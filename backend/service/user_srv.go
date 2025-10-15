package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"latih.in-be/helper"
	"latih.in-be/model"
	"latih.in-be/repository"
)

type UserService interface {
	Register(ctx context.Context, data model.RegisterCredential) error
	Login(ctx context.Context, cred model.LoginCredential) (*model.User, string, string, error)
	GetById(ctx context.Context, id int) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, data model.User, id int) (*model.User, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]model.User, error)
	GetByNim(ctx context.Context, nim string) (*model.User, error)
	GetByName(ctx context.Context, name string) ([]model.User, error)
	GetByRole(ctx context.Context, role string) ([]model.User, error)
	ChangePassword(ctx context.Context, id int, oldPassword, newPassword string) error
	ChangeRole(ctx context.Context, id int, role model.Role) error
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Register(ctx context.Context, data model.RegisterCredential) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	existingData, _ := s.repo.GetByEmail(ctx, data.Email)
	if existingData != nil {
		return fmt.Errorf("email %s already used", data.Email)
	}

	userData := model.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: string(hashedPassword),
		Major:    data.Major,
		Nim:      data.Nim,
		Faculty:  data.Faculty,
		Role:     model.RoleUser,
	}

	_, err = s.repo.Register(ctx, userData)
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	return nil
}

func (s *userService) Login(ctx context.Context, cred model.LoginCredential) (*model.User, string, string, error) {
	data, err := s.repo.GetByEmail(ctx, cred.Email)
	if err != nil {
		return nil, "", "", fmt.Errorf("user with email %s not found", cred.Email)
	}

	if bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(cred.Password)) != nil {
		return nil, "", "", fmt.Errorf("wrong password")
	}

	accessToken, err := helper.GenerateAccessToken(data)
	if err != nil {
		return nil, "", "", fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := helper.GenerateRefreshToken(data)
	if err != nil {
		return nil, "", "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return data, accessToken, refreshToken, nil
}

func (s *userService) GetById(ctx context.Context, id int) (*model.User, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return data, nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	data, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("user with email %s not found: %w", email, err)
	}
	return data, nil
}

func (s *userService) Update(ctx context.Context, data model.User, id int) (*model.User, error) {
	oldUser, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	if oldUser.ImgUrl != "" && oldUser.ImgUrl != data.ImgUrl {
		if err := helper.DeleteImage(oldUser.ImgUrl); err != nil {
			return nil, fmt.Errorf("failed to delete old image: %w", err)
		}
	}

	updatedUser, err := s.repo.Update(ctx, data, id)
	if err != nil {
		if strings.Contains(err.Error(), "Unknown column") {
			var fieldName string
			parts := strings.Split(err.Error(), "'")
			if len(parts) >= 2 {
				fieldName = parts[1]
			}

			val := helper.GetFieldValue(data, fieldName)

			return nil, fmt.Errorf("field '%s' with value '%v' is undefined", fieldName, val)
		}

		return nil, fmt.Errorf("update gagal: %v", err)
	}

	return updatedUser, nil
}

func (s *userService) Delete(ctx context.Context, id int) error {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	if err := helper.DeleteImage(user.ImgUrl); err != nil {
		return fmt.Errorf("failed to delete image: %w", err)
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}
	return nil
}

func (s *userService) GetAll(ctx context.Context) ([]model.User, error) {
	dataList, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	return dataList, nil
}

func (s *userService) GetByNim(ctx context.Context, nim string) (*model.User, error) {
	if len(nim) > 9 {
		return nil, fmt.Errorf("nim cannot be more than 9 characters: %s", nim)
	}

	data, err := s.repo.GetByNim(ctx, nim)
	if err != nil {
		return nil, fmt.Errorf("user with nim %q not found: %w", nim, err)
	}

	return data, nil
}

func (s *userService) GetByName(ctx context.Context, name string) ([]model.User, error) {
	if containsNumber(name) {
		return nil, fmt.Errorf("name cannot contain numbers")
	}

	dataList, err := s.repo.GetByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("user with name %q not found: %w", name, err)
	}
	return dataList, nil
}

func (s *userService) GetByRole(ctx context.Context, role string) ([]model.User, error) {
	if role != "admin" && role != "user" && role != "lecturer" {
		return nil, fmt.Errorf("invalid role: %s", role)
	}

	dataList, err := s.repo.GetByRole(ctx, role)
	if err != nil {
		return nil, fmt.Errorf("user with role %q not found: %w", role, err)
	}

	return dataList, nil
}

func (s *userService) ChangePassword(ctx context.Context, id int, oldPassword, newPassword string) error {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	if bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(oldPassword)) != nil {
		return fmt.Errorf("old password is incorrect")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash new password: %w", err)
	}

	if err := s.repo.ChangePassword(ctx, id, string(hashedPassword)); err != nil {
		return fmt.Errorf("cannot change password: %w", err)
	}

	return nil
}

func (s *userService) ChangeRole(ctx context.Context, id int, role model.Role) error {
	if err := s.repo.ChangeRole(ctx, id, role); err != nil {
		return fmt.Errorf("cannot change role: %w", err)
	}
	return nil
}

func containsNumber(s string) bool {
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			return true
		}
	}
	return false
}

func (s *userService) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	userId, err := helper.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("invalid or expired refresh token: %w", err)
	}

	user, err := s.repo.GetById(ctx, userId)
	if err != nil {
		return "", fmt.Errorf("user not found: %w", err)
	}

	newAccessToken, err := helper.GenerateAccessToken(user)
	if err != nil {
		return "", fmt.Errorf("failed to generate new access token: %w", err)
	}

	return newAccessToken, nil
}
