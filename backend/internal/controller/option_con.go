package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"latih.in-be/internal/model"
	"latih.in-be/internal/service"
	"latih.in-be/utils/helper"
	"latih.in-be/utils/response"
)

type OptionController struct {
	service service.OptionService
}

func NewOptionController(s service.OptionService) *OptionController {
	return &OptionController{service: s}
}

func (h *OptionController) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var data model.Option
	if err := c.ShouldBindJSON(&data); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Create(ctx, data); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(c, data, "exam score created")
}

func (h *OptionController) GetById(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid data id")
		return
	}

	data, err := h.service.GetById(ctx, id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	optionRes := response.OptionResponse(*data)

	helper.Success(c, optionRes, "data found")
}

func (h *OptionController) GetMany(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Query("question_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid limit")
		return
	}

	data, err := h.service.GetMany(ctx, id, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	optionsRes := response.OptionsResponse(data)

	helper.Success(c, optionsRes, "data found")
}

func (h *OptionController) Update(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	var data model.Option

	if err := c.ShouldBindJSON(&data); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	updatedData, err := h.service.Update(ctx, data, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	optionRes := response.OptionResponse(*updatedData)

	helper.Success(c, optionRes, "data updated")
}

func (h *OptionController) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	err = h.service.Delete(ctx, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(c, nil, "data deleted")
}
