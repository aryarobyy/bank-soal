package model

import "time"

type UserAnswer struct {
	Id            int         `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	ExamSessionId int         `json:"exam_session_id" gorm:"not null"`
	ExamSession   ExamSession `gorm:"foreignKey:ExamSessionId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	QuestionId    int         `json:"question_id" gorm:"not null"`
	Answer        string      `json:"answer" gorm:"not null"`
	IsCorrect     *bool       `json:"is_correct"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
