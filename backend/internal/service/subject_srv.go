package service

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"latih.in-be/internal/model"
	"latih.in-be/internal/repository"
)

type SubjectService interface {
	Create(ctx context.Context, data model.Subject) error
	GetById(ctx context.Context, id int) (*model.Subject, error)
	Update(ctx context.Context, data model.Subject, id int) (*model.Subject, error)
	Delete(ctx context.Context, id int) error
	GetMany(ctx context.Context, limit int, offset int) ([]model.Subject, int64, error)
	GetByCode(ctx context.Context, code string) (*model.Subject, error)
}

type subjectService struct {
	repo repository.SubjectRepository
}

func NewSubjectService(repo repository.SubjectRepository) SubjectService {
	return &subjectService{
		repo: repo,
	}
}

func (s *subjectService) Create(ctx context.Context, data model.Subject) error {
	err := s.repo.Create(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to create subject: %w", err)
	}

	return nil
}

func (s *subjectService) GetById(ctx context.Context, id int) (*model.Subject, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("data with id %d not found: %w", id, err)
	}
	return data, nil
}

func (s *subjectService) GetByCode(ctx context.Context, code string) (*model.Subject, error) {
	data, err := s.repo.GetByCode(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("data with id %s not found: %w", code, err)
	}
	return data, nil
}

func (s *subjectService) Update(ctx context.Context, data model.Subject, id int) (*model.Subject, error) {
	existingData, err := s.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("subject with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to check existing subject: %w", err)
	}

	if existingData == nil {
		return nil, fmt.Errorf("subject with id %d not found", id)
	}

	updatedData, err := s.repo.Update(ctx, data, id)
	if err != nil {
		return nil, fmt.Errorf("failed to update subject: %w", err)
	}

	return updatedData, nil
}

func (s *subjectService) GetMany(ctx context.Context, limit int, offset int) ([]model.Subject, int64, error) {
	data, total, err := s.repo.GetMany(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all data: %w", err)
	}
	return data, total, nil
}

func (s *subjectService) Delete(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}
	return nil
}
