package seeder

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"latih.in-be/internal/model"
)

func SeedUser(db *gorm.DB) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.DefaultCost)

	users := []model.User{
		{
			Name:     "Super Admin",
			Role:     model.RoleSuperAdmin,
			Username: func() *string { s := "ilhamgod14"; return &s }(),
			Password: string(hashedPassword),
		},
		{
			Name:     "Admin",
			Role:     model.RoleAdmin,
			Username: func() *string { s := "ilhamgod15"; return &s }(),
			Password: string(hashedPassword),
		},
	}

	for _, u := range users {
		db.Where("username = ?", u.Username).FirstOrCreate(&u)
	}
	return nil
}
