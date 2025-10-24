package service

import (
	"context"
	"fmt"
	"strings"

	"latih.in-be/utils/helper"
	"latih.in-be/internal/model"
	"latih.in-be/internal/repository"
)

type ExamService interface {
	Create(ctx context.Context, data model.Exam) error
	GetById(ctx context.Context, id int) (*model.Exam, error)
	Update(ctx context.Context, data model.Exam, id int) (*model.Exam, error)
	Delete(ctx context.Context, id int, userId int) error
	GetAll(ctx context.Context) ([]model.Exam, error)
}

type examService struct {
	repo     repository.ExamRepository
	userRepo repository.UserRepository
}

func NewExamService(repo repository.ExamRepository) ExamService {
	return &examService{
		repo: repo,
	}
}

func (s *examService) Create(ctx context.Context, data model.Exam) error {
	if data.FinishedAt.Before(*data.StartedAt) {
		return fmt.Errorf("finished_at must be after started_at")
	}

	if err := s.repo.Create(ctx, data); err != nil {
		return fmt.Errorf("failed to create exam: %w", err)
	}

	return nil
}

func (s *examService) GetById(ctx context.Context, id int) (*model.Exam, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("data with id %d not found: %w", id, err)
	}
	return data, nil
}

func (s *examService) Update(ctx context.Context, data model.Exam, id int) (*model.Exam, error) {
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

func (s *examService) GetAll(ctx context.Context) ([]model.Exam, error) {
	data, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all data: %w", err)
	}
	return data, nil
}

func (s *examService) Delete(ctx context.Context, id int, userId int) error {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return fmt.Errorf("data is unavaible %w", err)
	}

	user, err := s.userRepo.GetById(ctx, userId)
	if err != nil {
		return fmt.Errorf("user is unavaible %w", err)
	}

	if user.Id != data.CreatorId && user.Role != model.RoleAdmin {
		return fmt.Errorf("you are not the creator or admin")
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}
	return nil
}
