package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"latih.in-be/model"
)

type UserRepository interface {
	Register(ctx context.Context, user model.User) (*model.User, error)
	GetById(ctx context.Context, id int) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, user model.User, id int) (*model.User, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]model.User, error)
	GetByNim(ctx context.Context, nim string) (*model.User, error)
	GetByName(ctx context.Context, name string) ([]model.User, error)
	GetByRole(ctx context.Context, role string) ([]model.User, error)
	ChangePassword(ctx context.Context, id int, password string) error
	ChangeRole(ctx context.Context, id int, role model.Role) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Register(ctx context.Context, user model.User) (*model.User, error) {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetById(ctx context.Context, id int) (*model.User, error) {
	user := model.User{}

	if err := r.db.WithContext(ctx).Model(model.User{}).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user := model.User{}
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user model.User, id int) (*model.User, error) {
	updateData := map[string]interface{}{}

	if user.Name != "" {
		updateData["name"] = user.Name
	}
	if user.Nim != "" {
		updateData["nim"] = user.Nim
	}
	if user.Major != "" {
		updateData["major"] = user.Major
	}
	if user.Faculty != "" {
		updateData["faculty"] = user.Faculty
	}
	if user.AcademicYear != 0 {
		updateData["academic_year"] = user.AcademicYear
	}
	if user.Status != "" {
		updateData["status"] = user.Status
	}
	if user.ImgUrl != "" {
		updateData["img_url"] = user.ImgUrl
	}
	if user.Role != "" {
		updateData["role"] = user.Role
	}

	if len(updateData) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	if err := r.db.WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", id).
		Updates(updateData).Error; err != nil {
		return nil, err
	}

	var updated model.User
	if err := r.db.WithContext(ctx).Model(model.User{}).First(&updated, id).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(id).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	if err := r.db.WithContext(ctx).Model(model.User{}).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetByNim(ctx context.Context, nim string) (*model.User, error) {
	user := model.User{}
	if err := r.db.WithContext(ctx).Model(model.User{}).Where("nim = ?", nim).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByName(ctx context.Context, name string) ([]model.User, error) {
	var users []model.User
	if err := r.db.WithContext(ctx).Model(model.User{}).Where("name = ?", name).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetByRole(ctx context.Context, role string) ([]model.User, error) {
	var users []model.User
	if err := r.db.WithContext(ctx).Model(model.User{}).Where("role = ?", role).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) ChangePassword(ctx context.Context, id int, password string) error {
	if err := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("password", password).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) ChangeRole(ctx context.Context, id int, role model.Role) error {
	if err := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("role", role).Error; err != nil {
		return err
	}
	return nil
}
