package service

import (
	"context"
	"fmt"

	"latih.in-be/internal/model"
	"latih.in-be/internal/repository"
)

type ScoreService interface {
	Create(ctx context.Context, data model.Score) error
	GetById(ctx context.Context, id int) (*model.Score, error)
	Update(ctx context.Context, data model.Score, id int) (*model.Score, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context, qId int) ([]model.Score, error)
}

type scoreService struct {
	repo repository.ScoreRepository
}

func NewScoreService(repo repository.ScoreRepository) ScoreService {
	return &scoreService{
		repo: repo,
	}
}

func (s *scoreService) Create(ctx context.Context, data model.Score) error {
	err := s.repo.Create(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to create score: %w", err)
	}

	return nil
}

func (s *scoreService) GetById(ctx context.Context, id int) (*model.Score, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("data with id %d not found: %w", id, err)
	}
	return data, nil
}

func (s *scoreService) Update(ctx context.Context, data model.Score, id int) (*model.Score, error) {
	updatedData, err := s.repo.Update(ctx, data, id)
	if err != nil {
		return nil, fmt.Errorf("update failed: %w", err)
	}
	return updatedData, nil
}

func (s *scoreService) GetAll(ctx context.Context, qId int) ([]model.Score, error) {
	data, err := s.repo.GetAll(ctx, qId)
	if err != nil {
		return nil, fmt.Errorf("failed to get all data: %w", err)
	}
	return data, nil
}

func (s *scoreService) Delete(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}
	return nil
}
