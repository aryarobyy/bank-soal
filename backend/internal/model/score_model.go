package model

import "time"

type Score struct {
	Id         int       `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	ExamId     int       `json:"exam_id" gorm:"not null"`
	UserId     int       `json:"user_id" gorm:"not null"`
	QuestionId int       `json:"question_id" gorm:"not null"`
	OptionId   int       `json:"option_id" gorm:"not null"`
	IsCorrect  bool      `json:"is_correct"`
	Score      int       `json:"score"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
