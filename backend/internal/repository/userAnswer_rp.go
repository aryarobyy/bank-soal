package repository

import (
	"context"
	"log"

	"gorm.io/gorm"
	"latih.in-be/internal/model"
)

type UserAnswerRepository interface {
	Create(ctx context.Context, userAnswer *model.UserAnswer) error
	GetById(ctx context.Context, id int) (*model.UserAnswer, error)
	GetMany(ctx context.Context, limit int, offset int) ([]model.UserAnswer, int64, error)
	Update(ctx context.Context, id int, userAnswer *model.UserAnswer) (*model.UserAnswer, error)
	Delete(ctx context.Context, id int) error
	GetByExamSessionId(ctx context.Context, examSessionId int, limit int, offset int) ([]model.UserAnswer, int64, error)
	GetByQuestionId(ctx context.Context, questionId int, limit int, offset int) ([]model.UserAnswer, int64, error)
	GetUserAnswer(ctx context.Context, userId int, examSessionId int, limit int, offset int) ([]model.UserAnswer, int64, error)
	GetAllUserAnswers(ctx context.Context, userId int, examSessionId int) ([]model.UserAnswer, error)
}

type userAnswerRepository struct {
	db *gorm.DB
}

func NewUserAnswerRepository(db *gorm.DB) UserAnswerRepository {
	return &userAnswerRepository{db: db}
}

func (r *userAnswerRepository) Create(ctx context.Context, userAnswer *model.UserAnswer) error {
	if err := r.db.WithContext(ctx).Create(&userAnswer).Error; err != nil {
		return err
	}
	return nil
}

func (r *userAnswerRepository) GetById(ctx context.Context, id int) (*model.UserAnswer, error) {
	userAnswer := model.UserAnswer{}
	err := r.db.WithContext(ctx).
		Preload("ExamSession").
		First(&userAnswer, id).
		Error
	if err != nil {
		return nil, err
	}
	return &userAnswer, nil
}

func (r *userAnswerRepository) GetMany(ctx context.Context, limit int, offset int) ([]model.UserAnswer, int64, error) {
	var (
		userAnswers []model.UserAnswer
		total       int64
	)

	query := r.db.
		WithContext(ctx).
		Model(&model.UserAnswer{})

	if err := query.
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Preload("ExamSession").
		Limit(limit).
		Offset(offset).
		Find(&userAnswers).Error; err != nil {
		return nil, 0, err
	}

	return userAnswers, total, nil
}

func (r *userAnswerRepository) Update(ctx context.Context, id int, userAnswer *model.UserAnswer) (*model.UserAnswer, error) {
	updateData := map[string]interface{}{}

	if userAnswer.UserId != 0 {
		updateData["user_id"] = userAnswer.UserId
	}
	if userAnswer.ExamSessionId != 0 {
		updateData["exam_session_id"] = userAnswer.ExamSessionId
	}
	if userAnswer.QuestionId != 0 {
		updateData["question_id"] = userAnswer.QuestionId
	}
	if userAnswer.Answer != "" {
		updateData["answer"] = userAnswer.Answer
	}

	updateData["is_correct"] = userAnswer.IsCorrect

	var updated model.UserAnswer
	if err := r.db.WithContext(ctx).
		Preload("ExamSession").
		First(&updated, id).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

func (r *userAnswerRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&model.UserAnswer{}).
		Error
}

func (r *userAnswerRepository) GetByExamSessionId(ctx context.Context, examSessionId int, limit int, offset int) ([]model.UserAnswer, int64, error) {
	var (
		userAnswers []model.UserAnswer
		total       int64
	)

	query := r.db.WithContext(ctx).
		Model(&model.UserAnswer{}).
		Where("exam_session_id = ?", examSessionId)

	query = query.Debug()

	if err := query.
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	log.Printf("Total records found for exam_session_id %d: %d", examSessionId, total)

	if err := query.
		Preload("ExamSession").
		Limit(limit).
		Offset(offset).
		Find(&userAnswers).Error; err != nil {
		return nil, 0, err
	}

	return userAnswers, total, nil
}

func (r *userAnswerRepository) GetByQuestionId(ctx context.Context, questionId int, limit int, offset int) ([]model.UserAnswer, int64, error) {
	var (
		userAnswers []model.UserAnswer
		total       int64
	)

	query := r.db.WithContext(ctx).
		Model(&model.UserAnswer{}).
		Where("question_id = ?", questionId)

	if err := query.
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := query.Preload("ExamSession").
		Limit(limit).
		Offset(offset).
		Find(&userAnswers).Error; err != nil {
		return nil, 0, err
	}

	return userAnswers, total, nil
}

func (r *userAnswerRepository) GetUserAnswer(ctx context.Context, userId int, examSessionId int, limit int, offset int) ([]model.UserAnswer, int64, error) {
	var (
		userAnswers []model.UserAnswer
		total       int64
	)

	query := r.db.WithContext(ctx).
		Model(&model.UserAnswer{}).
		Where("user_id = ? AND exam_session_id = ?", userId, examSessionId)

	if err := query.
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := query.Preload("ExamSession").
		Limit(limit).
		Offset(offset).
		Find(&userAnswers).Error; err != nil {
		return nil, 0, err
	}

	return userAnswers, total, nil
}

func (r *userAnswerRepository) GetAllUserAnswers(ctx context.Context, userId int, examSessionId int) ([]model.UserAnswer, error) {
	var userAnswers []model.UserAnswer

	if err := r.db.WithContext(ctx).
		Where("user_id = ? AND exam_session_id = ?", userId, examSessionId).
		Find(&userAnswers).Error; err != nil {
		return nil, err
	}

	return userAnswers, nil
}
