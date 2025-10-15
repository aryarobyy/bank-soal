package model

import "time"

type Option struct {
	Id          int       `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	QuestionId  int       `json:"question_id" gorm:"not null;index"`
	OptionLabel string    `json:"option_label" validate:"required"`
	OptionText  string    `json:"option_text" validate:"required"`
	IsCorrect   bool      `json:"is_correct"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
