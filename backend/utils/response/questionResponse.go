package response

import "latih.in-be/internal/model"

func QuestionResponse(data model.Question) model.QuestionResponse {
	questionRes := model.QuestionResponse{
		Id:           data.Id,
		SubjectId:    data.SubjectId,
		Subject:      data.Subject,
		CreatorId:    data.CreatorId,
		QuestionText: data.QuestionText,
		Difficulty:   data.Difficulty,
		Answer:       data.Answer,
		Score:        data.Score,
		ImgUrl:       data.ImgUrl,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
		Options:      data.Options,
	}

	return questionRes
}

func QuestionsResponse(data []model.Question) []model.QuestionResponse {
	questionsRes := []model.QuestionResponse{}

	for _, question := range data {
		questionRes := QuestionResponse(question)
		questionsRes = append(questionsRes, questionRes)
	}

	return questionsRes
}