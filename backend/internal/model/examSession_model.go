package model

import "time"

type ExamSession struct {
	Id         int           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserId     int           `json:"user_id" gorm:"not null; index"`
	ExamId     int           `json:"exam_id" gorm:"not null; index"`
	StartedAt  time.Time     `json:"started_at" gorm:"autoCreateTime"`
	FinishedAt *time.Time    `json:"finished_at"`
	Status     SessionStatus `json:"status" gorm:"default:'in_progress'"`
	CurrentNo  int           `json:"current_no" gorm:"default:1"`
	Score      int           `json:"score"`
	MaxScore   int           `json:"max_score"`
	Percentage float64       `json:"percentage" gorm:"default:0.0"`
	IsPassed   bool          `json:"is_passed" gorm:"default:false"`

	CreatedAt time.Time `json:"created_at" gorm:"index"`
	UpdatedAt time.Time `json:"updated_at"`

	UserAnswers []UserAnswer `gorm:"foreignKey:ExamSessionId"`
}

type UpdateExamSession struct {
	UserId     int           `json:"user_id" gorm:"not null"`
	ExamId     int           `json:"exam_id" gorm:"not null"`
	StartedAt  time.Time     `json:"started_at" gorm:"not null"`
	FinishedAt *time.Time    `json:"finished_at"`
	Status     SessionStatus `json:"status" gorm:"default:'in_progress'"`
	CurrentNo  int           `json:"current_no" gorm:"not null"`
	IsPassed   *bool         `json:"is_passed"`
	Score      int           `json:"score"`
}

type UpdateCurrNo struct {
	CurrentNo int `json:"current_no" gorm:"not null"`
}

type UpdateStatus struct {
	Status SessionStatus `json:"status" gorm:"not null"`
}

type FinishExam struct {
	Id         int           `json:"session_id" gorm:"not null"` //id
	UserId     int           `json:"user_id" gorm:"not null"`
	ExamId     int           `json:"exam_id" gorm:"not null"`
	FinishedAt time.Time     `json:"finished_at"`
	Status     SessionStatus `json:"status" gorm:"default:'finished'"`
	Score      int           `json:"score"`
	MaxScore   int           `json:"max_score"`
	Percentage float64       `json:"percentage"`
	IsPassed   bool          `json:"is_passed"`
}

type SessionResponse struct {
	Id         int           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserId     int           `json:"user_id" gorm:"not null"`
	ExamId     int           `json:"exam_id" gorm:"not null"`
	StartedAt  time.Time     `json:"started_at" gorm:"not null"`
	FinishedAt *time.Time    `json:"finished_at"`
	Status     SessionStatus `json:"status" gorm:"default:'in_progress'"`
	CurrentNo  int           `json:"current_no" gorm:"not null"`
	Score      *float64      `json:"score"`
	MaxScore   int           `json:"max_score"`
	Percentage float64       `json:"percentage"`
	IsPassed   bool          `json:"is_passed" gorm:"default:false"`

	UserAnswers []UserAnswer `gorm:"foreignKey:ExamSessionId"`
}

type FinishexamResponse struct {
	StartedAt  time.Time     `json:"started_at"`
	FinishedAt time.Time     `json:"finished_at"`
	Status     SessionStatus `json:"status" gorm:"default:'finished'"`
	MaxScore   int           `json:"max_score"`
	Score      int           `json:"score"`
	Percentage float64       `json:"percentage"`
}

type ScoreResponse struct {
	MaxScore   int     `json:"max_score"`
	Score      int     `json:"score"`
	Percentage float64 `json:"percentage"`
}
