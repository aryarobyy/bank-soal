package repository

import (
	"context"
	"fmt"

	"latih.in-be/internal/model"

	"gorm.io/gorm"
)

type ExamSessionRepository interface {
	Create(ctx context.Context, e model.ExamSession) error
	GetById(ctx context.Context, id int) (*model.ExamSession, error)
	Update(ctx context.Context, id int, e model.UpdateExamSession) (*model.ExamSession, error)
	Delete(ctx context.Context, id int) error
	GetMany(ctx context.Context, userId int, limit int, offset int) ([]model.ExamSession, error)
	UpdateCurrNo(ctx context.Context, id int, no model.UpdateCurrNo) (*model.ExamSession, error)
	FinishExam(ctx context.Context, id int, e model.FinishExam) (*model.ExamSession, error)
	CheckUserSession(ctx context.Context, userId int, examId int) error
}

type examSessionRepository struct {
	db *gorm.DB
}

func NewExamSessionRepository(db *gorm.DB) ExamSessionRepository {
	return &examSessionRepository{db: db}
}

func (r *examSessionRepository) Create(ctx context.Context, e model.ExamSession) error {
	if err := r.db.WithContext(ctx).Create(&e).Error; err != nil {
		return err
	}
	return nil
}

func (r *examSessionRepository) GetById(ctx context.Context, id int) (*model.ExamSession, error) {
	e := model.ExamSession{}
	if err := r.db.WithContext(ctx).
		First(&e, id).
		Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *examSessionRepository) Update(ctx context.Context, id int, e model.UpdateExamSession) (*model.ExamSession, error) {
	updateData := map[string]interface{}{}

	if e.UserId != 0 {
		updateData["user_id"] = e.UserId
	}
	if e.ExamId != 0 {
		updateData["exam_id"] = e.ExamId
	}
	if !e.StartedAt.IsZero() {
		updateData["started_at"] = e.StartedAt
	}
	if e.FinishedAt != nil {
		updateData["finished_at"] = e.FinishedAt
	}
	if e.Status != "" {
		updateData["status"] = e.Status
	}
	if e.CurrentNo != 0 {
		updateData["current_no"] = e.CurrentNo
	}
	if e.Score != nil {
		updateData["score"] = e.Score
	}

	if len(updateData) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	if err := r.db.WithContext(ctx).
		Model(&model.ExamSession{}).
		Where("id = ?", id).
		Updates(updateData).Error; err != nil {
		return nil, err
	}

	var updated model.ExamSession
	if err := r.db.WithContext(ctx).
		First(&updated, id).
		Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

func (r *examSessionRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&model.ExamSession{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *examSessionRepository) GetMany(ctx context.Context, userId int, limit int, offset int) ([]model.ExamSession, error) {
	var sessions []model.ExamSession
	if err := r.db.WithContext(ctx).
		Where("user_id = ?", userId).
		Limit(limit).
		Offset(offset).
		Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *examSessionRepository) UpdateCurrNo(ctx context.Context, id int, no model.UpdateCurrNo) (*model.ExamSession, error) {
	if no.CurrentNo <= 0 {
		return nil, fmt.Errorf("invalid current number")
	}

	if err := r.db.WithContext(ctx).
		Model(&model.ExamSession{}).
		Where("id = ?", id).
		Update("current_no = ?", no.CurrentNo).Error; err != nil {
		return nil, err
	}

	var updated model.ExamSession
	if err := r.db.WithContext(ctx).
		First(&updated, id).
		Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

func (r *examSessionRepository) FinishExam(ctx context.Context, id int, e model.FinishExam) (*model.ExamSession, error) {
	if *e.Score == 0.0 {
		return nil, fmt.Errorf("invalid status")
	}

	if err := r.db.WithContext(ctx).
		Model(&model.ExamSession{}).
		Where("id = ?", id).
		Update("score", e.Score).Error; err != nil {
		return nil, err
	}

	var updated model.ExamSession
	if err := r.db.WithContext(ctx).
		First(&updated, id).
		Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

func (r *examSessionRepository) CheckUserSession(ctx context.Context, userId int, examId int) error {
	var session model.ExamSession

	err := r.db.WithContext(ctx).
		Where("user_id = ? AND exam_id = ? AND status = ?", userId, examId, model.SessionInProgress).
		First(&session).Error

	if err == nil {
		return fmt.Errorf("user already has an active session")
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}
