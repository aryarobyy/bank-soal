package model

import "time"

type User struct {
	Id           int       `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name         string    `json:"name" validate:"required"`
	Nim          *string   `json:"nim,omitempty" gorm:"unique"`
	Nip          *string   `json:"nip,omitempty" gorm:"unique"`
	Nidn         *string   `json:"nidn,omitempty" gorm:"unique"`
	Username     *string   `json:"username,omitempty" gorm:"unique"`
	ImgUrl       string    `json:"img_url,omitempty"`
	Email        string    `json:"email,omitempty" gorm:"unique;default:null"`
	Password     string    `json:"password" validate:"required,min=6"`
	Role         Role      `json:"role" validate:"oneof=admin user super_admin lecturer"`
	Major        string    `json:"major,omitempty"`
	AcademicYear string    `json:"academic_year,omitempty" validate:"len=4"`
	Faculty      string    `json:"faculty,omitempty"`
	Status       Status    `json:"status" validate:"oneof=passed not_passed"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type LoginCredential struct {
	LoginId  string `json:"login_id"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterCredential struct {
	Email        string `json:"email,omitempty" gorm:"unique;default:null"`
	Password     string `json:"password" validate:"required,min=6"`
	Name         string `json:"name" validate:"required"`
	Nim          string `json:"nim,omitempty"`
	Nip          string `json:"nip,omitempty"`
	Nidn         string `json:"nidn,omitempty"`
	Username     string `json:"username,omitempty"`
	Major        string `json:"major" validate:"required" gorm:"default:informatika"`
	Faculty      string `json:"faculty" validate:"required" gorm:"default:teknik"`
	Role         Role   `json:"role" gorm:"default:user"`
	AcademicYear string `json:"academic_year,omitempty" validate:"omitempty,len=4"`
}

type ChangePasswordCredential struct {
	NewPassword string `json:"new_password" binding:"required"`
}

type ChangeRoleCredential struct {
	Role Role `json:"role" binding:"required"`
}

type BulkUserCredential struct {
	Nim          string `json:"nim"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	Major        string `json:"major"`
	Faculty      string `json:"faculty"`
	AcademicYear string `json:"academic_year"`
}

type BulkUserOutput struct {
	Nim      string `json:"nim"`
	Password string `json:"password"`
}
