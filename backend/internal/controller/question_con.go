package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"latih.in-be/internal/model"
	"latih.in-be/internal/service"
	"latih.in-be/utils/helper"
)

type QuestionController struct {
	service service.QuestionService
}

func NewQuestionController(s service.QuestionService) *QuestionController {
	return &QuestionController{service: s}
}

func (h *QuestionController) Create(c *gin.Context) {
	var data model.Question
	if err := c.ShouldBindJSON(&data); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Create(c.Request.Context(), data); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(c, data, "data created")
}

func (h *QuestionController) GetById(c *gin.Context) {
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

func (h *QuestionController) GetAll(c *gin.Context) {
	data, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, data, "data found")
}

func (h *QuestionController) Update(c *gin.Context) {
	idStr1 := c.Param("id")
	id, err := strconv.Atoi(idStr1)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	idStr2 := c.Query("creator_id")
	userId, err := strconv.Atoi(idStr2)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	var data model.Question

	if err := c.ShouldBindJSON(&data); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	updatedData, err := h.service.Update(c, data, id, userId)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(c, updatedData, "data updated")
}

func (h *QuestionController) Delete(c *gin.Context) {
	idStr1 := c.Param("id")
	id, err := strconv.Atoi(idStr1)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	idStr2 := c.Query("user_id")
	userId, err := strconv.Atoi(idStr2)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	err = h.service.Delete(c, id, userId)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(c, nil, "data deleted")
}

func (h *QuestionController) CreateWithOptions(c *gin.Context) {
	var data model.Question

	if err := c.ShouldBindJSON(&data); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.Create(c, data); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, nil, "question created successfully")
}

func (h *QuestionController) CreateFromJson(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "File tidak ditemukan. Gunakan key 'file' untuk upload")
		return
	}

	if file.Header.Get("Content-Type") != "application/json" {
		helper.Error(c, http.StatusBadRequest, "file berformat json")
		return
	}

	if err := h.service.CreateFromJson(c, file); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	helper.Success(c, nil, "questions upload successfully")
}

func (h *QuestionController) GetByExam(c *gin.Context) {
	idStr := c.Query("exam_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid exam id")
		return
	}

	data, err := h.service.GetByExam(c.Request.Context(), id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, data, "data found")
}

func (h *QuestionController) GetByCreator(c *gin.Context) {
	idStr := c.Query("creator_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid exam id")
		return
	}

	data, err := h.service.GetByCreatorId(c.Request.Context(), id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, data, "data found")
}

func (h *QuestionController) GetByDiff(c *gin.Context) {
	diff := c.Query("diff")

	data, err := h.service.GetByDifficult(c.Request.Context(), diff)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, data, "data found")
}
