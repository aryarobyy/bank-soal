package model

import "time"

type ExamScore struct {
	Id         int    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	ExamId     int    `json:"exam_id" gorm:"not null; index"`
	UserId     int    `json:"user_id" gorm:"not null; index"`
	TotalScore int    `json:"total_score"`
	Status     Status `json:"status" gorm:"default:'not_passed'" validate:"oneof=passed not_passed"`

	User      *User     `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at" gorm:"index"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ExamScoreResponse struct {
	Id         int    `json:"id"`
	ExamId     int    `json:"exam_id"`
	UserId     int    `json:"user_id"`
	TotalScore int    `json:"total_score"`
	Status     Status `json:"status"`

	User *User `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
