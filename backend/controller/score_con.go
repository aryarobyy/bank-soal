package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"latih.in-be/helper"
	"latih.in-be/model"
	"latih.in-be/service"
)

type ScoreController struct {
	service service.ScoreService
}

func NewScoreController(s service.ScoreService) *ScoreController {
	return &ScoreController{service: s}
}

func (h *ScoreController) Create(c *gin.Context) {
	var data model.Score
	if err := c.ShouldBindJSON(&data); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Create(c.Request.Context(), data); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(c, data, "exam created")
}

func (h *ScoreController) GetById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid data id")
		return
	}

	data, err := h.service.GetById(c.Request.Context(), id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, data, "data found")
}

func (h *ScoreController) GetAll(c *gin.Context) {
	idStr := c.Param("qId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid question id")
		return
	}
	data, err := h.service.GetAll(c.Request.Context(), id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, data, "data found")
}

func (h *ScoreController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	var data model.Score

	if err := c.ShouldBindJSON(&data); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	updatedData, err := h.service.Update(c, data, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(c, updatedData, "data updated")
}

func (h *ScoreController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	err = h.service.Delete(c, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(c, nil, "data deleted")
}
