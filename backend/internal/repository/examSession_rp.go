package repository

import (
	"context"
	"fmt"

	"latih.in-be/internal/model"

	"gorm.io/gorm"
)

type ExamSessionRepository interface {
	Create(ctx context.Context, e model.ExamSession) (*model.ExamSession, error)
	GetById(ctx context.Context, id int) (*model.ExamSession, error)
	Update(ctx context.Context, id int, e model.UpdateExamSession) (*model.ExamSession, error)
	Delete(ctx context.Context, id int) error
	GetMany(ctx context.Context, examId int, limit int, offset int) ([]model.ExamSession, int64, error)
	UpdateCurrNo(ctx context.Context, id int, no model.UpdateCurrNo) (*model.ExamSession, error)
	FinishExam(ctx context.Context, id int, e model.FinishExam) (*model.ExamSession, error)
	CheckUserSession(ctx context.Context, userId int, examId int) (*model.ExamSession, error)
	GetScore(ctx context.Context, sessionId int, userId int) (*model.ExamSession, error)
	GetUserSession(ctx context.Context, userId int, limit int, offset int) ([]model.ExamSession, int64, error)
}

type examSessionRepository struct {
	db *gorm.DB
}

func NewExamSessionRepository(db *gorm.DB) ExamSessionRepository {
	return &examSessionRepository{db: db}
}

func (r *examSessionRepository) Create(ctx context.Context, e model.ExamSession) (*model.ExamSession, error) {
	if err := r.db.WithContext(ctx).
		Create(&e).
		Error; err != nil {
		return nil, err
	}

	return &e, nil
}

func (r *examSessionRepository) GetById(ctx context.Context, id int) (*model.ExamSession, error) {
	e := model.ExamSession{}
	if err := r.db.WithContext(ctx).
		Preload("UserAnswers").
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
	if e.Score != 0 {
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
	if err := r.db.WithContext(ctx).
		Model(&model.ExamSession{}).
		Where("id = ?", id).
		Delete(id).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *examSessionRepository) GetMany(ctx context.Context, examId int, limit int, offset int) ([]model.ExamSession, int64, error) {
	var (
		sessions []model.ExamSession
		total    int64
	)
	query := r.db.WithContext(ctx).
		Model(model.ExamSession{}).
		Where("exam_id = ?", examId).
		Limit(limit).
		Offset(offset)

	if err := query.
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Find(&sessions).
		Error; err != nil {
		return nil, 0, err
	}
	return sessions, total, nil
}

func (r *examSessionRepository) UpdateCurrNo(ctx context.Context, id int, no model.UpdateCurrNo) (*model.ExamSession, error) {
	if no.CurrentNo <= 0 {
		return nil, fmt.Errorf("invalid: no cant be 0")
	}

	if err := r.db.WithContext(ctx).
		Model(&model.ExamSession{}).
		Where("id = ?", id).
		Update("current_no", no.CurrentNo).Error; err != nil {
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
	if err := r.db.WithContext(ctx).
		Model(&model.ExamSession{}).
		Where("id = ?", id).
		Updates(&e).Error; err != nil {
		return nil, err
	}

	var updated model.ExamSession
	if err := r.db.WithContext(ctx).
		First(&updated, id).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

func (r *examSessionRepository) CheckUserSession(ctx context.Context, userId int, examId int) (*model.ExamSession, error) {
	var session model.ExamSession

	err := r.db.WithContext(ctx).
		Where("user_id = ? AND exam_id = ? AND status = ?", userId, examId, model.SessionInProgress).
		First(&session).Error

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *examSessionRepository) GetScore(ctx context.Context, sessionId int, userId int) (*model.ExamSession, error) {
	session := model.ExamSession{}

	if err := r.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", sessionId, userId).
		Select("max_score", "score", "percentage").
		First(&session).
		Error; err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *examSessionRepository) GetUserSession(ctx context.Context, userId int, limit int, offset int) ([]model.ExamSession, int64, error) {
	var (
		sessions []model.ExamSession
		total    int64
	)

	if err := r.db.WithContext(ctx).
		Model(model.ExamSession{}).
		Where("user_id = ?", userId).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).
		Model(model.ExamSession{}).
		Where("user_id = ?", userId).
		Find(&sessions).
		Error; err != nil {
		return nil, 0, err
	}

	return sessions, total, nil
}
