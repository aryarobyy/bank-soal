package service

import (
	"context"
	"fmt"

	"latih.in-be/internal/model"
	"latih.in-be/internal/repository"
)

type UserAnswerService interface {
	Create(ctx context.Context, userAnswer *model.UserAnswer) error
	GetById(ctx context.Context, id int) (*model.UserAnswer, error)
	GetMany(ctx context.Context, limit int, offset int) ([]model.UserAnswer, int64, error)
	Update(ctx context.Context, id int, userAnswer *model.UserAnswer) (*model.UserAnswer, error)
	Delete(ctx context.Context, id int) error
	GetByExamSessionId(ctx context.Context, examSessionId int, limit int, offset int) ([]model.UserAnswer, int64, error)
	GetByQuestionId(ctx context.Context, questionId int, limit int, offset int) ([]model.UserAnswer, int64, error)
	GetByUserId(ctx context.Context, userId int, limit int, offset int) ([]model.UserAnswer, int64, error)
}

type userAnswerService struct {
	repo       repository.UserAnswerRepository
	optionRepo repository.OptionRepository
}

func NewUserAnswerService(repo repository.UserAnswerRepository, optionRepo repository.OptionRepository) UserAnswerService {
	return &userAnswerService{
		repo:       repo,
		optionRepo: optionRepo,
	}
}

func (s *userAnswerService) Create(ctx context.Context, userAnswer *model.UserAnswer) error {
	checkAnswer, err := s.optionRepo.CheckCorrectAnswer(ctx, userAnswer.QuestionId, userAnswer.Answer)
	if err != nil {
		return fmt.Errorf("failed to check correct answer: %w", err)
	}

	userAnswer.IsCorrect = checkAnswer

	if err := s.repo.Create(ctx, userAnswer); err != nil {
		return fmt.Errorf("failed to create user answer: %w", err)
	}

	return nil
}

func (s *userAnswerService) GetById(ctx context.Context, id int) (*model.UserAnswer, error) {
	if id <= 0 {
		return nil, fmt.Errorf("user answer Id is required and must be greater than 0")
	}

	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user answer with id %d not found: %w", id, err)
	}
	return data, nil
}

func (s *userAnswerService) GetMany(ctx context.Context, limit int, offset int) ([]model.UserAnswer, int64, error) {
	data, total, err := s.repo.GetMany(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all user answers: %w", err)
	}
	return data, total, nil
}

func (s *userAnswerService) Update(ctx context.Context, id int, userAnswer *model.UserAnswer) (*model.UserAnswer, error) {
	if id <= 0 {
		return nil, fmt.Errorf("user answer Id is required and must be greater than 0")
	}

	_, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user answer with id %d not found: %w", id, err)
	}

	updated, err := s.repo.Update(ctx, id, userAnswer)
	if err != nil {
		return nil, fmt.Errorf("failed to update user answer: %w", err)
	}

	return updated, nil
}

func (s *userAnswerService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("user answer Id is required and must be greater than 0")
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete user answer: %w", err)
	}
	return nil
}

func (s *userAnswerService) GetByExamSessionId(ctx context.Context, examSessionId int, limit int, offset int) ([]model.UserAnswer, int64, error) {
	if examSessionId <= 0 {
		return nil, 0, fmt.Errorf("exam session Id is required and must be greater than 0")
	}

	data, total, err := s.repo.GetByExamSessionId(ctx, examSessionId, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get user answers by exam session Id %d: %w", examSessionId, err)
	}
	return data, total, nil
}

func (s *userAnswerService) GetByQuestionId(ctx context.Context, questionId int, limit int, offset int) ([]model.UserAnswer, int64, error) {
	if questionId <= 0 {
		return nil, 0, fmt.Errorf("question Id is required and must be greater than 0")
	}

	data, total, err := s.repo.GetByQuestionId(ctx, questionId, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get user answers by question Id %d: %w", questionId, err)
	}
	return data, total, nil
}

func (s *userAnswerService) GetByUserId(ctx context.Context, userId int, limit int, offset int) ([]model.UserAnswer, int64, error) {
	if userId <= 0 {
		return nil, 0, fmt.Errorf("user Id is required and must be greater than 0")
	}

	data, total, err := s.repo.GetByUserId(ctx, userId, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get user answers by user Id %d: %w", userId, err)
	}
	return data, total, nil
}
