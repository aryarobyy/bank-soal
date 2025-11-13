package repository

import (
	"context"

	"gorm.io/gorm"
	"latih.in-be/internal/model"
)

type SubjectRepository interface {
	Create(ctx context.Context, s model.Subject) error
	GetById(ctx context.Context, id int) (*model.Subject, error)
	GetMany(ctx context.Context, limit int, offset int) ([]model.Subject, int64, error)
	Update(ctx context.Context, s model.Subject, id int) (*model.Subject, error)
	Delete(ctx context.Context, id int) error
	GetByCode(ctx context.Context, code string) (*model.Subject, error)
}

type subjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) SubjectRepository {
	return &subjectRepository{db: db}
}

func (r *subjectRepository) Create(ctx context.Context, s model.Subject) error {
	if err := r.db.WithContext(ctx).
		Create(&s).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *subjectRepository) GetById(ctx context.Context, id int) (*model.Subject, error) {
	s := model.Subject{}

	if err := r.db.
		WithContext(ctx).
		First(&s, id).
		Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *subjectRepository) GetByCode(ctx context.Context, code string) (*model.Subject, error) {
	s := model.Subject{}

	if err := r.db.
		WithContext(ctx).
		Where("code = ?", code).
		Find(&s).
		Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *subjectRepository) GetMany(ctx context.Context, limit int, offset int) ([]model.Subject, int64, error) {
	var (
		s     []model.Subject
		total int64
	)

	if err := r.db.WithContext(ctx).
		Model(&model.Subject{}).
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).
		Model(model.Subject{}).
		Limit(limit).
		Offset(offset).
		Find(&s).
		Error; err != nil {
		return nil, 0, err
	}
	return s, total, nil
}

func (r *subjectRepository) Update(ctx context.Context, data model.Subject, id int) (*model.Subject, error) {
	var subject model.Subject

	if err := r.db.WithContext(ctx).First(&subject, id).Error; err != nil {
		return nil, err
	}

	if data.Title != "" {
		subject.Title = data.Title
	}
	if data.Code != "" {
		subject.Code = data.Code
	}

	if err := r.db.WithContext(ctx).Save(&subject).Error; err != nil {
		return nil, err
	}

	return &subject, nil
}

func (r *subjectRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).
		Model(model.Subject{}).
		Where("id = ?", id).
		Delete(id).
		Error; err != nil {
		return err
	}
	return nil
}
