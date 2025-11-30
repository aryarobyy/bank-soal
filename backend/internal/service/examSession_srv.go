package service

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
	"latih.in-be/internal/model"
	"latih.in-be/internal/repository"
)

type ExamSessionService interface {
	Create(ctx context.Context, e model.ExamSession, userId int, examId int) (*model.ExamSession, error)
	GetById(ctx context.Context, id int) (*model.ExamSession, error)
	Update(ctx context.Context, id int, e model.UpdateExamSession) (*model.ExamSession, error)
	Delete(ctx context.Context, id int) error
	GetMany(ctx context.Context, examId int, limit int, offset int) ([]model.ExamSession, int64, error)
	UpdateCurrNo(ctx context.Context, id int, no model.UpdateCurrNo) (*model.ExamSession, error)
	FinishExam(ctx context.Context, userId int, id int) (*model.ExamSession, error)
	GetScore(ctx context.Context, sessionId int, userId int) (*model.ExamSession, error)
	GetUserSession(ctx context.Context, userId int, limit int, offset int) ([]model.ExamSession, int64, error)
}

type examSessionService struct {
	repo         repository.ExamSessionRepository
	examRepo     repository.ExamRepository
	answerRepo   repository.UserAnswerRepository
	questionRepo repository.QuestionRepository
}

func NewExamSessionService(
	repo repository.ExamSessionRepository,
	examRepo repository.ExamRepository,
	answerRepo repository.UserAnswerRepository,
	questionRepo repository.QuestionRepository,
) ExamSessionService {
	return &examSessionService{
		repo:         repo,
		examRepo:     examRepo,
		answerRepo:   answerRepo,
		questionRepo: questionRepo,
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

	println(exam)
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

	if data.FinishedAt != nil && data.FinishedAt.After(time.Now()) {
		return nil, fmt.Errorf("session %d already finished", id)
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

func (s *examSessionService) GetMany(ctx context.Context, examId int, limit int, offset int) ([]model.ExamSession, int64, error) {
	data, total, err := s.repo.GetMany(ctx, examId, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get sessions: %w", err)
	}
	if data != nil || len(data) == 0 {
		return nil, 0, fmt.Errorf("no exam sessions found")
	}

	return data, total, nil
}

func (s *examSessionService) UpdateCurrNo(ctx context.Context, id int, no model.UpdateCurrNo) (*model.ExamSession, error) {
	updated, err := s.repo.UpdateCurrNo(ctx, id, no)
	if err != nil {
		return nil, fmt.Errorf("failed to update current question: %w", err)
	}
	return updated, nil
}

func (s *examSessionService) FinishExam(ctx context.Context, userId int, id int) (*model.ExamSession, error) {
	session, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find session: %w", err)
	}

	userScore, maxScore, err := s.calculateScore(ctx, userId, id, session.ExamId)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate score: %w", err)
	}

	var percentageScore float64
	if maxScore > 0 {
		percentageScore = (float64(userScore) / float64(maxScore)) * 100
	}

	now := time.Now()
	session.FinishedAt = &now
	session.Percentage = percentageScore
	session.Status = model.SessionFinished
	session.Score = userScore
	session.MaxScore = maxScore
	if percentageScore >= 75.0 {
		session.IsPassed = true
	}

	data := model.FinishExam{
		FinishedAt: *session.FinishedAt,
		Status:     session.Status,
		Score:      session.Score,
		Percentage: session.Percentage,
		MaxScore:   session.MaxScore,
	}

	updated, err := s.repo.FinishExam(ctx, id, data)
	if err != nil {
		return nil, fmt.Errorf("failed to finish exam: %w", err)
	}

	return updated, nil
}

func (s *examSessionService) calculateScore(ctx context.Context, userId, sessionId int, examId int) (int, int, error) {
	userAnswers, err := s.answerRepo.GetAllUserAnswers(ctx, userId, sessionId)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get user answers: %w", err)
	}

	examQuestions, err := s.questionRepo.GetByExamId(ctx, examId)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get exam questions: %w", err)
	}

	maxScore := 0
	for _, eq := range examQuestions {
		question, err := s.questionRepo.GetById(ctx, eq.Id)
		if err != nil {
			return 0, 0, fmt.Errorf("failed to get question %d: %w", eq.Id, err)
		}
		maxScore += question.Score
	}

	userScore := 0
	for _, answer := range userAnswers {
		fmt.Println("skaksdosak saksdk", answer.Id, answer.QuestionId, answer.IsCorrect)
		if answer.IsCorrect {
			question, err := s.questionRepo.GetById(ctx, answer.QuestionId)
			if err != nil {
				return 0, 0, fmt.Errorf("failed to get question %d: %w", answer.QuestionId, err)
			}
			userScore += question.Score
		}
	}
	fmt.Print("tsaysadus saskas", userScore, maxScore)
	return userScore, maxScore, nil
}

func (s *examSessionService) GetScore(ctx context.Context, sessionId int, userId int) (*model.ExamSession, error) {
	data, err := s.repo.GetScore(ctx, sessionId, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get score: %w", err)
	}
	return data, nil
}

func (s *examSessionService) GetUserSession(ctx context.Context, userId int, limit int, offset int) ([]model.ExamSession, int64, error) {
	data, total, err := s.repo.GetUserSession(ctx, userId, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get user session: %w", err)
	}
	return data, total, nil
}
