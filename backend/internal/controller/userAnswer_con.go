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

type UserAnswerController struct {
	service service.UserAnswerService
}

func NewUserAnswerController(service service.UserAnswerService) *UserAnswerController {
	return &UserAnswerController{service}
}

func (h *UserAnswerController) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var data model.UserAnswer
	if err := c.ShouldBindJSON(&data); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid input format")
		return
	}

	if err := h.service.Create(ctx, &data); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, data, "user answer created")
}

func (h *UserAnswerController) GetById(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Query("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	data, err := h.service.GetById(ctx, id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	userAnswerRes := response.UserAnswerResponse(*data)

	helper.Success(c, userAnswerRes, "user answer found")
}

func (h *UserAnswerController) GetMany(c *gin.Context) {
	ctx := c.Request.Context()

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	data, total, err := h.service.GetMany(ctx, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	userAnswersRes := response.UserAnswersResponse(data)

	helper.Success(c, gin.H{"data": userAnswersRes, "total": total}, "user answers found")
}

func (h *UserAnswerController) Update(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	var input model.UserAnswer
	if err := c.ShouldBindJSON(&input); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid input format")
		return
	}

	updated, err := h.service.Update(ctx, id, &input)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	userAnswerRes := response.UserAnswerResponse(*updated)

	helper.Success(c, userAnswerRes, "user answer updated")
}

func (h *UserAnswerController) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.service.Delete(ctx, id); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, nil, "user answer deleted")
}

func (h *UserAnswerController) GetByExamSessionId(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Query("session_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid session id")
		return
	}

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	data, total, err := h.service.GetByExamSessionId(ctx, id, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if total == 0 {
		helper.Error(c, http.StatusNotFound, "no user answers found for this exam session")
		return
	}

	userAnswersRes := response.UserAnswersResponse(data)

	helper.Success(c, gin.H{"data": userAnswersRes, "total": total}, "user answers found")
}

func (h *UserAnswerController) GetByQuestionId(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Query("question_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid question id")
		return
	}

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	data, total, err := h.service.GetByQuestionId(ctx, id, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if total == 0 {
		helper.Error(c, http.StatusNotFound, "no user answers found for this question")
		return
	}

	userAnswersRes := response.UserAnswersResponse(data)

	helper.Success(c, gin.H{"data": userAnswersRes, "total": total}, "user answers found")
}

func (h *UserAnswerController) GetUserAnswer(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Query("user_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}

	sessionIdStr := c.Query("session_id")
	sessionId, err := strconv.Atoi(sessionIdStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid session id")
		return
	}

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	data, total, err := h.service.GetUserAnswer(ctx, id, sessionId, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if total == 0 {
		helper.Error(c, http.StatusNotFound, "no user answers found for this user")
		return
	}

	userAnswersRes := response.UserAnswersResponse(data)

	helper.Success(c, gin.H{"data": userAnswersRes, "total": total}, "user answers found")
}
