package response

import "latih.in-be/internal/model"

func SessionResponse(data model.ExamSession) model.SessionResponse {
	sessionRes := model.SessionResponse{
		Id:         data.Id,
		UserId:     data.UserId,
		ExamId:     data.ExamId,
		StartedAt:  data.StartedAt,
		FinishedAt: data.FinishedAt,
		Status:     data.Status,
		CurrentNo:  data.CurrentNo,
	}
	{
		s := float64(data.Score)
		sessionRes.Score = &s
	}

	return sessionRes
}

func SessionsResponse(data []model.ExamSession) []model.SessionResponse {
	sessionsRes := []model.SessionResponse{}

	for _, session := range data {
		sessionRes := SessionResponse(session)
		sessionsRes = append(sessionsRes, sessionRes)
	}

	return sessionsRes
}
