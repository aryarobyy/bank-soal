package service

import (
	"context"
	"fmt"

	"latih.in-be/internal/model"
	"latih.in-be/internal/repository"
)

type XlsPathService interface {
	SaveXlsPath(ctx context.Context, filepath string) error
	GetById(ctx context.Context, id int) (*model.XlsPath, error)
	Delete(ctx context.Context, id int) error
	GetMany(ctx context.Context, limit int, offset int) ([]model.XlsPath, error)
}

type xlspathService struct {
	repo repository.XlsPathRepository
}

func NewXlsPathService(repo repository.XlsPathRepository) XlsPathService {
	return &xlspathService{
		repo: repo,
	}
}

func (s *xlspathService) SaveXlsPath(ctx context.Context, filepath string) error {
	xls := model.XlsPath{
		FilePath: filepath,
	}

	err := s.repo.Create(ctx, xls)
	if err != nil {
		return fmt.Errorf("failed to create xlspath: %w", err)
	}

	return nil
}

func (s *xlspathService) GetById(ctx context.Context, id int) (*model.XlsPath, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("data with id %d not found: %w", id, err)
	}
	return data, nil
}

func (s *xlspathService) GetMany(ctx context.Context, limit int, offset int) ([]model.XlsPath, error) {
	data, err := s.repo.GetMany(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get all data: %w", err)
	}
	return data, nil
}

func (s *xlspathService) Delete(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}
	return nil
}
