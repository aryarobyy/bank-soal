package model

import "time"

type Exam struct {
	Id          int        `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Title       string     `json:"title" validate:"required"`
	Description string     `json:"description,omitempty"`
	Difficulty  Difficulty `json:"difficulty" gorm:"type:enum('easy','medium','hard');not null" validate:"oneof=easy medium hard"`
	LongTime    int        `json:"long_time" validate:"required,min=1"` // menit
	CreatorId   int        `json:"creator_id" validate:"required"`
	StartedAt   *time.Time `json:"started_at" validate:"required"`
	FinishedAt  *time.Time `json:"finished_at" validate:"required, gtfield=StartedAt"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
