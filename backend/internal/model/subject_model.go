package model

type Subject struct {
	Id    int          `json:"id" gorm:"primaryKey;autoIncrement"`
	Title SubjectTitle `json:"title" gorm:"type:varchar(255);not null"`
	Code  string       `json:"code" gorm:"type:varchar(50);uniqueIndex;not null"`
}
