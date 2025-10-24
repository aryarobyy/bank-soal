package model

import (
	"time"
)

type Question struct {
	Id           int        `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	ExamId       int        `json:"exam_id"`
	SubjectId    int        `json:"subject_id" gorm:"index;not null"`
	Subject      Subject    `gorm:"foreignKey:SubjectId"`
	CreatorId    int        `json:"creator_id" gorm:"not null;index"`
	QuestionText string     `json:"question_text" validate:"required"`
	Difficulty   Difficulty `json:"difficulty" gorm:"type:enum('easy','medium','hard');not null" validate:"oneof=easy medium hard"`
	Answer       string     `json:"answer,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`

	Options []Option `json:"options" gorm:"foreignKey:QuestionId;constraint:OnDelete:CASCADE;"`
}
