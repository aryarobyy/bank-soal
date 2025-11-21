package model

import "time"

type ExamScore struct {
	Id         int       `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	ExamId     int       `json:"exam_id" gorm:"not null"`
	UserId     int       `json:"user_id" gorm:"not null"`
	TotalScore int       `json:"total_score"`
	Status     Status    `json:"status" gorm:"default:'not_passed'" validate:"oneof=passed not_passed"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ExamScoreResponse struct {
	Id         int       `json:"id"`
	ExamId     int       `json:"exam_id"`
	UserId     int       `json:"user_id"`
	TotalScore int       `json:"total_score"`
	Status     Status    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
