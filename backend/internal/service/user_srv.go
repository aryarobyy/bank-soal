package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"latih.in-be/internal/model"
	"latih.in-be/internal/repository"
	"latih.in-be/utils/helper"
	"latih.in-be/utils/update"
)

type UserService interface {
	Register(ctx context.Context, data model.RegisterCredential, requesterRole model.Role) error
	Login(ctx context.Context, cred model.LoginCredential) (*model.User, string, string, error)
	GetById(ctx context.Context, id int) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, c *gin.Context, data model.UpdateUser, id int, requesterRole model.Role, currentId int) (*model.User, error)
	Delete(ctx context.Context, id int, requesterRole model.Role) error
	GetMany(ctx context.Context, limit int, offset int) ([]model.User, int64, error)
	GetByNim(ctx context.Context, nim string, requesterRole model.Role) (*model.User, error)
	GetByUsn(ctx context.Context, username string, requesterRole model.Role) (*model.User, error)
	GetByNidn(ctx context.Context, nidn string, requesterRole model.Role) (*model.User, error)
	GetByName(ctx context.Context, name string, limit int, offset int) ([]model.User, int64, error)
	GetByRole(ctx context.Context, role string, limit int, offset int, requesterRole string) ([]model.User, int64, error)
	ChangePassword(ctx context.Context, id int, newPassword string, role model.Role) error
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

func (s *userService) Register(ctx context.Context, data model.RegisterCredential, requesterRole model.Role) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	existingEmail, _ := s.repo.GetByEmail(ctx, data.Email)
	if existingEmail != nil {
		return fmt.Errorf("email %s already used", data.Email)
	}

	existingUsn, _ := s.repo.GetByUsn(ctx, data.Username)
	if existingUsn != nil {
		return fmt.Errorf("username %s already used", data.Username)
	}

	rules := map[string]int{
		"name":    256,
		"email":   512,
		"faculty": 128,
		"major":   256,
	}

	if err := helper.ValidateFieldLengths(data, rules); err != nil {
		return err
	}

	if data.Role == model.RoleAdmin && requesterRole != model.RoleSuperAdmin {
		return fmt.Errorf("you cant access this role")
	}
	if data.Role == model.RoleSuperAdmin {
		return fmt.Errorf("you cant access this role")
	}

	switch data.Role {
	case model.RoleLecturer:
		if data.Nip == "" || data.Nidn == "" {
			return fmt.Errorf("lecturer must have both NIP and NIDN")
		}
	case model.RoleUser:
		if data.Nim == "" {
			return fmt.Errorf("user must have NIM")
		}
	case model.RoleAdmin:
	default:
		return fmt.Errorf("invalid role: %s", data.Role)
	}

	finalAcademicYear := ""
	if data.Role != model.RoleLecturer {
		finalAcademicYear = data.AcademicYear
	}

	finalUsername := data.Username
	if data.Role == model.RoleLecturer || data.Role == model.RoleUser {
		finalUsername = ""
	}

	registerCred := model.User{
		Name:         data.Name,
		Email:        helper.BindAndConvertToPtr(data.Email),
		Password:     string(hashedPassword),
		Major:        data.Major,
		Faculty:      data.Faculty,
		AcademicYear: finalAcademicYear,
		Nim:          helper.BindAndConvertToPtr(data.Nim),
		Nip:          helper.BindAndConvertToPtr(data.Nip),
		Nidn:         helper.BindAndConvertToPtr(data.Nidn),
		Username:     helper.BindAndConvertToPtr(finalUsername),
		Role:         data.Role,
	}

	_, err = s.repo.Register(ctx, registerCred)
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

	if err := helper.ValidateFieldLengths(cred, rules); err != nil {
		return nil, "", "", err
	}

	switch loginType {
	case "nidn":
		data, err = s.repo.GetByNidn(ctx, loginId)
		if err != nil {
			return nil, "", "", fmt.Errorf("user not found")
		}
		if data.Role != model.RoleLecturer {
			return nil, "", "", fmt.Errorf("you cant login use nidn %s", err)
		}
	case "nim":
		data, err = s.repo.GetByNim(ctx, loginId)
		if err != nil {
			return nil, "", "", fmt.Errorf("user not found")
		}
		if data.Role != model.RoleUser {
			return nil, "", "", fmt.Errorf("you cant login use nim %s", err)
		}
	case "username":
		data, err = s.repo.GetByUsn(ctx, loginId)
		if err != nil {
			return nil, "", "", fmt.Errorf("user not found")
		}
		if data.Role != model.RoleAdmin && data.Role != model.RoleSuperAdmin {
			return nil, "", "", fmt.Errorf("you cant login use username %s", err)
		}
	default:
		return nil, "", "", fmt.Errorf("user not found")
	}

	if err != nil || data == nil {
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

	if data.Role == model.RoleSuperAdmin {
		return nil, fmt.Errorf("user not found")
	}

	return data, nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	data, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("user with email %s not found: %w", email, err)
	}

	if data.Role == model.RoleSuperAdmin {
		return nil, fmt.Errorf("user not found")
	}

	return data, nil
}

func (s *userService) Update(ctx context.Context, c *gin.Context, data model.UpdateUser, id int, requesterRole model.Role, currentId int) (*model.User, error) {
	oldUser, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	effectiveRole := oldUser.Role
	if data.Role != nil {
		effectiveRole = *data.Role
	}

	if err := update.ValidateAuthorization(effectiveRole, oldUser, data, requesterRole, currentId); err != nil {
		return nil, err
	}

	update.NormalizeRoleTransition(oldUser, &data, effectiveRole)

	if err := update.ValidateRoleRequirements(data, effectiveRole); err != nil {
		return nil, err
	}

	if err := update.ValidateRoleTransitionRequirements(oldUser, data, effectiveRole); err != nil {
		return nil, err
	}

	update.MergeDefaults(oldUser, &data, effectiveRole)

	if err := update.HandleUserImageUpload(c, oldUser, &data, id); err != nil {
		return nil, err
	}

	updatedUser, err := s.repo.Update(ctx, data, id)
	if err != nil {
		return nil, update.FormatUpdateUserError(err, data)
	}

	return updatedUser, nil
}

func (s *userService) Delete(ctx context.Context, id int, requesterRole model.Role) error {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	if user.Role == model.RoleSuperAdmin && requesterRole != model.RoleSuperAdmin {
		return fmt.Errorf("user not found")
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
		return nil, 0, fmt.Errorf("failed to get users: %w", err)
	}

	filtered := make([]model.User, 0)
	hiddenCount := int64(0)

	for _, u := range dataList {
		if u.Role == model.RoleSuperAdmin {
			hiddenCount++
			continue
		}
		filtered = append(filtered, u)
	}

	totalWithoutSA := total - hiddenCount

	return filtered, totalWithoutSA, nil
}

func (s *userService) GetByNim(ctx context.Context, nim string, requesterRole model.Role) (*model.User, error) {
	if len(nim) > 9 {
		return nil, fmt.Errorf("nim cannot be more than 9 characters: %s", nim)
	}

	data, err := s.repo.GetByNim(ctx, nim)
	if err != nil {
		return nil, fmt.Errorf("user with nim %q not found: %w", nim, err)
	}

	if data.Role == model.RoleSuperAdmin && requesterRole != model.RoleSuperAdmin {
		return nil, fmt.Errorf("user not found")
	}

	return data, nil
}

func (s *userService) GetByNidn(ctx context.Context, nidn string, requesterRole model.Role) (*model.User, error) {
	if len(nidn) > 10 {
		return nil, fmt.Errorf("nidn cannot be more than 9 characters: %s", nidn)
	}

	data, err := s.repo.GetByNidn(ctx, nidn)
	if err != nil {
		return nil, fmt.Errorf("user with nidn %q not found: %w", nidn, err)
	}

	if data.Role == model.RoleSuperAdmin && requesterRole != model.RoleSuperAdmin {
		return nil, fmt.Errorf("user not found")
	}

	return data, nil
}

func (s *userService) GetByUsn(ctx context.Context, username string, requesterRole model.Role) (*model.User, error) {
	if len(username) > 256 {
		return nil, fmt.Errorf("username cannot be more than 256 characters: %s", username)
	}

	data, err := s.repo.GetByUsn(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("user with username %q not found: %w", username, err)
	}

	if data.Role == model.RoleSuperAdmin && requesterRole != model.RoleSuperAdmin {
		return nil, fmt.Errorf("user not found")
	}

	return data, nil
}

func (s *userService) GetByName(ctx context.Context, name string, limit int, offset int) ([]model.User, int64, error) {
	if helper.ContainsNumber(name) {
		return nil, 0, fmt.Errorf("name cannot contain numbers")
	}

	dataList, total, err := s.repo.GetByName(ctx, name, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("user with name %q not found: %w", name, err)
	}

	filtered := make([]model.User, 0)
	hiddenCount := int64(0)

	for _, u := range dataList {
		if u.Role == model.RoleSuperAdmin {
			hiddenCount++
			continue
		}
		filtered = append(filtered, u)
	}

	totalWithoutSA := total - hiddenCount

	return filtered, totalWithoutSA, nil
}

func (s *userService) GetByRole(ctx context.Context, role string, limit int, offset int, requesterRole string) ([]model.User, int64, error) {
	modelRole := model.Role(role)
	requesterRoleModel := model.Role(requesterRole)

	if modelRole == model.RoleSuperAdmin && requesterRoleModel != model.RoleSuperAdmin {
		return nil, 0, fmt.Errorf("user not found")
	}

	if modelRole != model.RoleAdmin && modelRole != model.RoleUser && modelRole != model.RoleLecturer {
		return nil, 0, fmt.Errorf("invalid role: %s", role)
	}

	dataList, total, err := s.repo.GetByRole(ctx, modelRole, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("user with role %q not found: %w", role, err)
	}

	return dataList, total, nil
}

func (s *userService) ChangePassword(ctx context.Context, id int, newPassword string, role model.Role) error {
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

	if role == model.RoleAdmin && user.Role == model.RoleSuperAdmin {
		return fmt.Errorf("admin cannot change super admin role")
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

func (s *userService) ChangeRole(ctx context.Context, id int, role model.Role, requesterRole model.Role) error {
	if role == model.RoleAdmin && model.Role(requesterRole) != model.RoleSuperAdmin {
		return fmt.Errorf("you dont have permission to assign admin role")
	}

	if err := s.repo.ChangeRole(ctx, id, role); err != nil {
		return fmt.Errorf("cannot change role: %w", err)
	}
	return nil
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
