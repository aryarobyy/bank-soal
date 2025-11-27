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

type ExamSessionController struct {
	service service.ExamSessionService
}

func NewExamSessionController(s service.ExamSessionService) *ExamSessionController {
	return &ExamSessionController{
		service: s,
	}
}

func (h *ExamSessionController) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req model.ExamSession
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	userIdVal, exists := c.Get("user_id")
	if !exists {
		helper.Error(c, http.StatusUnauthorized, "user id not found in context %w")
		return
	}
	userId := userIdVal.(int)

	exam, err := h.service.Create(ctx, req, userId, req.ExamId)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, exam, "exam created")
}

func (h *ExamSessionController) GetById(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	data, err := h.service.GetById(ctx, id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, "session not found %s")
		return
	}

	sessionRes := response.SessionResponse(*data)

	helper.Success(c, sessionRes, "session found")
}

func (h *ExamSessionController) Update(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req model.UpdateExamSession
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.service.Update(ctx, id, req)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	sessionRes := response.SessionResponse(*data)

	helper.Success(c, sessionRes, "update session success")
}

func (h *ExamSessionController) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.service.Delete(ctx, id); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, nil, "data deleted")
}

func (h *ExamSessionController) GetMany(c *gin.Context) {
	ctx := c.Request.Context()

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.service.GetMany(ctx, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	sessionsRes := response.SessionsResponse(data)

	helper.Success(c, sessionsRes, "data updated")
}

func (h *ExamSessionController) UpdateCurrNo(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req model.UpdateCurrNo
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	data, err := h.service.UpdateCurrNo(ctx, id, req)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, data, "data found")
}

func (h *ExamSessionController) FinishExam(c *gin.Context) {
	ctx := c.Request.Context()

	var req model.FinishExam
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	data, err := h.service.FinishExam(ctx, req.UserId, req.Id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, data, "data found")
}
