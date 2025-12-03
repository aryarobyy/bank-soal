package response

import "latih.in-be/internal/model"

func UserResponse(data model.User) model.UserResponse {
	userRes := model.UserResponse{
		Id:           data.Id,
		Name:         &data.Name,
		Nim:          data.Nim,
		Nip:          data.Nip,
		Username:     data.Username,
		ImgUrl:       &data.ImgUrl,
		Email:        data.Email,
		Role:         data.Role,
		Major:        &data.Major,
		AcademicYear: &data.AcademicYear,
		Faculty:      &data.Faculty,
		Status:       data.Status,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}

	return userRes
}

func UsersResponse(data []model.User) []model.UserResponse {
	usersRes := []model.UserResponse{}

	for _, user := range data {
		userRes := UserResponse(user)
		usersRes = append(usersRes, userRes)
	}

	return usersRes
}
