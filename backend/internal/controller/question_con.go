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
	var data *model.Question
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

func (h *QuestionController) GetMany(c *gin.Context) {
	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid limit")
		return
	}
	data, err := h.service.GetMany(c, limit, offset)
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

	userIdVal, exists := c.Get("user_id")
	if !exists {
		helper.Error(c, http.StatusUnauthorized, "user id not found in context")
		return
	}
	userId := userIdVal.(int)

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

func (h *QuestionController) CreateWithOptions(c *gin.Context) {
	var data *model.Question

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

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid limit")
		return
	}

	data, err := h.service.GetByExam(c, id, limit, offset)
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

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid limit")
		return
	}

	data, err := h.service.GetByCreatorId(c, id, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, data, "data found")
}

func (h *QuestionController) GetByDiff(c *gin.Context) {
	diff := c.Query("diff")

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid limit")
		return
	}

	data, err := h.service.GetByDifficult(c, diff, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, data, "data found")
}

func (h *QuestionController) GetBySubject(c *gin.Context) {
	subjectStr := c.Query("subject_id")
	subject := 0

	if subjectStr != "" {
		if l, err := strconv.Atoi(subjectStr); err == nil && l > 0 {
			subject = l
		}
	}

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid limit")
		return
	}

	data, err := h.service.GetBySubject(c, subject, limit, offset)

	helper.Success(c, data, "data found")
}
