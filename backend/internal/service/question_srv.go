package service

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"latih.in-be/internal/model"
	"latih.in-be/internal/repository"
	"latih.in-be/utils/helper"
)

type QuestionService interface {
	Create(ctx context.Context, c *gin.Context, data *model.Question) error
	GetById(ctx context.Context, id int) (*model.Question, error)
	Update(ctx context.Context, c *gin.Context, req *model.Question, id int, userId int) (*model.Question, error)
	Delete(ctx context.Context, id int, userId int) error
	GetMany(ctx context.Context, limit int, offset int) ([]model.Question, int64, error)
	CreateWithOptions(ctx context.Context, data model.Question) error
	CreateFromJson(ctx context.Context, file *multipart.FileHeader) error
	GetByExam(ctx context.Context, examId int, limit int, offset int) ([]model.Question, int64, error)
	GetByCreatorId(ctx context.Context, creatorId int, limit int, offset int) ([]model.Question, int64, error)
	GetByDifficult(ctx context.Context, diff string, limit int, offset int) ([]model.Question, int64, error)
	GetBySubject(ctx context.Context, subjectId int, limit int, offset int) ([]model.Question, int64, error)
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

func (s *questionService) Create(ctx context.Context, c *gin.Context, data *model.Question) error {

	if len(data.Options) == 0 {
		return fmt.Errorf("options can't be null")
	}

	switch data.Difficulty {
	case model.DifficultyEasy:
		if data.Score == 0 {
			data.Score = 5
		}
		if data.Score < 3 || data.Score > 8 {
			return fmt.Errorf("score for easy difficulty must be between 3 and 8")
		}

	case model.DifficultyMedium:
		if data.Score == 0 {
			data.Score = 10
		}
		if data.Score < 10 || data.Score > 15 {
			return fmt.Errorf("score for medium difficulty must be between 10 and 15")
		}

	case model.DifficultyHard:
		if data.Score == 0 {
			data.Score = 20
		}
		if data.Score < 18 || data.Score > 23 {
			return fmt.Errorf("score for hard difficulty must be between 18 and 23")
		}

	default:
		return fmt.Errorf("invalid difficulty level")
	}

	if err := s.repo.Create(ctx, data); err != nil {
		return fmt.Errorf("failed to create question: %w", err)
	}

	imgDir := "./storages/images/question"
	_, err := helper.UploadImage(c, data.Id, imgDir)
	if err != nil {
		return fmt.Errorf("failed to upload image: %w", err)
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

func (s *questionService) Update(ctx context.Context, c *gin.Context, req *model.Question, id int, userId int) (*model.Question, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("question with id %d not found", id)
		}
		return nil, fmt.Errorf("data is unavailable: %w", err)
	}

	user, err := s.userRepo.GetById(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("user is unavailable: %w", err)
	}

	if user.Id != data.CreatorId && user.Role != model.RoleAdmin {
		return nil, fmt.Errorf("you are not the creator or admin")
	}

	if req.Difficulty != "" {
		switch req.Difficulty {
		case model.DifficultyEasy:
			if req.Score != 0 && (req.Score < 3 || req.Score > 8) {
				return nil, fmt.Errorf("score for easy difficulty must be between 3 and 8")
			}
		case model.DifficultyMedium:
			if req.Score != 0 && (req.Score < 10 || req.Score > 15) {
				return nil, fmt.Errorf("score for medium difficulty must be between 10 and 15")
			}
		case model.DifficultyHard:
			if req.Score != 0 && (req.Score < 18 || req.Score > 23) {
				return nil, fmt.Errorf("score for hard difficulty must be between 18 and 23")
			}
		default:
			return nil, fmt.Errorf("invalid difficulty level")
		}
	}

	fileHeader, err := c.FormFile("image")
	hasNewImage := err == nil

	if hasNewImage {
		ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			return nil, fmt.Errorf("invalid image format. Only JPG and PNG are allowed")
		}

		if data.ImgUrl != "" {
			if err := helper.DeleteImage(data.ImgUrl); err != nil {
				fmt.Printf("Warning: failed to delete old image: %v\n", err)
			}
		}

		imgDir := "./storages/images/question"
		imageURL, err := helper.UploadImage(c, id, imgDir)
		if err != nil {
			return nil, fmt.Errorf("failed to upload image: %w", err)
		}

		req.ImgUrl = imageURL
	} else {
		req.ImgUrl = data.ImgUrl
	}

	updatedData, err := s.repo.Update(ctx, *req, id)
	if err != nil {
		if strings.Contains(err.Error(), "Unknown column") {
			parts := strings.Split(err.Error(), "'")
			if len(parts) >= 2 {
				fieldName := parts[1]
				val := helper.GetFieldValue(*req, fieldName)
				return nil, fmt.Errorf("field '%s' with value '%v' is undefined", fieldName, val)
			}
		}
		return updatedData, fmt.Errorf("update gagal: %v", err)
	}

	if len(req.Options) > 0 {
		existingOptions := data.Options

		minLen := len(req.Options)
		if len(existingOptions) < minLen {
			minLen = len(existingOptions)
		}

		for i := 0; i < minLen; i++ {
			optToUpdate := req.Options[i]
			optToUpdate.Id = existingOptions[i].Id
			optToUpdate.QuestionId = id

			_, err := s.optRepo.Update(ctx, optToUpdate, existingOptions[i].Id)
			if err != nil {
				return nil, fmt.Errorf("failed to update option at index %d: %w", i, err)
			}
		}

		if len(req.Options) > len(existingOptions) {
			for i := len(existingOptions); i < len(req.Options); i++ {
				newOpt := req.Options[i]
				newOpt.QuestionId = id
				err := s.optRepo.Create(ctx, newOpt)
				if err != nil {
					return nil, fmt.Errorf("failed to create option at index %d: %w", i, err)
				}
			}
		}

		if len(existingOptions) > len(req.Options) {
			for i := len(req.Options); i < len(existingOptions); i++ {
				err := s.optRepo.Delete(ctx, existingOptions[i].Id)
				if err != nil {
					return nil, fmt.Errorf("failed to delete option id %d: %w", existingOptions[i].Id, err)
				}
			}
		}
	}

	result, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated data: %w", err)
	}

	return result, nil
}

func (s *questionService) GetMany(ctx context.Context, limit int, offset int) ([]model.Question, int64, error) {
	data, total, err := s.repo.GetMany(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all data: %w", err)
	}
	return data, total, nil
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

func (s *questionService) GetByExam(ctx context.Context, examId int, limit int, offset int) ([]model.Question, int64, error) {
	data, total, err := s.repo.GetByExam(ctx, examId, limit, offset)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, fmt.Errorf("question with id %d not found", examId)
		}
		return nil, 0, fmt.Errorf("data with exam id %d not found: %w", examId, err)
	}
	return data, total, nil
}

func (s *questionService) GetByCreatorId(ctx context.Context, creatorId int, limit int, offset int) ([]model.Question, int64, error) {
	data, total, err := s.repo.GetByCreatorId(ctx, creatorId, limit, offset)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, fmt.Errorf("creator with id %d not found", creatorId)
		}
		return nil, 0, fmt.Errorf("data with creator id %d not found: %w", creatorId, err)
	}
	return data, total, nil
}

func (s *questionService) GetByDifficult(ctx context.Context, diff string, limit int, offset int) ([]model.Question, int64, error) {
	data, total, err := s.repo.GetByDifficult(ctx, diff, limit, offset)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, fmt.Errorf("question with difficulty %s not found", diff)
		}
		return nil, 0, fmt.Errorf("data with difficulty %s not found: %w", diff, err)
	}
	return data, total, nil
}

func (s *questionService) GetBySubject(ctx context.Context, subjectId int, limit int, offset int) ([]model.Question, int64, error) {
	data, total, err := s.repo.GetBySubject(ctx, subjectId, limit, offset)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, fmt.Errorf("question with subject id %d not found", subjectId)
		}
		return nil, 0, fmt.Errorf("data with subjectId %d not found: %w", subjectId, err)
	}
	return data, total, nil
}
