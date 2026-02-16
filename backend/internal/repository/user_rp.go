package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"latih.in-be/internal/model"
)

type UserRepository interface {
	Register(ctx context.Context, user model.User) (*model.User, error)
	GetById(ctx context.Context, id int) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, user model.UpdateUser, id int) (*model.User, error)
	Delete(ctx context.Context, id int) error
	GetMany(ctx context.Context, limit int, offset int) ([]model.User, int64, error)
	GetByNim(ctx context.Context, nim string) (*model.User, error)
	GetByNip(ctx context.Context, nip string) (*model.User, error)
	GetByUsn(ctx context.Context, username string) (*model.User, error)
	GetByName(ctx context.Context, name string, limit int, offset int) ([]model.User, int64, error)
	GetByRole(ctx context.Context, role model.Role, limit int, offset int) ([]model.User, int64, error)
	ChangePassword(ctx context.Context, id int, password string) error
	ChangeRole(ctx context.Context, id int, data model.User) error
	BulkInsert(ctx context.Context, users []model.User) ([]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Register(ctx context.Context, user model.User) (*model.User, error) {
	if err := r.db.
		WithContext(ctx).
		Create(&user).
		Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetById(ctx context.Context, id int) (*model.User, error) {
	user := model.User{}

	if err := r.db.
		WithContext(ctx).
		Model(model.User{}).
		First(&user, id).
		Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user := model.User{}
	if err := r.db.
		WithContext(ctx).
		Where("email = ?", email).
		First(&user).
		Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user model.UpdateUser, id int) (*model.User, error) {
	updateData := map[string]interface{}{}

	if user.Name != nil {
		updateData["name"] = user.Name
	}
	if user.Username != nil {
		updateData["username"] = user.Username
	}
	if user.Nim != nil {
		if *user.Nim != "" {
			updateData["nim"] = user.Nim
		} else {
			updateData["nim"] = nil
		}
	}
	if user.Nip != nil {
		if *user.Nip != "" {
			updateData["nip"] = user.Nip
		} else {
			updateData["nip"] = nil
		}
	}
	if user.Major != nil {
		updateData["major"] = user.Major
	}
	if user.Email != nil {
		updateData["email"] = user.Email
	}
	if user.Faculty != nil {
		updateData["faculty"] = user.Faculty
	}
	if user.AcademicYear != nil {
		updateData["academic_year"] = user.AcademicYear
	}
	if user.Status != nil {
		updateData["status"] = user.Status
	}
	if user.ImgUrl != nil {
		updateData["img_url"] = user.ImgUrl
	}
	if user.Role != nil {
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
	if err := r.db.
		WithContext(ctx).
		Model(model.User{}).
		First(&updated, id).
		Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.
		WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", id).
		Delete(id).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetMany(ctx context.Context, limit int, offset int) ([]model.User, int64, error) {
	var (
		users []model.User
		total int64
	)
	if err := r.db.
		WithContext(ctx).
		Model(&model.User{}).
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.
		WithContext(ctx).
		Model(model.User{}).
		Limit(limit).
		Offset(offset).
		Find(&users).
		Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (r *userRepository) GetByNim(ctx context.Context, nim string) (*model.User, error) {
	user := model.User{}
	if err := r.db.
		WithContext(ctx).
		Model(model.User{}).
		Where("nim = ?", nim).
		First(&user).
		Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByNip(ctx context.Context, nip string) (*model.User, error) {
	user := model.User{}
	if err := r.db.
		WithContext(ctx).
		Model(model.User{}).
		Where("nip = ?", nip).
		First(&user).
		Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByUsn(ctx context.Context, username string) (*model.User, error) {
	user := model.User{}
	if err := r.db.
		WithContext(ctx).
		Model(model.User{}).
		Where("username = ?", username).
		First(&user).
		Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByName(ctx context.Context, name string, limit int, offset int) ([]model.User, int64, error) {
	var (
		users []model.User
		total int64
	)

	query := r.db.
		WithContext(ctx).
		Model(&model.User{}).
		Where("name LIKE ?", "%"+name+"%")

	if err := query.
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Limit(limit).
		Offset(offset).
		Find(&users).
		Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (r *userRepository) GetByRole(ctx context.Context, role model.Role, limit int, offset int) ([]model.User, int64, error) {
	var (
		users []model.User
		total int64
	)

	query := r.db.
		WithContext(ctx).
		Model(model.User{}).
		Where("role = ?", role)

	if err := query.
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Limit(limit).
		Offset(offset).
		Find(&users).
		Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (r *userRepository) ChangePassword(ctx context.Context, id int, password string) error {
	if err := r.db.
		WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", id).
		Update("password", password).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) ChangeRole(ctx context.Context, id int, data model.User) error {
	updates := map[string]interface{}{
		"role": data.Role,
	}

	if data.Nim != nil {
		updates["nim"] = data.Nim
	} else {
		updates["nim"] = nil
	}

	if data.Nip != nil {
		updates["nip"] = data.Nip
	} else {
		updates["nip"] = nil
	}

	if data.AcademicYear != "" {
		updates["academic_year"] = data.AcademicYear
	}

	if data.Username != nil {
		updates["username"] = data.Username
	} else {
		updates["username"] = nil
	}

	return r.db.
		WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", id).
		Updates(updates).
		Error
}

func (r *userRepository) BulkInsert(ctx context.Context, users []model.User) ([]model.User, error) {
	if err := r.db.
		WithContext(ctx).
		CreateInBatches(&users, 100). //batch size 100
		Error; err != nil {
		return nil, err
	}

	return users, nil
}
