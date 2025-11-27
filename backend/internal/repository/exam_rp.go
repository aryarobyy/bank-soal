package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"latih.in-be/internal/model"
)

type ExamRepository interface {
	Create(ctx context.Context, req model.CreateExam) error
	GetById(ctx context.Context, id int) (*model.Exam, error)
	Update(ctx context.Context, e model.Exam, id int) (*model.Exam, error)
	Delete(ctx context.Context, id int) error
	GetMany(ctx context.Context, limit int, offset int) ([]model.Exam, int64, error)
	StartSession(ctx context.Context, id int) (*model.Exam, error)
	UpdateScore(ctx context.Context, examId int, score int) error
	AddQuestions(ctx context.Context, examId int, questionIds []int) error
	ReplaceQuestions(ctx context.Context, examId int, questionIds []int) error
	RemoveQuestions(ctx context.Context, examId int, questionIds []int) error
}

type examRepository struct {
	db *gorm.DB
}

func NewExamRepository(db *gorm.DB) ExamRepository {
	return &examRepository{db: db}
}

func (r *examRepository) Create(ctx context.Context, req model.CreateExam) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		exam := model.Exam{
			Title:       req.Title,
			Description: req.Description,
			Difficulty:  model.Difficulty(req.Difficulty),
			LongTime:    req.LongTime,
			CreatorId:   req.CreatorId,
			StartedAt:   req.StartedAt,
			FinishedAt:  req.FinishedAt,
			Score:       req.Score,
		}

		if err := tx.Create(&exam).Error; err != nil {
			return err
		}

		if len(req.QuestionIds) > 0 {
			examQuestions := make([]model.ExamQuestion, 0, len(req.QuestionIds))
			for _, qid := range req.QuestionIds {
				examQuestions = append(examQuestions, model.ExamQuestion{
					ExamId:     exam.Id,
					QuestionId: qid,
				})
			}

			if err := tx.Create(&examQuestions).Error; err != nil {
				return err
			}
		}

		return nil
	})
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
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("exam_id = ?", id).Delete(&model.ExamQuestion{}).Error; err != nil {
			return fmt.Errorf("failed to delete exam questions: %w", err)
		}

		if err := tx.Where("id = ?", id).Delete(&model.Exam{}).Error; err != nil {
			return fmt.Errorf("failed to delete exam: %w", err)
		}
		return nil
	})
}

func (r *examRepository) GetMany(ctx context.Context, limit int, offset int) ([]model.Exam, int64, error) {
	var (
		e     []model.Exam
		total int64
	)

	if err := r.db.WithContext(ctx).
		Model(&model.Exam{}).
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).
		Model(&model.Exam{}).
		Limit(limit).
		Offset(offset).
		Find(&e).
		Error; err != nil {
		return nil, 0, err
	}
	return e, total, nil
}

func (r *examRepository) StartSession(ctx context.Context, id int) (*model.Exam, error) {
	var exam model.Exam
	if err := r.db.WithContext(ctx).
		Preload("Questions.Options").
		First(&exam, id).Error; err != nil {
		return nil, fmt.Errorf("failed to load exam with questions: %w", err)
	}
	return &exam, nil
}

func (r *examRepository) UpdateScore(ctx context.Context, examId int, score int) error {
	if err := r.db.WithContext(ctx).
		Model(&model.Exam{}).
		Where("id = ?", examId).
		Update("score", score).
		Error; err != nil {
		return fmt.Errorf("failed to update exam: %w", err)
	}
	return nil
}

func (r *examRepository) AddQuestions(ctx context.Context, examId int, questionIds []int) error {
	var examQuestions []model.ExamQuestion
	for _, qid := range questionIds {
		examQuestions = append(examQuestions, model.ExamQuestion{
			ExamId:     examId,
			QuestionId: qid,
		})
	}

	if err := r.db.WithContext(ctx).
		Create(&examQuestions).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *examRepository) ReplaceQuestions(ctx context.Context, examId int, questionIds []int) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("exam_id = ?", examId).Delete(&model.ExamQuestion{}).Error; err != nil {
			return err
		}

		if len(questionIds) > 0 {
			var examQuestions []model.ExamQuestion
			for _, qid := range questionIds {
				examQuestions = append(examQuestions, model.ExamQuestion{
					ExamId:     examId,
					QuestionId: qid,
				})
			}
			if err := tx.Create(&examQuestions).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *examRepository) RemoveQuestions(ctx context.Context, examId int, questionIds []int) error {
	if err := r.db.WithContext(ctx).
		Where("exam_id = ? AND question_id IN ?", examId, questionIds).
		Delete(&model.ExamQuestion{}).
		Error; err != nil {
		return err
	}
	return nil
}
