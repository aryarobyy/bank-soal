package service

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"latih.in-be/internal/model"
	"latih.in-be/internal/repository"
)

type ExamSessionService interface {
	Create(ctx context.Context, e model.ExamSession, userId int, examId int) error
	GetById(ctx context.Context, id int) (*model.ExamSession, error)
	Update(ctx context.Context, id int, e model.UpdateExamSession) (*model.ExamSession, error)
	Delete(ctx context.Context, id int) error
	GetMany(ctx context.Context, userId int, limit int, offset int) ([]model.ExamSession, error)
	UpdateCurrNo(ctx context.Context, id int, no model.UpdateCurrNo) (*model.ExamSession, error)
	FinishExam(ctx context.Context, id int, e model.FinishExam) (*model.ExamSession, error)
}

type examSessionService struct {
	repo repository.ExamSessionRepository
}

func NewExamSessionService(repo repository.ExamSessionRepository) ExamSessionService {
	return &examSessionService{
		repo: repo,
	}
}

func (s *examSessionService) Create(ctx context.Context, e model.ExamSession, userId int, examId int) error {
	err := s.repo.CheckUserSession(ctx, userId, examId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("failed to check existing session: %w", err)
	}

	if err := s.repo.Create(ctx, e); err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}
	return nil
}

func (s *examSessionService) GetById(ctx context.Context, id int) (*model.ExamSession, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("session with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to get session: %w", err)
	}
	return data, nil
}

func (s *examSessionService) Update(ctx context.Context, id int, e model.UpdateExamSession) (*model.ExamSession, error) {
	updated, err := s.repo.Update(ctx, id, e)
	if err != nil {
		return nil, fmt.Errorf("failed to update session: %w", err)
	}
	return updated, nil
}

func (s *examSessionService) Delete(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}
	return nil
}

func (s *examSessionService) GetMany(ctx context.Context, userId int, limit int, offset int) ([]model.ExamSession, error) {
	data, err := s.repo.GetMany(ctx, userId, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get sessions: %w", err)
	}
	return data, nil
}

func (s *examSessionService) UpdateCurrNo(ctx context.Context, id int, no model.UpdateCurrNo) (*model.ExamSession, error) {
	updated, err := s.repo.UpdateCurrNo(ctx, id, no)
	if err != nil {
		return nil, fmt.Errorf("failed to update current question: %w", err)
	}
	return updated, nil
}

func (s *examSessionService) FinishExam(ctx context.Context, id int, e model.FinishExam) (*model.ExamSession, error) {
	session, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find session: %w", err)
	}

	session.FinishedAt = &e.FinishedAt
	session.Score = e.Score

	data := model.FinishExam{
		FinishedAt: *session.FinishedAt,
		Score:      session.Score,
	}

	updated, err := s.repo.FinishExam(ctx, id, data)
	if err != nil {
		return nil, fmt.Errorf("failed to finish exam: %w", err)
	}

	return updated, nil
}
