package service

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strings"

	"gorm.io/gorm"
	"latih.in-be/internal/model"
	"latih.in-be/internal/repository"
	"latih.in-be/utils/helper"
)

type QuestionService interface {
	Create(ctx context.Context, data model.Question) error
	GetById(ctx context.Context, id int) (*model.Question, error)
	Update(ctx context.Context, newData model.Question, id int, userId int) (*model.Question, error)
	Delete(ctx context.Context, id int, userId int) error
	GetMany(ctx context.Context, limit int, offset int) ([]model.Question, error)
	CreateWithOptions(ctx context.Context, data model.Question) error
	CreateFromJson(ctx context.Context, file *multipart.FileHeader) error
	GetByExam(ctx context.Context, examId int, limit int, offset int) ([]model.Question, error)
	GetByCreatorId(ctx context.Context, creatorId int, limit int, offset int) ([]model.Question, error)
	GetByDifficult(ctx context.Context, diff string, limit int, offset int) ([]model.Question, error)
	GetBySubject(ctx context.Context, subjectId int, limit int, offset int) ([]model.Question, error)
}

type questionService struct {
	repo     repository.QuestionRepository
	userRepo repository.UserRepository
	optRepo  repository.OptionRepository
}

func NewQuestionService(repo repository.QuestionRepository, userRepo repository.UserRepository, optRepo repository.OptionRepository) QuestionService {
	return &questionService{
		repo:     repo,
		userRepo: userRepo,
		optRepo:  optRepo,
	}
}

func (s *questionService) Create(ctx context.Context, data model.Question) error {
	if helper.IsValidSubjectTitle(data.Subject.Title) {
		return fmt.Errorf("invalid subject: %s", data.Subject.Title)
	}

	if err := s.repo.Create(ctx, data); err != nil {
		return fmt.Errorf("failed to create question: %w", err)
	}

	return nil
}

func (s *questionService) GetById(ctx context.Context, id int) (*model.Question, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("question with id %d not found", id)
		}
		return nil, fmt.Errorf("data with id %d not found: %w", id, err)
	}
	return data, nil
}

func (s *questionService) Update(ctx context.Context, newData model.Question, id int, userId int) (*model.Question, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("question with id %d not found", id)
		}
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

func (s *questionService) GetMany(ctx context.Context, limit int, offset int) ([]model.Question, error) {
	data, err := s.repo.GetMany(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get all data: %w", err)
	}
	return data, nil
}

func (s *questionService) Delete(ctx context.Context, id int, userId int) error {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return fmt.Errorf("data is unavaible %w", err)
	}

	user, err := s.userRepo.GetById(ctx, userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user with id %d not found", id)
		}
		return fmt.Errorf("user is unavaible %w", err)
	}

	if user.Id != data.CreatorId && user.Role != model.RoleAdmin {
		return fmt.Errorf("you are not the creator or admin")
	}

	if err := s.optRepo.DeleteByQuestionId(ctx, id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("question with id %d not found", id)
		}
		return fmt.Errorf("failed to delete data: %w", err)
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}
	return nil
}

func (s *questionService) CreateWithOptions(ctx context.Context, data model.Question) error {
	if data.Difficulty != "easy" && data.Difficulty != "medium" && data.Difficulty != "hard" {
		return fmt.Errorf("invalid difficulty")
	}

	valid := false
	for _, opt := range data.Options {
		if opt.IsCorrect {
			valid = true
			break
		}
	}
	if !valid {
		return fmt.Errorf("at least one option must be correct")
	}

	if err := s.repo.CreateWithOptions(ctx, data); err != nil {
		return fmt.Errorf("failed to create question: %w", err)
	}

	return nil
}

func (s *questionService) CreateFromJson(ctx context.Context, file *multipart.FileHeader) error {
	fileContent, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed opening file: %w", err)
	}
	defer fileContent.Close()

	var questions []model.Question
	decoder := json.NewDecoder(fileContent)
	if err := decoder.Decode(&questions); err != nil {
		return fmt.Errorf("invalid format file: %w", err)
	}

	if len(questions) == 0 {
		return fmt.Errorf("file json empty")
	}

	for i, q := range questions {
		if q.QuestionText == "" {
			return fmt.Errorf("question cannot be empty at index %d", i)
		}
		if q.CreatorId == 0 {
			return fmt.Errorf("creatorId cannot be empty at index %d", i)
		}
		if q.Difficulty == "" {
			return fmt.Errorf("difficulty cannot be empty at index %d", i)
		}
	}

	if err := s.repo.CreateBatch(ctx, questions); err != nil {
		return fmt.Errorf("failed to save to database: %w", err)
	}
	return nil
}

func (s *questionService) GetByExam(ctx context.Context, examId int, limit int, offset int) ([]model.Question, error) {
	data, err := s.repo.GetByExam(ctx, examId, limit, offset)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("question with id %d not found", examId)
		}
		return nil, fmt.Errorf("data with exam id %d not found: %w", examId, err)
	}
	return data, nil
}

func (s *questionService) GetByCreatorId(ctx context.Context, creatorId int, limit int, offset int) ([]model.Question, error) {
	data, err := s.repo.GetByExam(ctx, creatorId, limit, offset)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("creator with id %d not found", creatorId)
		}
		return nil, fmt.Errorf("data with creator id %d not found: %w", creatorId, err)
	}
	return data, nil
}

func (s *questionService) GetByDifficult(ctx context.Context, diff string, limit int, offset int) ([]model.Question, error) {
	data, err := s.repo.GetByDifficult(ctx, diff, limit, offset)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("question with difficulty %s not found", diff)
		}
		return nil, fmt.Errorf("data with difficulty %s not found: %w", diff, err)
	}
	return data, nil
}

func (s *questionService) GetBySubject(ctx context.Context, subjectId int, limit int, offset int) ([]model.Question, error) {
	data, err := s.repo.GetBySubject(ctx, subjectId, limit, offset)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("question with subject id %d not found", subjectId)
		}
		return nil, fmt.Errorf("data with subjectId %d not found: %w", subjectId, err)
	}
	return data, nil
}
