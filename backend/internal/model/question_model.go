package model

import (
	"time"
)

type Question struct {
	Id           int        `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	SubjectId    int        `json:"subject_id" gorm:"index;not null"`
	Subject      Subject    `gorm:"foreignKey:SubjectId"`
	CreatorId    int        `json:"creator_id" gorm:"not null;index"`
	QuestionText string     `json:"question_text" validate:"required"`
	Difficulty   Difficulty `json:"difficulty" gorm:"type:enum('easy','medium','hard');not null;index" validate:"oneof=easy medium hard"`
	Answer       string     `json:"answer,omitempty"`
	Score        int        `json:"score" gorm:"not null;default:0"`
	ImgUrl       string     `json:"img_url,omitempty"`
	CreatedAt    time.Time  `json:"created_at" gorm:"index"`
	UpdatedAt    time.Time  `json:"updated_at"`

	Options []Option `json:"options" gorm:"foreignKey:QuestionId;constraint:OnDelete:CASCADE;" validate:"required"`
	Exams   []Exam   `json:"exams" gorm:"many2many:exam_questions;"`
}

type UpdateQuestion struct {
	SubjectId    *int        `json:"subject_id,omitempty"`
	CreatorId    *int        `json:"creator_id,omitempty"` //(Gatau Kepake gak)
	QuestionText *string     `json:"question_text,omitempty"`
	Difficulty   *Difficulty `json:"difficulty,omitempty"`
	Answer       *string     `json:"answer,omitempty"`
	Score        *int        `json:"score,omitempty"`
	ImgUrl       *string     `json:"img_url,omitempty"`

	Options   []Option `json:"options" gorm:"foreignKey:QuestionId;constraint:OnDelete:CASCADE;"`
	ImgDelete *bool    `json:"img_delete,omitempty" gorm:"default:false"`
}

type QuestionResponse struct {
	Id           int        `json:"id"`
	SubjectId    int        `json:"subject_id"`
	Subject      Subject    `json:"subject"`
	CreatorId    int        `json:"creator_id"`
	QuestionText string     `json:"question_text"`
	Difficulty   Difficulty `json:"difficulty"`
	Answer       string     `json:"answer,omitempty"`
	Score        int        `json:"score"`
	ImgUrl       string     `json:"img_url,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Options      []Option   `json:"options"`
}
