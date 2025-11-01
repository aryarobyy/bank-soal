package service

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"latih.in-be/internal/repository"
)

type ExamQuestionService interface {
	AddQuestionToExam(ctx context.Context, examId int, questionIds []int) error
	UpdateQuestionsInExam(ctx context.Context, examId int, questionIds []int) error
	RemoveQuestionsFromExam(ctx context.Context, examId int, questionIds []int) error
}

type examQuestionService struct {
	repo      repository.ExamQuestionRepository
	questRepo repository.QuestionRepository
	examRepo  repository.ExamRepository
}

func NewExamQuestionService(
	repo repository.ExamQuestionRepository,
	questRepo repository.QuestionRepository,
	examRepo repository.ExamRepository,
) ExamQuestionService {
	return &examQuestionService{
		repo:      repo,
		questRepo: questRepo,
		examRepo:  examRepo,
	}
}

func (s *examQuestionService) AddQuestionToExam(ctx context.Context, examId int, questionIds []int) error {
	exam, err := s.examRepo.GetById(ctx, examId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("exam with id %d not found", examId)
		}
		return fmt.Errorf("failed to get exam: %w", err)
	}

	for _, qid := range questionIds {
		q, err := s.questRepo.GetById(ctx, qid)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("question with id %d not found", qid)
			}
			return fmt.Errorf("failed to get question %d: %w", qid, err)
		}
		if q == nil {
			return fmt.Errorf("invalid question id: %d", qid)
		}
	}

	if err := s.repo.AddQuestionToExam(ctx, exam.Id, questionIds); err != nil {
		return fmt.Errorf("failed to add questions to exam: %w", err)
	}

	return nil
}

func (s *examQuestionService) UpdateQuestionsInExam(ctx context.Context, examId int, questionIds []int) error {
	if _, err := s.examRepo.GetById(ctx, examId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("exam with id %d not found", examId)
		}
		return err
	}

	for _, qid := range questionIds {
		if _, err := s.questRepo.GetById(ctx, qid); err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("question with id %d not found", qid)
			}
			return err
		}
	}

	if err := s.repo.UpdateQuestionsInExam(ctx, examId, questionIds); err != nil {
		return fmt.Errorf("failed to update questions in exam: %w", err)
	}

	return nil
}

func (s *examQuestionService) RemoveQuestionsFromExam(ctx context.Context, examId int, questionIds []int) error {
	if _, err := s.examRepo.GetById(ctx, examId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("exam with id %d not found", examId)
		}
		return err
	}

	if err := s.repo.RemoveQuestionsFromExam(ctx, examId, questionIds); err != nil {
		return fmt.Errorf("failed to remove questions from exam: %w", err)
	}

	return nil
}
