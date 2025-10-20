package service

import (
	"context"
	"fmt"

	"latih.in-be/model"
	"latih.in-be/repository"
)

type OptionService interface {
	Create(ctx context.Context, data model.Option) error
	GetById(ctx context.Context, id int) (*model.Option, error)
	Update(ctx context.Context, data model.Option, id int) (*model.Option, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context, qId int) ([]model.Option, error)
}

type optionService struct {
	repo repository.OptionRepository
}

func NewOptionService(repo repository.OptionRepository) OptionService {
	return &optionService{
		repo: repo,
	}
}

func (s *optionService) Create(ctx context.Context, data model.Option) error {
	err := s.repo.Create(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to create option: %w", err)
	}

	return nil
}

func (s *optionService) GetById(ctx context.Context, id int) (*model.Option, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("data with id %d not found: %w", id, err)
	}
	return data, nil
}

func (s *optionService) Update(ctx context.Context, data model.Option, id int) (*model.Option, error) {
	updatedData, err := s.repo.Update(ctx, data, id)
	if err != nil {
		return nil, fmt.Errorf("update failed: %w", err)
	}
	return updatedData, nil
}

func (s *optionService) GetAll(ctx context.Context, qId int) ([]model.Option, error) {
	data, err := s.repo.GetAll(ctx, qId)
	if err != nil {
		return nil, fmt.Errorf("failed to get all data: %w", err)
	}
	return data, nil
}

func (s *optionService) Delete(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}
	return nil
}
