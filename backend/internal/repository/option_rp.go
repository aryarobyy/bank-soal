package repository

import (
	"context"

	"gorm.io/gorm"
	"latih.in-be/internal/model"
)

type OptionRepository interface {
	Create(ctx context.Context, o model.Option) error
	GetById(ctx context.Context, id int) (*model.Option, error)
	GetMany(ctx context.Context, qId int, limit int, offset int) ([]model.Option, error)
	Update(ctx context.Context, o model.Option, id int) (*model.Option, error)
	Delete(ctx context.Context, id int) error
	DeleteByQuestionId(ctx context.Context, qId int) error
	CheckCorrectAnswer(ctx context.Context, qId int, answer string) (bool, error)
}

type optionRepository struct {
	db *gorm.DB
}

func NewOptionRepository(db *gorm.DB) OptionRepository {
	return &optionRepository{db: db}
}

func (r *optionRepository) Create(ctx context.Context, o model.Option) error {
	if err := r.db.WithContext(ctx).
		Create(&o).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *optionRepository) GetById(ctx context.Context, id int) (*model.Option, error) {
	o := model.Option{}

	if err := r.db.
		WithContext(ctx).
		First(&o, id).
		Error; err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *optionRepository) GetMany(ctx context.Context, qId int, limit int, offset int) ([]model.Option, error) {
	var o []model.Option

	if err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Where("question_id = ?", qId).
		Find(&o).
		Error; err != nil {
		return nil, err
	}
	return o, nil
}

func (r *optionRepository) Update(ctx context.Context, o model.Option, id int) (*model.Option, error) {
	if err := r.db.WithContext(ctx).
		Model(model.Option{}).
		Where("id = ?", id).
		Updates(o).Error; err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *optionRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).
		Model(model.Option{}).
		Where("id = ?", id).
		Delete(id).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *optionRepository) DeleteByQuestionId(ctx context.Context, qId int) error {
	if err := r.db.WithContext(ctx).
		Model(model.Option{}).
		Where("question_id = ?", qId).
		Delete(qId).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *optionRepository) CheckCorrectAnswer(ctx context.Context, qId int, answer string) (bool, error) {
	var option model.Option

	if err := r.db.WithContext(ctx).
		Select("is_correct").
		Where("question_id = ? AND option_label = ?", qId, answer).
		First(&option).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, err
		}
		return false, err
	}

	return option.IsCorrect, nil
}
