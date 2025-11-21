package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"latih.in-be/internal/model"
)

type QuestionRepository interface {
	Create(ctx context.Context, q *model.Question) error
	GetById(ctx context.Context, id int) (*model.Question, error)
	GetMany(ctx context.Context, limit int, offset int) ([]model.Question, int64, error)
	GetByExam(ctx context.Context, examId int, limit int, offset int) ([]model.Question, int64, error)
	Update(ctx context.Context, q model.UpdateQuestion, id int) (*model.Question, error)
	Delete(ctx context.Context, id int) error
	CreateWithOptions(ctx context.Context, question model.Question) error
	CreateBatch(ctx context.Context, q []model.Question) error
	GetByDifficult(ctx context.Context, diff string, limit int, offset int) ([]model.Question, int64, error)
	GetByCreatorId(ctx context.Context, creatorId int, limit int, offset int) ([]model.Question, int64, error)
	GetBySubject(ctx context.Context, subjectId int, limit int, offset int) ([]model.Question, int64, error)
}

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &questionRepository{db: db}
}

func (r *questionRepository) Create(ctx context.Context, q *model.Question) error {
	if err := r.db.WithContext(ctx).Create(q).Error; err != nil {
		return err
	}

	if err := r.db.WithContext(ctx).
		Preload("Subject").
		First(q, q.Id).Error; err != nil {
		return err
	}

	return nil
}

func (r *questionRepository) GetById(ctx context.Context, id int) (*model.Question, error) {
	var q model.Question
	if err := r.db.WithContext(ctx).
		Preload("Subject").
		Preload("Options").
		First(&q, id).
		Error; err != nil {
		return nil, err
	}
	return &q, nil
}

func (r *questionRepository) GetMany(ctx context.Context, limit int, offset int) ([]model.Question, int64, error) {
	var (
		q     []model.Question
		total int64
	)

	if err := r.db.WithContext(ctx).
		Model(&model.Question{}).
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).
		Model(&model.Question{}).
		Preload("Subject").
		Preload("Options").
		Limit(limit).
		Offset(offset).
		Find(&q).
		Error; err != nil {
		return nil, 0, err
	}
	return q, total, nil
}

func (r *questionRepository) GetByExam(ctx context.Context, examId int, limit int, offset int) ([]model.Question, int64, error) {
	var (
		q     []model.Question
		total int64
	)

	baseQuery := r.db.WithContext(ctx).
		Model(&model.Question{}).
		Joins("JOIN exam_questions eq ON eq.question_id = questions.id").
		Where("eq.exam_id = ?", examId)

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := baseQuery.
		Preload("Subject").
		Preload("Options").
		Limit(limit).
		Offset(offset).
		Find(&q).Error; err != nil {
		return nil, 0, err
	}

	return q, total, nil
}

func (r *questionRepository) Update(ctx context.Context, q model.UpdateQuestion, id int) (*model.Question, error) {
	updateData := map[string]interface{}{}

	if q.SubjectId != nil {
		updateData["subject_id"] = *q.SubjectId
	}
	if q.Answer != nil {
		updateData["answer"] = *q.Answer
	}
	if q.CreatorId != nil {
		updateData["creator_id"] = *q.CreatorId
	}
	if q.QuestionText != nil {
		updateData["question_text"] = *q.QuestionText
	}
	if q.Difficulty != nil {
		updateData["difficulty"] = *q.Difficulty
	}
	if q.Score != nil {
		updateData["score"] = *q.Score
	}
	if q.ImgUrl != nil {
		updateData["img_url"] = *q.ImgUrl
	}

	if len(updateData) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	if err := r.db.WithContext(ctx).
		Model(&model.Question{}).
		Where("id = ?", id).
		Updates(updateData).
		Error; err != nil {
		return nil, err
	}

	var updatedQuestion model.Question
	if err := r.db.WithContext(ctx).
		Preload("Subject").
		Preload("Options").
		First(&updatedQuestion, id).
		Error; err != nil {
		return nil, err
	}

	return &updatedQuestion, nil
}

func (r *questionRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&model.Question{}).
		Delete(id).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *questionRepository) CreateWithOptions(ctx context.Context, question model.Question) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&question).Error; err != nil {
			return err
		}

		for i := range question.Options {
			question.Options[i].QuestionId = question.Id
		}

		if len(question.Options) > 0 {
			if err := tx.Create(&question.Options).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *questionRepository) CreateBatch(ctx context.Context, q []model.Question) error {
	if err := r.db.
		WithContext(ctx).
		Create(&q).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *questionRepository) GetByDifficult(ctx context.Context, diff string, limit int, offset int) ([]model.Question, int64, error) {
	var (
		q     []model.Question
		total int64
	)

	query := r.db.WithContext(ctx).
		Model(&model.Question{}).
		Where("difficulty = ?", diff)

	if err := query.
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Preload("Subject").
		Preload("Options").
		Limit(limit).
		Offset(offset).
		Find(&q).
		Error; err != nil {
		return nil, 0, err
	}
	return q, total, nil
}

func (r *questionRepository) GetByCreatorId(ctx context.Context, creatorId int, limit int, offset int) ([]model.Question, int64, error) {
	var (
		q     []model.Question
		total int64
	)

	query := r.db.WithContext(ctx).
		Model(&model.Question{}).
		Where("creator_id = ?", creatorId)

	if err := query.
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Preload("Subject").
		Preload("Options").
		Limit(limit).
		Offset(offset).
		Find(&q).
		Error; err != nil {
		return nil, 0, err
	}
	return q, total, nil
}

func (r *questionRepository) GetBySubject(ctx context.Context, subjectId int, limit int, offset int) ([]model.Question, int64, error) {
	var (
		q     []model.Question
		total int64
	)

	query := r.db.WithContext(ctx).
		Model(&model.Question{}).
		Where("subject_id = ?", subjectId)

	if err := query.
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Preload("Subject").
		Preload("Options").
		Limit(limit).
		Offset(offset).
		Find(&q).
		Error; err != nil {
		return nil, 0, err
	}
	return q, total, nil
}
