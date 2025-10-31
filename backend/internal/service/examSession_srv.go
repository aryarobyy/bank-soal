package service

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"latih.in-be/internal/model"
	"latih.in-be/internal/repository"
)

type ExamSessionService interface {
	Create(ctx context.Context, e model.ExamSession, userId int, examId int) (*model.ExamSession, error)
	GetById(ctx context.Context, id int) (*model.ExamSession, error)
	Update(ctx context.Context, id int, e model.UpdateExamSession) (*model.ExamSession, error)
	Delete(ctx context.Context, id int) error
	GetMany(ctx context.Context, userId int, limit int, offset int) ([]model.ExamSession, error)
	UpdateCurrNo(ctx context.Context, id int, no model.UpdateCurrNo) (*model.ExamSession, error)
	FinishExam(ctx context.Context, id int, e model.FinishExam) (*model.ExamSession, error)
}

type examSessionService struct {
	repo     repository.ExamSessionRepository
	examRepo repository.ExamRepository
}

func NewExamSessionService(repo repository.ExamSessionRepository, examRepo repository.ExamRepository) ExamSessionService {
	return &examSessionService{
		repo:     repo,
		examRepo: examRepo,
	}
}

func (s *examSessionService) Create(ctx context.Context, e model.ExamSession, userId int, examId int) (*model.ExamSession, error) {
	if err := s.repo.CheckUserSession(ctx, userId, examId); err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("failed to check existing session: %w", err)
	}

	e.UserId = userId
	e.ExamId = examId

	session, err := s.repo.Create(ctx, e)
	if err != nil {
		return nil, fmt.Errorf("failed to create exam session: %w", err)
	}

	exam, err := s.examRepo.StartSession(ctx, examId)
	if err != nil {
		return session, fmt.Errorf("failed to load exam questions: %w", err)
	}
	fmt.Println(exam)
	return session, nil
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
