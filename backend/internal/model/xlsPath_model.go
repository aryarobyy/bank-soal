package model

import "time"

type XlsPath struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	FilePath  string `json:"file_path" gorm:"type:varchar(500)"`
	CreatedAt time.Time
}
