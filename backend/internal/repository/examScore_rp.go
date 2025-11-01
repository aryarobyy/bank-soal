package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"latih.in-be/internal/model"
)

type ExamScoreRepository interface {
	Create(ctx context.Context, e model.ExamScore) error
	GetById(ctx context.Context, id int) (*model.ExamScore, error)
	GetMany(ctx context.Context, examId int, limit int, offset int) ([]model.ExamScore, error)
	Update(ctx context.Context, e model.ExamScore, id int) (*model.ExamScore, error)
	Delete(ctx context.Context, id int) error
}

type examScoreRepository struct {
	db *gorm.DB
}

func NewExamScoreRepository(db *gorm.DB) ExamScoreRepository {
	return &examScoreRepository{db: db}
}

func (r *examScoreRepository) Create(ctx context.Context, e model.ExamScore) error {
	if err := r.db.WithContext(ctx).
		Create(&e).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *examScoreRepository) GetById(ctx context.Context, id int) (*model.ExamScore, error) {
	e := model.ExamScore{}

	if err := r.db.WithContext(ctx).
		First(&e, id).
		Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *examScoreRepository) GetMany(ctx context.Context, examId int, limit int, offset int) ([]model.ExamScore, error) {
	var e []model.ExamScore

	if err := r.db.WithContext(ctx).
		Where("exam_id = ?", examId).
		Limit(limit).
		Offset(offset).
		Error; err != nil {
		return nil, err
	}
	return e, nil
}

func (r *examScoreRepository) Update(ctx context.Context, e model.ExamScore, id int) (*model.ExamScore, error) {
	updateData := map[string]interface{}{}

	if e.ExamId != 0 {
		updateData["exam_id"] = e.ExamId
	}
	if e.Status != "" {
		updateData["status"] = e.Status
	}
	if e.TotalScore != 0 {
		updateData["total_score"] = e.TotalScore
	}
	if e.UserId != 0 {
		updateData["user_id"] = e.UserId
	}
	if len(updateData) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	if err := r.db.WithContext(ctx).
		Model(&model.ExamScore{}).
		Where("id = ?", id).
		Updates(updateData).Error; err != nil {
		return nil, err
	}

	var updated model.ExamScore
	if err := r.db.WithContext(ctx).First(&updated, id).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *examScoreRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).
		Model(model.Exam{}).
		Where("id = ?", id).
		Delete(id).
		Error; err != nil {
		return err
	}
	return nil
}
