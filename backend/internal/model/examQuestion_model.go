package model

type ExamQuestion struct {
	ExamId     int `json:"exam_id" gorm:"primaryKey"`
	QuestionId int `json:"question_id" gorm:"primaryKey"`

	Questions Question `gorm:"foreignKey:QuestionId;references:Id"`
}
