package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"latih.in-be/internal/model"
	"latih.in-be/internal/service"
	"latih.in-be/utils/helper"
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

func (h *ExamController) GetMany(c *gin.Context) {
	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid limit")
		return
	}
	data, total, err := h.service.GetMany(c, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, gin.H{"data": data, "total": total}, "data found")
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

	userIdVal, exists := c.Get("user_id")
	if !exists {
		helper.Error(c, http.StatusUnauthorized, "user id not found in context")
		return
	}
	userId := userIdVal.(int)

	updatedData, err := h.service.Update(c, data, id, userId)
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

	userIdVal, exists := c.Get("user_id")
	if !exists {
		helper.Error(c, http.StatusUnauthorized, "user id not found in context")
		return
	}
	userId := userIdVal.(int)

	err = h.service.Delete(c, id, userId)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(c, nil, "data deleted")
}
