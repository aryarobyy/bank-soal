package service

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"

	"latih.in-be/model"
	"latih.in-be/repository"
)

type QuestionService interface {
	Create(ctx context.Context, data model.Question) error
	GetById(ctx context.Context, id int) (*model.Question, error)
	Update(ctx context.Context, data model.Question, id int) (*model.Question, error)
	Delete(ctx context.Context, id int, userId int) error
	GetAll(ctx context.Context) ([]model.Question, error)
	CreateWithOptions(ctx context.Context, data model.Question) error
	CreateFromJson(ctx context.Context, file *multipart.FileHeader) error
	GetByExam(ctx context.Context, examId int) ([]model.Question, error)
	GetByCreatorId(ctx context.Context, creatorId int) ([]model.Question, error)
	GetByDifficult(ctx context.Context, diff string) ([]model.Question, error)
}

type questionService struct {
	repo     repository.QuestionRepository
	userRepo repository.UserRepository
}

func NewQuestionService(repo repository.QuestionRepository) QuestionService {
	return &questionService{
		repo: repo,
	}
}

func (s *questionService) Create(ctx context.Context, data model.Question) error {
	if IsValidSubjectTitle(data.Subject.Title) {
		return fmt.Errorf("invalid subject: %s", data.Subject)
	}

	if err := s.repo.Create(ctx, data); err != nil {
		return fmt.Errorf("failed to create question: %w", err)
	}

	return nil
}

func (s *questionService) GetById(ctx context.Context, id int) (*model.Question, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("data with id %d not found: %w", id, err)
	}
	return data, nil
}

func (s *questionService) Update(ctx context.Context, data model.Question, id int) (*model.Question, error) {
	updatedData, err := s.repo.Update(ctx, data, id)
	if err != nil {
		return nil, fmt.Errorf("update failed: %w", err)
	}
	return updatedData, nil
}

func (s *questionService) GetAll(ctx context.Context) ([]model.Question, error) {
	data, err := s.repo.GetAll(ctx)
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

func (s *questionService) GetByExam(ctx context.Context, examId int) ([]model.Question, error) {
	data, err := s.repo.GetByExam(ctx, examId)
	if err != nil {
		return nil, fmt.Errorf("data with exam id %d not found: %w", examId, err)
	}
	return data, nil
}

func (s *questionService) GetByCreatorId(ctx context.Context, creatorId int) ([]model.Question, error) {
	data, err := s.repo.GetByExam(ctx, creatorId)
	if err != nil {
		return nil, fmt.Errorf("data with creator id %d not found: %w", creatorId, err)
	}
	return data, nil
}

func (s *questionService) GetByDifficult(ctx context.Context, diff string) ([]model.Question, error) {
	data, err := s.repo.GetByDifficult(ctx, diff)
	if err != nil {
		return nil, fmt.Errorf("data with difficulty %s not found: %w", diff, err)
	}
	return data, nil
}

func IsValidSubjectTitle(title model.SubjectTitle) bool {
	switch title {
	case model.SubjectKalkulus,
		model.SubjectMatDis,
		model.SubjectAutomata,
		model.SubjectData,
		model.SubjectMetNum:
		return true
	default:
		return false
	}
}
