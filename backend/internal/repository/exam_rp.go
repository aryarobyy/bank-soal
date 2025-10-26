package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"latih.in-be/internal/model"
)

type ExamRepository interface {
	Create(ctx context.Context, e model.Exam) error
	GetById(ctx context.Context, id int) (*model.Exam, error)
	Update(ctx context.Context, e model.Exam, id int) (*model.Exam, error)
	Delete(ctx context.Context, id int) error
	GetMany(ctx context.Context, limit int, offset int) ([]model.Exam, error)
}

type examRepository struct {
	db *gorm.DB
}

func NewExamRepository(db *gorm.DB) ExamRepository {
	return &examRepository{db: db}
}

func (r *examRepository) Create(ctx context.Context, e model.Exam) error {
	if err := r.db.WithContext(ctx).
		Create(&e).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *examRepository) GetById(ctx context.Context, id int) (*model.Exam, error) {
	e := model.Exam{}

	if err := r.db.WithContext(ctx).
		Model(model.Exam{}).
		First(&e, id).
		Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *examRepository) Update(ctx context.Context, e model.Exam, id int) (*model.Exam, error) {
	updateData := map[string]interface{}{}

	if e.CreatorId != 0 {
		updateData["creator_id"] = e.CreatorId
	}
	if e.Description != "" {
		updateData["description"] = e.Description
	}
	if e.Difficulty != "" {
		updateData["difficulty"] = e.Difficulty
	}
	if e.StartedAt == nil || e.StartedAt.IsZero() {
		updateData["started_at"] = e.StartedAt
	}
	if e.FinishedAt == nil || e.FinishedAt.IsZero() {
		updateData["finished_at"] = e.FinishedAt
	}
	if e.LongTime != 0 {
		updateData["long_time"] = e.LongTime
	}
	if e.Title != "" {
		updateData["title"] = e.Title
	}

	if len(updateData) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	if err := r.db.WithContext(ctx).
		Model(&model.Exam{}).
		Where("id = ?", id).
		Updates(updateData).Error; err != nil {
		return nil, err
	}

	var updated model.Exam
	if err := r.db.WithContext(ctx).
		First(&updated, id).
		Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *examRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).
		Model(model.Exam{}).
		Where("id = ?", id).
		Delete(id).Error; err != nil {
		return err
	}
	return nil
}

func (r *examRepository) GetMany(ctx context.Context, limit int, offset int) ([]model.Exam, error) {
	var e []model.Exam
	if err := r.db.WithContext(ctx).
		Find(&e).
		Limit(limit).
		Offset(offset).
		Error; err != nil {
		return nil, err
	}
	return e, nil
}
