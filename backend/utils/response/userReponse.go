package response

import (
	"latih.in-be/internal/model"
	"latih.in-be/utils/helper"
)

func UserResponse(data model.User) interface{} {
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

	return helper.SanitizeUserResponse(userRes)
}

func UsersResponse(data []model.User) interface{} {
	var usersRes []model.UserResponse

	for _, user := range data {
		userRes := model.UserResponse{
			Id:           user.Id,
			Name:         &user.Name,
			Nim:          user.Nim,
			Nip:          user.Nip,
			Username:     user.Username,
			ImgUrl:       &user.ImgUrl,
			Email:        user.Email,
			Role:         user.Role,
			Major:        &user.Major,
			AcademicYear: &user.AcademicYear,
			Faculty:      &user.Faculty,
			Status:       user.Status,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
		}
		usersRes = append(usersRes, userRes)
	}

	return helper.SanitizeUserResponse(usersRes)
}
