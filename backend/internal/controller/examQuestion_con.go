package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"latih.in-be/internal/service"
	"latih.in-be/utils/helper"
)

type ExamQuestionHandler struct {
	service service.ExamQuestionService
}

func NewExamQuestionHandler(s service.ExamQuestionService) *ExamQuestionHandler {
	return &ExamQuestionHandler{service: s}
}

func (h *ExamQuestionHandler) AddQuestionsToExam(c *gin.Context) {
	examIdStr := c.Param("examId")
	examId, err := strconv.Atoi(examIdStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid exam Id")
		return
	}

	var body struct {
		QuestionIds []int `json:"question_ids"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}
	if len(body.QuestionIds) == 0 {
		helper.Error(c, http.StatusBadRequest, "question_ids cannot be empty")
		return
	}

	if err := h.service.AddQuestionToExam(c, examId, body.QuestionIds); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, nil, "questions added to exam successfully")
}

func (h *ExamQuestionHandler) UpdateQuestionsInExam(c *gin.Context) {
	examIdStr := c.Param("id")
	examId, err := strconv.Atoi(examIdStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid exam Id")
		return
	}

	var body struct {
		QuestionIds []int `json:"question_ids"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}
	if len(body.QuestionIds) == 0 {
		helper.Error(c, http.StatusBadRequest, "question_ids cannot be empty")
		return
	}

	if err := h.service.UpdateQuestionsInExam(c, examId, body.QuestionIds); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, nil, "exam questions updated successfully")
}

func (h *ExamQuestionHandler) RemoveQuestionsFromExam(c *gin.Context) {
	examIdStr := c.Param("id")
	examId, err := strconv.Atoi(examIdStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid exam ID")
		return
	}

	var body struct {
		QuestionIds []int `json:"question_ids"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}
	if len(body.QuestionIds) == 0 {
		helper.Error(c, http.StatusBadRequest, "question_ids cannot be empty")
		return
	}

	if err := h.service.RemoveQuestionsFromExam(c, examId, body.QuestionIds); err != nil {
		helper.Error(c, http.StatusInternalServerError, "failed to delete questions")
		return
	}

	helper.Success(c, nil, "data deleted successfully")
}
