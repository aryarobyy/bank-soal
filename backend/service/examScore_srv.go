package service

import (
	"context"
	"fmt"
	"strings"

	"latih.in-be/helper"
	"latih.in-be/model"
	"latih.in-be/repository"
)

type ExamScoreService interface {
	Create(ctx context.Context, data model.ExamScore) error
	GetById(ctx context.Context, id int) (*model.ExamScore, error)
	Update(ctx context.Context, data model.ExamScore, id int) (*model.ExamScore, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]model.ExamScore, error)
}

type examScoreService struct {
	repo repository.ExamScoreRepository
}

func NewExamScoreService(repo repository.ExamScoreRepository) ExamScoreService {
	return &examScoreService{
		repo: repo,
	}
}

func (s *examScoreService) Create(ctx context.Context, data model.ExamScore) error {
	err := s.repo.Create(ctx, data)
	if err != nil {
		if strings.Contains(err.Error(), "Unknown column") {
			var fieldName string
			parts := strings.Split(err.Error(), "'")
			if len(parts) >= 2 {
				fieldName = parts[1]
			}

			val := helper.GetFieldValue(data, fieldName)

			return fmt.Errorf("field '%s' with value '%v' is undefined", fieldName, val)
		}

		return fmt.Errorf("create gagal: %v", err)
	}
	return nil
}

func (s *examScoreService) GetById(ctx context.Context, id int) (*model.ExamScore, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("data with id %d not found: %w", id, err)
	}
	return data, nil
}

func (s *examScoreService) Update(ctx context.Context, data model.ExamScore, id int) (*model.ExamScore, error) {
	updatedData, err := s.repo.Update(ctx, data, id)
	if err != nil {
		if strings.Contains(err.Error(), "Unknown column") {
			var fieldName string
			parts := strings.Split(err.Error(), "'")
			if len(parts) >= 2 {
				fieldName = parts[1]
			}

			val := helper.GetFieldValue(data, fieldName)

			return nil, fmt.Errorf("field '%s' with value '%v' is undefined", fieldName, val)
		}

		return nil, fmt.Errorf("update gagal: %v", err)
	}
	return updatedData, nil
}

func (s *examScoreService) GetAll(ctx context.Context) ([]model.ExamScore, error) {
	data, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all data: %w", err)
	}
	return data, nil
}

func (s *examScoreService) Delete(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}
	return nil
}
