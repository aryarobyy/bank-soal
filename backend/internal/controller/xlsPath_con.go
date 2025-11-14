package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"latih.in-be/internal/service"
	"latih.in-be/utils/helper"
)

type XlsPathController struct {
	service service.XlsPathService
}

func NewXlsPathController(s service.XlsPathService) *XlsPathController {
	return &XlsPathController{service: s}
}

func (h *XlsPathController) GetById(c *gin.Context) {
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

	helper.Success(c, data, "data found")
}

func (h *XlsPathController) GetMany(c *gin.Context) {
	ctx := c.Request.Context()

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid limit and offset")
		return
	}

	data, err := h.service.GetMany(ctx, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, data, "data found")
}

func (h *XlsPathController) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Query("id")
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
