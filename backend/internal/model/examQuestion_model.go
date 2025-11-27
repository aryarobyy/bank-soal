package model

type ExamQuestion struct {
	ExamId     int `json:"exam_id" gorm:"primaryKey;constraint:OnDelete:CASCADE"`
	QuestionId int `json:"question_id" gorm:"primaryKey;constraint:OnDelete:CASCADE"`

	Question *Question `json:"question,omitempty" gorm:"foreignKey:QuestionId;references:Id;constraint:OnDelete:CASCADE"`
}
