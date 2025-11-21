package response

import "latih.in-be/internal/model"

func SubjectResponse(data model.Subject) model.SubjectResponse {
	subjectRes := model.SubjectResponse{
		Id:    data.Id,
		Title: data.Title,
		Code:  data.Code,
	}

	return subjectRes
}

func SubjectsResponse(data []model.Subject) []model.SubjectResponse {
	subjectsRes := []model.SubjectResponse{}

	for _, subject := range data {
		subjectRes := SubjectResponse(subject)
		subjectsRes = append(subjectsRes, subjectRes)
	}

	return subjectsRes
}