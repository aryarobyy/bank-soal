package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

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

	if questionText != "" && len(questionText) > 5000 {
		helper.Error(c, http.StatusBadRequest, "question_text is too long")
		return
	}

	if answer != "" && len(answer) > 1000 {
		helper.Error(c, http.StatusBadRequest, "answer is too long")
		return
	}

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
		dec := json.NewDecoder(strings.NewReader(optionsJson))
		dec.DisallowUnknownFields()

		if err := dec.Decode(&data.Options); err != nil {
			helper.Error(c, http.StatusBadRequest, "invalid options format")
			return
		}

		if len(data.Options) > 4 {
			helper.Error(c, http.StatusBadRequest, "too many options")
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

	subjectIdStr := c.PostForm("subject_id")
	creatorIdStr := c.PostForm("creator_id")
	questionText := c.PostForm("question_text")
	difficulty := c.PostForm("difficulty")
	answer := c.PostForm("answer")
	scoreStr := c.PostForm("score")
	optionsJson := c.PostForm("options")

	if optionsJson == "" {
		helper.Error(c, http.StatusBadRequest, "options cannot be empty")
		return
	}

	data.SubjectId = helper.BindToInt(subjectIdStr)
	data.CreatorId = helper.BindToInt(creatorIdStr)
	data.QuestionText = questionText
	data.Difficulty = model.Difficulty(difficulty)
	data.Answer = answer
	data.Score = helper.BindToInt(scoreStr)

	if optionsJson != "" {
		if err := json.Unmarshal([]byte(optionsJson), &data.Options); err != nil {
			helper.Error(c, http.StatusBadRequest, "invalid options format")
			return
		}
	}

	if data.QuestionText == "" {
		helper.Error(c, http.StatusBadRequest, "question_text is required")
		return
	}

	if string(data.Difficulty) == "" {
		helper.Error(c, http.StatusBadRequest, "difficulty is required")
		return
	}

	if len(data.Options) == 0 {
		helper.Error(c, http.StatusBadRequest, "options are required")
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
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "File not found. Use key 'file' untuk upload")
		return
	}

	if file.Size > 10*1024*1024 {
		helper.Error(c, http.StatusBadRequest, "File is too big. Max 10MB")
		return
	}

	if file.Header.Get("Content-Type") != "application/json" && !strings.HasSuffix(strings.ToLower(file.Filename), ".json") {
		helper.Error(c, http.StatusBadRequest, "File must be json")
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

func (h *QuestionController) GetRandomQuestion(c *gin.Context) {
	ctx := c.Request.Context()

	totalStr := c.Query("total")
	subjectIdStr := c.Query("subject_id")
	creatorIdStr := c.Query("creator_id")

	total := helper.BindToInt(totalStr)
	subjectId := helper.BindToInt(subjectIdStr)
	creatorId := helper.BindToInt(creatorIdStr)

	data, err := h.service.GetRandomQuestion(ctx, total, &subjectId, &creatorId)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	questionsRes := response.QuestionsResponse(data)

	helper.Success(c, questionsRes, "data found")
}

func (h *QuestionController) GetByCreatorNSubject(c *gin.Context) {
	ctx := c.Request.Context()

	creatorIdStr := c.Query("creator_id")
	subjectIdStr := c.Query("subject_id")

	creatorId := helper.BindToInt(creatorIdStr)
	subjectId := helper.BindToInt(subjectIdStr)

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	data, total, err := h.service.GetByCreatorNSubject(ctx, creatorId, subjectId, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	questionsRes := response.QuestionsResponse(data)

	helper.Success(c, gin.H{"data": questionsRes, "total": total}, "data found")
}

func (h *QuestionController) GetByCreatorNDifficult(c *gin.Context) {
	ctx := c.Request.Context()

	creatorIdStr := c.Query("creator_id")
	diff := c.Query("diff")

	creatorId := helper.BindToInt(creatorIdStr)

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	data, total, err := h.service.GetByCreatorNDifficult(ctx, creatorId, diff, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	questionsRes := response.QuestionsResponse(data)

	helper.Success(c, gin.H{"data": questionsRes, "total": total}, "data found")
}
