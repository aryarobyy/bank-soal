package repository

import (
	"context"

	"gorm.io/gorm"
	"latih.in-be/internal/model"
)

type ScoreRepository interface {
	Create(ctx context.Context, o model.Score) error
	GetById(ctx context.Context, id int) (*model.Score, error)
	GetAll(ctx context.Context, qId int) ([]model.Score, error)
	Update(ctx context.Context, o model.Score, id int) (*model.Score, error)
	Delete(ctx context.Context, id int) error
}

type scoreRepository struct {
	db *gorm.DB
}

func NewScoreRepository(db *gorm.DB) ScoreRepository {
	return &scoreRepository{db: db}
}

func (r *scoreRepository) Create(ctx context.Context, s model.Score) error {
	if err := r.db.WithContext(ctx).Create(&s).Error; err != nil {
		return err
	}
	return nil
}

func (r *scoreRepository) GetById(ctx context.Context, id int) (*model.Score, error) {
	s := model.Score{}

	if err := r.db.WithContext(ctx).First(&s, id).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *scoreRepository) GetAll(ctx context.Context, qId int) ([]model.Score, error) {
	var s []model.Score

	if err := r.db.WithContext(ctx).Where("question_id = ?", qId).Error; err != nil {
		return nil, err
	}
	return s, nil
}

func (r *scoreRepository) Update(ctx context.Context, s model.Score, id int) (*model.Score, error) {
	if err := r.db.WithContext(ctx).Model(model.Score{}).Where("id = ?", id).Updates(s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *scoreRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Model(model.Exam{}).Where("id = ?", id).Delete(id).Error; err != nil {
		return err
	}
	return nil
}
