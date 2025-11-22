package response

import "latih.in-be/internal/model"

func UserAnswerResponse(data model.UserAnswer) model.UserAnswerResponse {
	userAnswerRes := model.UserAnswerResponse{
		Id:            data.Id,
		ExamSessionId: data.ExamSessionId,
		UserId:        data.UserId,
		QuestionId:    data.QuestionId,
		Answer:        data.Answer,
		IsCorrect:     data.IsCorrect,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}

	return userAnswerRes
}

func UserAnswersResponse(data []model.UserAnswer) []model.UserAnswerResponse {
	userAnswersRes := []model.UserAnswerResponse{}

	for _, userAnswer := range data {
		userAnswerRes := UserAnswerResponse(userAnswer)
		userAnswersRes = append(userAnswersRes, userAnswerRes)
	}

	return userAnswersRes
}