package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"latih.in-be/internal/model"
	"latih.in-be/internal/service"
	"latih.in-be/utils/helper"
	"latih.in-be/utils/response"
)

type QuestionController struct {
	service service.QuestionService
}

func NewQuestionController(s service.QuestionService) *QuestionController {
	return &QuestionController{service: s}
}

func (h *QuestionController) GetById(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Query("id")
	id := helper.BindToInt(idStr)

	data, err := h.service.GetById(ctx, id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	questionRes := response.QuestionResponse(*data)

	helper.Success(c, questionRes, "data found")
}

func (h *QuestionController) GetMany(c *gin.Context) {
	ctx := c.Request.Context()

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	data, total, err := h.service.GetMany(ctx, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	questionsRes := response.QuestionsResponse(data)

	helper.Success(c, gin.H{"data": questionsRes, "total": total}, "data found")
}

func (h *QuestionController) Update(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Param("id")
	id := helper.BindToInt(idStr)

	userIdVal, exists := c.Get("user_id")
	if !exists {
		helper.Error(c, http.StatusUnauthorized, "user id not found")
		return
	}
	userId := userIdVal.(int)

	subjectId := c.PostForm("subject_id")
	creatorId := c.PostForm("creator_id")
	questionText := c.PostForm("question_text")
	difficulty := c.PostForm("difficulty")
	answer := c.PostForm("answer")
	score := c.PostForm("score")
	optionsJson := c.PostForm("options")
	imgDelete := c.PostForm("img_delete")

	data := model.UpdateQuestion{
		SubjectId:    helper.BindToIntPtr(subjectId),
		CreatorId:    helper.BindToIntPtr(creatorId),
		Score:        helper.BindToIntPtr(score),
		QuestionText: helper.BindAndConvertToPtr(questionText),
		Difficulty:   (*model.Difficulty)(helper.BindAndConvertToPtr(difficulty)),
		Answer:       helper.BindAndConvertToPtr(answer),
		ImgDelete:    helper.BindAndConvertToBoolPtr(imgDelete),
	}

	if optionsJson != "" {
		if err := json.Unmarshal([]byte(optionsJson), &data.Options); err != nil {
			helper.Error(c, http.StatusBadRequest, "invalid options format")
			return
		}
	}

	updated, err := h.service.Update(ctx, c, data, id, userId)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	questionRes := response.QuestionResponse(*updated)

	helper.Success(c, questionRes, "question updated")
}

func (h *QuestionController) Delete(c *gin.Context) {
	ctx := c.Request.Context()

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

	err = h.service.Delete(ctx, id, userId)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(c, nil, "data deleted")
}

func (h *QuestionController) CreateWithOptions(c *gin.Context) {
	ctx := c.Request.Context()

	var data model.Question

	data.SubjectId, _ = strconv.Atoi(c.PostForm("subject_id"))
	data.CreatorId, _ = strconv.Atoi(c.PostForm("creator_id"))
	data.QuestionText = c.PostForm("question_text")
	data.Difficulty = model.Difficulty(c.PostForm("difficulty"))
	data.Answer = c.PostForm("answer")
	data.Score, _ = strconv.Atoi(c.PostForm("score"))

	optionsJson := c.PostForm("options")
	if optionsJson == "" {
		helper.Error(c, http.StatusBadRequest, "options cannot be empty")
		return
	}

	if err := json.Unmarshal([]byte(optionsJson), &data.Options); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid options format")
		return
	}

	if err := h.service.Create(ctx, c, &data); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	questionRes := response.QuestionResponse(data)

	helper.Success(c, questionRes, "question created successfully")
}

func (h *QuestionController) CreateFromJson(c *gin.Context) {
	ctx := c.Request.Context()

	file, err := c.FormFile("file")
	println("SLSKSKAOKS", file)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "File tidak ditemukan. Gunakan key 'file' untuk upload")
		return
	}

	if file.Header.Get("Content-Type") != "application/json" {
		helper.Error(c, http.StatusBadRequest, "file berformat json")
		return
	}

	if err := h.service.CreateFromJson(ctx, file); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	helper.Success(c, nil, "questions upload successfully")
}

func (h *QuestionController) GetByExam(c *gin.Context) {
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

	data, total, err := h.service.GetByExam(ctx, id, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	questionsRes := response.QuestionsResponse(data)

	helper.Success(c, gin.H{"data": questionsRes, "total": total}, "data found")
}

func (h *QuestionController) GetByCreator(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Query("creator_id")
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

	data, total, err := h.service.GetByCreatorId(ctx, id, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	questionsRes := response.QuestionsResponse(data)

	helper.Success(c, gin.H{"data": questionsRes, "total": total}, "data found")
}

func (h *QuestionController) GetByDiff(c *gin.Context) {
	ctx := c.Request.Context()

	diff := c.Query("diff")

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	data, total, err := h.service.GetByDifficult(ctx, diff, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	questionsRes := response.QuestionsResponse(data)

	helper.Success(c, gin.H{"data": questionsRes, "total": total}, "data found")
}

func (h *QuestionController) GetBySubject(c *gin.Context) {
	ctx := c.Request.Context()

	subjectStr := c.Query("subject_id")
	subject := 0

	if subjectStr != "" {
		if l, err := strconv.Atoi(subjectStr); err == nil && l > 0 {
			subject = l
		}
	}

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	data, total, err := h.service.GetBySubject(ctx, subject, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	questionsRes := response.QuestionsResponse(data)

	helper.Success(c, gin.H{"data": questionsRes, "total": total}, "data found")
}
