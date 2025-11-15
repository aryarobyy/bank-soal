package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"latih.in-be/internal/model"
	"latih.in-be/internal/repository"
	"latih.in-be/utils/helper"
)

type UserService interface {
	Register(ctx context.Context, data model.RegisterCredential) error
	Login(ctx context.Context, cred model.LoginCredential) (*model.User, string, string, error)
	GetById(ctx context.Context, id int) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, data model.User, id int) (*model.User, error)
	Delete(ctx context.Context, id int) error
	GetMany(ctx context.Context, limit int, offset int) ([]model.User, int64, error)
	GetByNim(ctx context.Context, nim string) (*model.User, error)
	GetByUsn(ctx context.Context, username string) (*model.User, error)
	GetByNidn(ctx context.Context, nidn string) (*model.User, error)
	GetByName(ctx context.Context, name string, limit int, offset int) ([]model.User, int64, error)
	GetByRole(ctx context.Context, role string, limit int, offset int) ([]model.User, int64, error)
	ChangePassword(ctx context.Context, id int, newPassword string) error
	ChangeRole(ctx context.Context, id int, role model.Role, userRole model.Role) error
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
	BulkInsert(ctx context.Context, batchUser model.BulkUserCredential, prefix string, start int, end int) ([]model.BulkUserOutput, error)
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

	rules := map[string]int{
		"name":    256,
		"email":   512,
		"faculty": 128,
		"major":   256,
	}

	if err := helper.ValidateFieldLengths(data, rules); err != nil {
		return fmt.Errorf(err.Error())
	}

	switch data.Role {
	case string(model.RoleLecturer):
		if data.Nip == "" || data.Nidn == "" {
			return fmt.Errorf("lecturer must have both NIP and NIDN")
		}
	case string(model.RoleUser):
		if data.Nim == "" {
			return fmt.Errorf("user must have NIM")
		}
	case string(model.RoleAdmin):
	default:
		return fmt.Errorf("invalid role: %s", data.Role)
	}

	var nimPtr, nipPtr, nidnPtr, academicYearPtr, usernamePtr *string

	if data.Nim != "" {
		nimPtr = &data.Nim
	}
	if data.Nip != "" {
		nipPtr = &data.Nip
	}
	if data.Nidn != "" {
		nidnPtr = &data.Nidn
	}
	if data.Username != "" {
		usernamePtr = &data.Username
	}
	if data.Role == string(model.RoleAdmin) {
		academicYearPtr = &data.AcademicYear
	}

	if data.Role == string(model.RoleLecturer) || data.Role == string(model.RoleAdmin) {
		academicYearPtr = nil
	}

	if data.Role == string(model.RoleLecturer) || data.Role == string(model.RoleUser) {
		usernamePtr = nil
	}

	var academicYearVal string
	if academicYearPtr != nil {
		academicYearVal = *academicYearPtr
	}

	userData := model.User{
		Name:         data.Name,
		Email:        data.Email,
		Password:     string(hashedPassword),
		Major:        data.Major,
		Faculty:      data.Faculty,
		AcademicYear: academicYearVal,
		Nim:          nimPtr,
		Nip:          nipPtr,
		Username:     usernamePtr,
		Nidn:         nidnPtr,
		Role:         model.Role(data.Role),
	}

	_, err = s.repo.Register(ctx, userData)
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	return nil
}

func (s *userService) Login(ctx context.Context, cred model.LoginCredential) (*model.User, string, string, error) {
	loginId := cred.LoginId

	loginType := helper.DetectLoginType(loginId)

	var (
		data *model.User
		err  error
	)

	rules := map[string]int{
		"login_id": 256,
	}

	if err := helper.ValidateFieldLengths(data, rules); err != nil {
		return nil, "", "", err
	}

	switch loginType {
	case "nidn":
		data, err = s.repo.GetByNidn(ctx, loginId)
	case "nim":
		data, err = s.repo.GetByNim(ctx, loginId)
	default:
		data, err = s.repo.GetByUsn(ctx, loginId)
	}

	if err != nil {
		return nil, "", "", fmt.Errorf("user not found")
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

	effectiveRole := data.Role
	if effectiveRole == "" {
		effectiveRole = oldUser.Role
	}

	var academicYearPtr *string
	if effectiveRole == model.RoleAdmin {
		academicYearPtr = &data.AcademicYear
	} else {
		academicYearPtr = nil
	}

	if academicYearPtr != nil {
		data.AcademicYear = *academicYearPtr
	} else {
		data.AcademicYear = ""
	}

	roleChanged := effectiveRole != oldUser.Role

	if roleChanged {
		if effectiveRole == "lecturer" {
			if (data.Nip == nil || *data.Nip == "") && (data.Nidn == nil || *data.Nidn == "") {
				return nil, fmt.Errorf("lecturer must provide either Nip or Nidn")
			}
			emptyNim := ""
			data.Nim = &emptyNim
		}

		if effectiveRole == "user" {
			if data.Nim == nil || *data.Nim == "" {
				return nil, fmt.Errorf("user must provide Nim")
			}
			emptyStr := ""
			data.Nip = &emptyStr
			data.Nidn = &emptyStr
		}

		data.Role = effectiveRole
	} else {
		if effectiveRole == "lecturer" {
			if data.Nim != nil && *data.Nim != "" {
				return nil, fmt.Errorf("only user can have Nim")
			}
		}

		if effectiveRole == "user" {
			if (data.Nip != nil && *data.Nip != "") || (data.Nidn != nil && *data.Nidn != "") {
				return nil, fmt.Errorf("only lecturers can have Nip or Nidn")
			}
		}
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

func (s *userService) GetMany(ctx context.Context, limit int, offset int) ([]model.User, int64, error) {
	dataList, total, err := s.repo.GetMany(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all users: %w", err)
	}

	return dataList, total, nil
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

func (s *userService) GetByNidn(ctx context.Context, nidn string) (*model.User, error) {
	if len(nidn) > 10 {
		return nil, fmt.Errorf("nidn cannot be more than 9 characters: %s", nidn)
	}

	data, err := s.repo.GetByNidn(ctx, nidn)
	if err != nil {
		return nil, fmt.Errorf("user with nidn %q not found: %w", nidn, err)
	}

	return data, nil
}

func (s *userService) GetByUsn(ctx context.Context, username string) (*model.User, error) {
	if len(username) > 256 {
		return nil, fmt.Errorf("username cannot be more than 9 characters: %s", username)
	}

	data, err := s.repo.GetByUsn(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("user with username %q not found: %w", username, err)
	}

	return data, nil
}

func (s *userService) GetByName(ctx context.Context, name string, limit int, offset int) ([]model.User, int64, error) {
	if containsNumber(name) {
		return nil, 0, fmt.Errorf("name cannot contain numbers")
	}

	dataList, total, err := s.repo.GetByName(ctx, name, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("user with name %q not found: %w", name, err)
	}
	return dataList, total, nil
}

func (s *userService) GetByRole(ctx context.Context, role string, limit int, offset int) ([]model.User, int64, error) {
	if role != "admin" && role != "user" && role != "lecturer" {
		return nil, 0, fmt.Errorf("invalid role: %s", role)
	}

	dataList, total, err := s.repo.GetByRole(ctx, role, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("user with role %q not found: %w", role, err)
	}

	return dataList, total, nil
}

func (s *userService) ChangePassword(ctx context.Context, id int, newPassword string) error {
	if newPassword == "" {
		return fmt.Errorf("new password cannot be empty")
	}

	if len(newPassword) < 6 {
		return fmt.Errorf("password must be at least 6 characters")
	}

	user, err := s.GetById(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	if user == nil {
		return fmt.Errorf("user not found")
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

func (s *userService) ChangeRole(ctx context.Context, id int, role model.Role, userRole model.Role) error {
	if role == model.RoleAdmin && model.Role(userRole) != model.RoleSuperAdmin {
		return fmt.Errorf("you dont have permission to assign admin role")
	}

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

func (s *userService) BulkInsert(ctx context.Context, batchUser model.BulkUserCredential, prefix string, start int, end int) ([]model.BulkUserOutput, error) {
	nims := helper.GenerateNim(prefix, start, end)

	var users []model.User
	var credentials []model.BulkUserOutput

	for _, nim := range nims {
		plainPw, _ := helper.GenerateRandomPassword(12)
		hashed, _ := bcrypt.GenerateFromPassword([]byte(plainPw), bcrypt.DefaultCost)

		users = append(users, model.User{
			Nim:          &nim,
			Password:     string(hashed),
			Role:         model.RoleUser,
			Major:        "Informatika",
			Faculty:      "Teknik",
			AcademicYear: batchUser.AcademicYear,
		})

		credentials = append(credentials, model.BulkUserOutput{
			Nim:      nim,
			Password: plainPw,
		})
	}

	_, err := s.repo.BulkInsert(ctx, users)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			return nil, fmt.Errorf("failed to insert: some data has been registered (duplicate)")
		}
		return nil, fmt.Errorf("failed to bulk insert users: %w", err)
	}

	return credentials, nil
}
