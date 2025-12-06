package update

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/model"
	"latih.in-be/utils/helper"
)

func DifficultyValidation(question model.UpdateQuestion) error {
	if question.Difficulty != nil {
		scoreToCheck := *question.Score
		if question.Score != nil {
			scoreToCheck = *question.Score
		}

		switch *question.Difficulty {
		case model.DifficultyEasy:
			if scoreToCheck < 3 || scoreToCheck > 8 {
				return fmt.Errorf("score for easy difficulty must be between 3 and 8, got %d", scoreToCheck)
			}
		case model.DifficultyMedium:
			if scoreToCheck < 10 || scoreToCheck > 15 {
				return fmt.Errorf("score for medium difficulty must be between 10 and 15, got %d", scoreToCheck)
			}
		case model.DifficultyHard:
			if scoreToCheck < 18 || scoreToCheck > 23 {
				return fmt.Errorf("score for hard difficulty must be between 18 and 23, got %d", scoreToCheck)
			}
		default:
			return fmt.Errorf("invalid difficulty level")
		}
	}
	return nil
}

func HandleQuestionImageUpload(c *gin.Context, oldQuest *model.Question, data *model.UpdateQuestion) error {
	if data.ImgDelete != nil && *data.ImgDelete {
		if oldQuest.ImgUrl != "" {
			if err := helper.DeleteImage(oldQuest.ImgUrl); err != nil {
				return fmt.Errorf("failed to delete image: %w", err)
			}
		}
		emptyUrl := ""
		data.ImgUrl = &emptyUrl
		return nil
	}

	file, _ := c.FormFile("image")
	if file == nil {
		data.ImgUrl = &oldQuest.ImgUrl
		return nil
	}

	if oldQuest.ImgUrl != "" {
		if err := helper.DeleteImage(oldQuest.ImgUrl); err != nil {
			return fmt.Errorf("failed to delete old image: %w", err)
		}
	}

	node, err := snowflake.NewNode(1)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	id := node.Generate()

	imgDir := "./storages/images/question"
	newImageUrl, err := helper.UploadImage(c, int(id.Int64()), imgDir)
	if err != nil {
		return fmt.Errorf("failed to upload image: %w", err)
	}

	data.ImgUrl = &newImageUrl
	return nil
}

func OptionValidation(ctx context.Context,
	question model.UpdateQuestion,
	oldQuest *model.Question,
	id int,
	s interface {
		Update(context.Context, model.Option, int) (interface{}, error)
		Create(context.Context, model.Option) error
		Delete(context.Context, int) error
	}) error {
	existingOptions := oldQuest.Options

	minLen := len(question.Options)
	if len(existingOptions) < minLen {
		minLen = len(existingOptions)
	}

	correctCount := 0
	for i := 0; i < minLen; i++ {
		optToUpdate := question.Options[i]
		optToUpdate.Id = existingOptions[i].Id
		optToUpdate.QuestionId = id

		if optToUpdate.IsCorrect {
			correctCount++
		}

		_, err := s.Update(ctx, optToUpdate, existingOptions[i].Id)
		if err != nil {
			return fmt.Errorf("failed to update option at index %d: %w", i, err)
		}
	}
	fmt.Println("skdsksalsal", correctCount)
	if correctCount != 1 {
		return fmt.Errorf("exactly one option must be correct")
	}

	if len(question.Options) > len(existingOptions) {
		for i := len(existingOptions); i < len(question.Options); i++ {
			newOpt := question.Options[i]
			newOpt.QuestionId = id
			err := s.Create(ctx, newOpt)
			if err != nil {
				return fmt.Errorf("failed to create option at index %d: %w", i, err)
			}
		}
	}

	if len(existingOptions) > len(question.Options) {
		for i := len(question.Options); i < len(existingOptions); i++ {
			err := s.Delete(ctx, existingOptions[i].Id)
			if err != nil {
				return fmt.Errorf("failed to delete option id %d: %w", existingOptions[i].Id, err)
			}
		}
	}
	return nil
}

func FormatUpdateQuestError(err error, data model.UpdateQuestion) error {
	if strings.Contains(err.Error(), "Unknown column") {
		parts := strings.Split(err.Error(), "'")
		if len(parts) >= 2 {
			fieldName := parts[1]
			val := helper.GetFieldValue(data, fieldName)
			return fmt.Errorf("field '%s' with value '%v' is undefined", fieldName, val)
		}
	}
	return fmt.Errorf("update gagal: %v", err)
}
