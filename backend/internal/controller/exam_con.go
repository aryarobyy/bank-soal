package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"latih.in-be/utils/helper"
	"latih.in-be/internal/model"
	"latih.in-be/internal/service"
)

type ExamController struct {
	service service.ExamService
}

func NewExamController(s service.ExamService) *ExamController {
	return &ExamController{service: s}
}

func (h *ExamController) Create(c *gin.Context) {
	var data model.Exam
	if err := c.ShouldBindJSON(&data); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if data.Difficulty != "" &&
		data.Difficulty != "easy" &&
		data.Difficulty != "medium" &&
		data.Difficulty != "hard" {
		helper.Error(c, http.StatusBadRequest, "invalid difficulty value (must be 'easy', 'medium', or 'hard')")
		return
	}

	if err := h.service.Create(c.Request.Context(), data); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(c, data, "exam created")
}

func (h *ExamController) GetById(c *gin.Context) {
	idStr := c.Query("id")
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

func (h *ExamController) GetAll(c *gin.Context) {
	data, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, data, "data found")
}

func (h *ExamController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var data model.Exam
	if err := c.ShouldBindJSON(&data); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if data.Difficulty != "" &&
		data.Difficulty != "easy" &&
		data.Difficulty != "medium" &&
		data.Difficulty != "hard" {
		helper.Error(c, http.StatusBadRequest, "invalid difficulty value (must be 'easy', 'medium', or 'hard')")
		return
	}

	updatedData, err := h.service.Update(c, data, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, updatedData, "data updated")
}

func (h *ExamController) Delete(c *gin.Context) {
	idStr1 := c.Param("id")
	id, err := strconv.Atoi(idStr1)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	idStr2 := c.Param("userId")
	idUser, err := strconv.Atoi(idStr2)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	err = h.service.Delete(c, id, idUser)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(c, nil, "data deleted")
}
