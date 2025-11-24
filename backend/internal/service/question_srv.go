package service

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"latih.in-be/internal/model"
	"latih.in-be/internal/repository"
	"latih.in-be/utils/helper"
	"latih.in-be/utils/update"
)

type QuestionService interface {
	Create(ctx context.Context, c *gin.Context, data *model.Question) error
	GetById(ctx context.Context, id int) (*model.Question, error)
	Update(ctx context.Context, c *gin.Context, req model.UpdateQuestion, id int, userId int) (*model.Question, error)
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

// optionRepoAdapter adapts repository.OptionRepository so its Update method
// matches the signature expected by update.OptionValidation (returns interface{}).
type optionRepoAdapter struct {
	r repository.OptionRepository
}

func (a optionRepoAdapter) Create(ctx context.Context, opt model.Option) error {
	return a.r.Create(ctx, opt)
}

func (a optionRepoAdapter) Delete(ctx context.Context, id int) error {
	return a.r.Delete(ctx, id)
}

func (a optionRepoAdapter) Update(ctx context.Context, opt model.Option, id int) (interface{}, error) {
	res, err := a.r.Update(ctx, opt, id)
	if err != nil {
		return nil, err
	}
	return res, nil
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

	node, err := snowflake.NewNode(1)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	id := node.Generate()
	imgDir := "./storages/images/question"
	imageURL, err := helper.UploadImage(c, int(id.Int64()), imgDir)
	if err != nil {
		return fmt.Errorf("failed to upload image: %w", err)
	}

	data.ImgUrl = imageURL

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

func (s *questionService) Update(ctx context.Context, c *gin.Context, req model.UpdateQuestion, id int, userId int) (*model.Question, error) {
	oldQuest, err := s.repo.GetById(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("question with id %d not found", id)
		}
		return nil, fmt.Errorf("data with id %d not found: %w", id, err)
	}

	user, err := s.userRepo.GetById(ctx, userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user with id %d not found", userId)
		}
		return nil, fmt.Errorf("user is unavailable: %w", err)
	}

	if user.Id != oldQuest.CreatorId && user.Role != model.RoleAdmin {
		return nil, fmt.Errorf("you are not the creator or admin")
	}

	if err = update.DifficultyValidation(req); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if err := update.HandleQuestionImageUpload(c, oldQuest, &req); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if req.CreatorId != nil && user.Role != model.RoleAdmin {
		return nil, fmt.Errorf("only admins can change the creator of a question")
	}

	adapter := optionRepoAdapter{r: s.optRepo}
	if err := update.OptionValidation(ctx, req, oldQuest, id, adapter); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	hasChanges := false
	if req.SubjectId != nil || req.CreatorId != nil || req.QuestionText != nil ||
		req.Difficulty != nil || req.Answer != nil || req.Score != nil ||
		req.ImgUrl != nil || len(req.Options) > 0 {
		hasChanges = true
	}

	if !hasChanges {
		return nil, fmt.Errorf("no fields to update")
	}

	updatedData, err := s.repo.Update(ctx, req, id)
	if err != nil {
		return nil, update.FormatUpdateQuestError(err, req)
	}

	return updatedData, nil
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
