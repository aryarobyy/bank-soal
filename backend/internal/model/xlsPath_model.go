package model

import "time"

type XlsPath struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	FilePath  string `json:"file_path" gorm:"type:varchar(500)"`
	CreatedAt time.Time
}

type XlsPathResponse struct {
	Id        uint      `json:"id"`
	FilePath  string    `json:"file_path"`
	CreatedAt time.Time `json:"created_at"`
}