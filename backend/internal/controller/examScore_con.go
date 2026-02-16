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

type ExamScoreController struct {
	service service.ExamScoreService
}

func NewExamScoreController(s service.ExamScoreService) *ExamScoreController {
	return &ExamScoreController{service: s}
}

func (h *ExamScoreController) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var data model.ExamScore
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

func (h *ExamScoreController) GetById(c *gin.Context) {
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

	examScoreRes := response.ExamScoreResponse(*data)

	helper.Success(c, examScoreRes, "data found")
}

func (h *ExamScoreController) GetMany(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Query("exam_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid exam id")
		return
	}

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.service.GetMany(ctx, id, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}
	if len(data) == 0 {
		helper.Error(c, http.StatusNotFound, "no Data found")
		return
	}

	examScoresRes := response.ExamScoresResponse(data)

	helper.Success(c, examScoresRes, "data found")
}

func (h *ExamScoreController) Update(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid score id")
		return
	}
	var data model.ExamScore

	if err := c.ShouldBindJSON(&data); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	updatedData, err := h.service.Update(ctx, data, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	examScoreRes := response.ExamScoreResponse(*updatedData)

	helper.Success(c, examScoreRes, "data updated")
}

func (h *ExamScoreController) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid score id")
		return
	}
	err = h.service.Delete(ctx, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(c, nil, "data deleted")
}
