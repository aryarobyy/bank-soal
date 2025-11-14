package repository

import (
	"context"

	"gorm.io/gorm"
	"latih.in-be/internal/model"
)

type XlsPathRepository interface {
	Create(ctx context.Context, o model.XlsPath) error
	GetById(ctx context.Context, id int) (*model.XlsPath, error)
	GetMany(ctx context.Context, limit int, offset int) ([]model.XlsPath, error)
	Delete(ctx context.Context, id int) error
}

type xlspathRepository struct {
	db *gorm.DB
}

func NewXlsPathRepository(db *gorm.DB) XlsPathRepository {
	return &xlspathRepository{db: db}
}

func (r *xlspathRepository) Create(ctx context.Context, o model.XlsPath) error {
	if err := r.db.WithContext(ctx).
		Create(&o).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *xlspathRepository) GetById(ctx context.Context, id int) (*model.XlsPath, error) {
	o := model.XlsPath{}

	if err := r.db.
		WithContext(ctx).
		First(&o, id).
		Error; err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *xlspathRepository) GetMany(ctx context.Context, limit int, offset int) ([]model.XlsPath, error) {
	var o []model.XlsPath

	if err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&o).
		Error; err != nil {
		return nil, err
	}
	return o, nil
}

func (r *xlspathRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).
		Model(model.XlsPath{}).
		Where("id = ?", id).
		Delete(id).
		Error; err != nil {
		return err
	}
	return nil
}
