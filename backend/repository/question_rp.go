package repository

import (
	"context"

	"gorm.io/gorm"
	"latih.in-be/model"
)

type QuestionRepository interface {
	Create(ctx context.Context, q model.Question) error
	GetById(ctx context.Context, id int) (*model.Question, error)
	GetAll(ctx context.Context) ([]model.Question, error)
	GetByExam(ctx context.Context, examId int) ([]model.Question, error)
	Update(ctx context.Context, q model.Question, id int) (*model.Question, error)
	Delete(ctx context.Context, id int) error
	CreateWithOptions(ctx context.Context, question model.Question) error
	CreateBatch(ctx context.Context, q []model.Question) error
	GetByDifficult(ctx context.Context, diff string) ([]model.Question, error)
	GetByCreatorId(ctx context.Context, creatorId int) ([]model.Question, error)
}

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &questionRepository{db: db}
}

func (r *questionRepository) Create(ctx context.Context, q model.Question) error {
	if err := r.db.WithContext(ctx).Create(&q).Error; err != nil {
		return err
	}
	return nil
}

func (r *questionRepository) GetById(ctx context.Context, id int) (*model.Question, error) {
	q := model.Question{}

	if err := r.db.WithContext(ctx).First(&q, id).Error; err != nil {
		return nil, err
	}
	return &q, nil
}

func (r *questionRepository) GetAll(ctx context.Context) ([]model.Question, error) {
	var q []model.Question
	if err := r.db.WithContext(ctx).Find(&q).Error; err != nil {
		return nil, err
	}
	return q, nil
}

func (r *questionRepository) GetByExam(ctx context.Context, examId int) ([]model.Question, error) {
	var q []model.Question

	if err := r.db.WithContext(ctx).Model(model.Question{}).Where("exam_id = ?", examId).Find(&q).Error; err != nil {
		return nil, err
	}
	return q, nil
}

func (r *questionRepository) Update(ctx context.Context, q model.Question, id int) (*model.Question, error) {
	if err := r.db.WithContext(ctx).Model(model.Question{}).Where("id = ?", id).Updates(q).Error; err != nil {
		return nil, err
	}
	return &q, nil
}

func (r *questionRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Model(model.Exam{}).Where("id = ?", id).Delete(id).Error; err != nil {
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
	if err := r.db.WithContext(ctx).Create(&q).Error; err != nil {
		return err
	}
	return nil
}

func (r *questionRepository) GetByDifficult(ctx context.Context, diff string) ([]model.Question, error) {
	var q []model.Question

	if err := r.db.WithContext(ctx).Model(model.Question{}).Where("difficulty = ?", diff).Find(&q).Error; err != nil {
		return nil, err
	}
	return q, nil
}

func (r *questionRepository) GetByCreatorId(ctx context.Context, creatorId int) ([]model.Question, error) {
	var q []model.Question

	if err := r.db.WithContext(ctx).Model(model.Question{}).Where("creator_id = ?", creatorId).Find(&q).Error; err != nil {
		return nil, err
	}
	return q, nil
}
