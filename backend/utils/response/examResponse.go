package response

import "latih.in-be/internal/model"

func ExamResponse(data model.Exam) model.ExamResponse {
	examRes := model.ExamResponse{
		Id:          data.Id,
		Title:       data.Title,
		Description: data.Description,
		Difficulty:  data.Difficulty,
		LongTime:    data.LongTime,
		CreatorId:   data.CreatorId,
		StartedAt:   data.StartedAt,
		FinishedAt:  data.FinishedAt,
		Score:       data.Score,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}

	return examRes
}

func ExamsResponse(data []model.Exam) []model.ExamResponse {
	examsRes := []model.ExamResponse{}

	for _, exam := range data {
		examRes := ExamResponse(exam)
		examsRes = append(examsRes, examRes)
	}

	return examsRes
}