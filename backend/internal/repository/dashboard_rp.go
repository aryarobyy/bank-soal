package repository

import (
	"context"

	"gorm.io/gorm"
	"latih.in-be/internal/model"
)

type DashboardSummary struct {
	UserCount int64 `json:"user_count"`
	// LecturerCount int64 `json:"lecturer_count"`
	// StudentCount  int64 `json:"student_count"`
	ExamCount     int64 `json:"exam_count"`
	SessionCount  int64 `json:"session_count"`
	QuestionCount int64 `json:"question_count"`
	SubjectCount  int64 `json:"subject_count"`
}

type DashboardRepository interface {
	GetSummary(ctx context.Context) (*DashboardSummary, error)
}

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{db: db}
}

func (r *dashboardRepository) GetSummary(ctx context.Context) (*DashboardSummary, error) {
	summ := &DashboardSummary{}

	if err := r.db.WithContext(ctx).
		Model(&model.User{}).
		Count(&summ.UserCount).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).
		Model(&model.Exam{}).
		Count(&summ.ExamCount).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).
		Model(&model.ExamSession{}).
		Count(&summ.SessionCount).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).
		Model(&model.Question{}).
		Count(&summ.QuestionCount).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).
		Model(&model.Subject{}).
		Count(&summ.SubjectCount).Error; err != nil {
		return nil, err
	}

	return summ, nil
}
