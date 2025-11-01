package model

import "time"

type User struct {
	Id           int       `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name         string    `json:"name" validate:"required"`
	Nim          string    `json:"nim,omitempty" gorm:"unique" validate:"required"`
	ImgUrl       string    `json:"img_url,omitempty"`
	Email        string    `json:"email" gorm:"unique" validate:"required,email"`
	Password     string    `json:"password" validate:"required,min=6"`
	Role         Role      `json:"role" validate:"oneof=admin user super_admin lecturer"`
	Major        string    `json:"major,omitempty"`
	AcademicYear int       `json:"academic_year,omitempty" validate:"numeric,len=4"`
	Faculty      string    `json:"faculty,omitempty"`
	Status       Status    `json:"status" validate:"oneof=passed not_passed"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type LoginCredential struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterCredential struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required"`
	Nim      string `json:"nim" validate:"required"`
	Major    string `json:"major" validate:"required"`
	Faculty  string `json:"faculty" validate:"required"`
	Role     string `json:"role" gorm:"default:user"`
}

type ChangePasswordCredential struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type ChangeRoleCredential struct {
	Role Role `json:"role" binding:"required"`
}
