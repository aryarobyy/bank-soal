package response

import "latih.in-be/internal/model"

func ExamScoreResponse(data model.ExamScore) model.ExamScoreResponse {
	examScoreRes := model.ExamScoreResponse{
		Id:         data.Id,
		ExamId:     data.ExamId,
		UserId:     data.UserId,
		TotalScore: data.TotalScore,
		Status:     data.Status,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}

	return examScoreRes
}

func ExamScoresResponse(data []model.ExamScore) []model.ExamScoreResponse {
	examScoresRes := []model.ExamScoreResponse{}

	for _, examScore := range data {
		examScoreRes := ExamScoreResponse(examScore)
		examScoresRes = append(examScoresRes, examScoreRes)
	}

	return examScoresRes
}