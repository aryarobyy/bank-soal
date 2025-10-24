package seeder

import (
	"fmt"

	"gorm.io/gorm"
	"latih.in-be/model"
)

func SeedSubjects(db *gorm.DB) error {
	subjects := []model.Subject{
		{Title: "Kalkulus", Code: "MFG-101"},
		{Title: "Matematika Diskrit", Code: "TIF-1203"},
		{Title: "Teori Bahasa dan Automata", Code: "TIF-2204"},
		{Title: "Basis Data Lanjut", Code: "TIF-2206"},
		{Title: "Metode Numerik", Code: "TIF-3107"},
	}

	for _, subject := range subjects {
		var existing model.Subject
		if err := db.Where("code = ?", subject.Code).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&subject).Error; err != nil {
				return fmt.Errorf("failedinsert subject %s: %v", subject.Title, err)
			}
			fmt.Println("Inserted:", subject.Title)
		} else {
			fmt.Println("Skip:", subject.Title, "(sudah ada)")
		}
	}
	return nil
}
