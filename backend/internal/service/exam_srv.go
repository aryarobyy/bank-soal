package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
	"latih.in-be/internal/model"
	"latih.in-be/internal/repository"
	"latih.in-be/utils/helper"
)

type ExamService interface {
	Create(ctx context.Context, data model.CreateExam) error
	GetById(ctx context.Context, id int) (*model.Exam, error)
	Update(ctx context.Context, newData model.Exam, id int, userId int) (*model.Exam, error)
	Delete(ctx context.Context, id int, userId int) error
	GetMany(ctx context.Context, limit int, offset int) ([]model.Exam, int64, error)
	AddQuestions(ctx context.Context, examId int, questionIds []int) error
	ReplaceQuestions(ctx context.Context, examId int, questionIds []int) error
	RemoveQuestions(ctx context.Context, examId int, questionIds []int) error
	GetByCreator(ctx context.Context, creatorId int, limit int, offset int) ([]model.Exam, error)
}

type examService struct {
	repo         repository.ExamRepository
	userRepo     repository.UserRepository
	questionRepo repository.QuestionRepository
}

func NewExamService(
	repo repository.ExamRepository,
	userRepo repository.UserRepository,
	questionRepo repository.QuestionRepository,
) ExamService {
	return &examService{
		repo:         repo,
		userRepo:     userRepo,
		questionRepo: questionRepo,
	}
}

func (s *examService) Create(ctx context.Context, data model.CreateExam) error {
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

func (s *examService) Update(ctx context.Context, newData model.Exam, id int, userId int) (*model.Exam, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("data is unavaible %w", err)
	}

	user, err := s.userRepo.GetById(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("user is unavaible %w", err)
	}

	if user.Id != data.CreatorId && user.Role != model.RoleAdmin {
		return nil, fmt.Errorf("you are not the creator or admin")
	}

	updatedData, err := s.repo.Update(ctx, newData, id)
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

func (s *examService) GetMany(ctx context.Context, limit int, offset int) ([]model.Exam, int64, error) {
	data, total, err := s.repo.GetMany(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all data: %w", err)
	}
	return data, total, nil
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

func (s *examService) AddQuestions(ctx context.Context, examId int, questionIds []int) error {
	exam, err := s.repo.GetById(ctx, examId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("exam with id %d not found", examId)
		}
		return fmt.Errorf("failed to get exam: %w", err)
	}

	addedScore := 0

	for _, qid := range questionIds {
		q, err := s.questionRepo.GetById(ctx, qid)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("question with id %d not found", qid)
			}
			return fmt.Errorf("failed to get question %d: %w", qid, err)
		}

		addedScore += q.Score
	}

	finalScore := exam.Score + addedScore

	if err := s.repo.AddQuestions(ctx, examId, questionIds); err != nil {
		return fmt.Errorf("failed to add questions: %w", err)
	}

	if err := s.repo.UpdateScore(ctx, examId, finalScore); err != nil {
		return fmt.Errorf("failed to update exam score: %w", err)
	}

	return nil
}

func (s *examService) ReplaceQuestions(ctx context.Context, examId int, questionIds []int) error {
	_, err := s.repo.GetById(ctx, examId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("exam with id %d not found", examId)
		}
		return fmt.Errorf("failed to get exam: %w", err)
	}

	newTotal := 0

	for _, qid := range questionIds {
		q, err := s.questionRepo.GetById(ctx, qid)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("question with id %d not found", qid)
			}
			return fmt.Errorf("failed to get question %d: %w", qid, err)
		}
		newTotal += q.Score
	}

	if err := s.repo.ReplaceQuestions(ctx, examId, questionIds); err != nil {
		return fmt.Errorf("failed to replace exam questions: %w", err)
	}

	if err := s.repo.UpdateScore(ctx, examId, newTotal); err != nil {
		return fmt.Errorf("failed to update exam total score: %w", err)
	}

	return nil
}

func (s *examService) RemoveQuestions(ctx context.Context, examId int, questionIds []int) error {
	exam, err := s.repo.GetById(ctx, examId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("exam with id %d not found", examId)
		}
		return fmt.Errorf("failed to get exam: %w", err)
	}

	reduce := 0

	for _, qid := range questionIds {
		q, err := s.questionRepo.GetById(ctx, qid)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("question with id %d not found", qid)
			}
			return fmt.Errorf("failed to get question %d: %w", qid, err)
		}
		reduce += q.Score
	}

	if err := s.repo.RemoveQuestions(ctx, examId, questionIds); err != nil {
		return fmt.Errorf("failed to remove exam questions: %w", err)
	}

	final := exam.Score - reduce
	if final < 0 {
		final = 0
	}

	if err := s.repo.UpdateScore(ctx, examId, final); err != nil {
		return fmt.Errorf("failed to update exam total score: %w", err)
	}

	return nil
}

func (s *examService) GetByCreator(ctx context.Context, creatorId int, limit int, offset int) ([]model.Exam, error) {
	data, err := s.repo.GetByCreator(ctx, creatorId, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get exams %w", err)
	}
	return data, nil
}
