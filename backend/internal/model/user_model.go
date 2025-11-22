package model

import "time"

type User struct {
	Id           int       `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name         string    `json:"name" validate:"required"`
	Nim          *string   `json:"nim" gorm:"unique"`
	Nip          *string   `json:"nip" gorm:"unique"`
	Nidn         *string   `json:"nidn" gorm:"unique"`
	Username     *string   `json:"username" gorm:"unique"`
	ImgUrl       string    `json:"img_url,omitempty"`
	Email        *string   `json:"email" gorm:"unique"`
	Password     string    `json:"password" validate:"required,min=6"`
	Role         Role      `json:"role" validate:"oneof=admin user super_admin lecturer"`
	Major        string    `json:"major,omitempty"`
	AcademicYear string    `json:"academic_year,omitempty" validate:"len=4"`
	Faculty      string    `json:"faculty,omitempty"`
	Status       Status    `json:"status" validate:"oneof=passed not_passed"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UpdateUser struct {
	Name         *string `json:"name" validate:"omitempty"`
	Nim          *string `json:"nim" gorm:"unique"`
	Nip          *string `json:"nip" gorm:"unique"`
	Nidn         *string `json:"nidn" gorm:"unique"`
	Username     *string `json:"username,omitempty" gorm:"unique"`
	ImgUrl       *string `json:"img_url,omitempty"`
	Email        *string `json:"email" gorm:"unique;"`
	Role         *Role   `json:"role,omitempty" validate:"omitempty,oneof=admin user super_admin lecturer"`
	Major        *string `json:"major,omitempty"`
	AcademicYear *string `json:"academic_year,omitempty" validate:"omitempty,len=4"`
	Faculty      *string `json:"faculty,omitempty"`
	Status       *Status `json:"status,omitempty" validate:"omitempty,oneof=passed not_passed"`

	ImgDelete *bool `json:"img_delete,omitempty" gorm:"default:false"`
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
	Username     string `json:"username,omitempty" gorm:"unique"`
	Major        string `json:"major" validate:"required" gorm:"default:informatika"`
	Faculty      string `json:"faculty" validate:"required" gorm:"default:teknik"`
	Role         Role   `json:"role" gorm:"default:user"`
	AcademicYear string `json:"academic_year,omitempty" validate:"omitempty,len=4"`
}

type ChangePasswordCredential struct {
	NewPassword string `json:"new_password" binding:"required"`
}

type ChangeRoleCredential struct {
	Nim          *string `json:"nim,omitempty"`
	Nip          *string `json:"nip,omitempty"`
	Nidn         *string `json:"nidn,omitempty"`
	Username     *string `json:"username,omitempty"`
	AcademicYear *string `json:"academic_year,omitempty"`
	Role         Role    `json:"role" binding:"required"`
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

type UserResponse struct {
	Id           int       `json:"id"`
	Name         *string   `json:"name,omitempty"`
	Nim          *string   `json:"nim,omitempty"`
	Nip          *string   `json:"nip,omitempty"`
	Nidn         *string   `json:"nidn,omitempty"`
	Username     *string   `json:"username,omitempty"`
	ImgUrl       *string   `json:"img_url,omitempty"`
	Email        *string   `json:"email,omitempty"`
	Major        *string   `json:"major,omitempty"`
	AcademicYear *string   `json:"academic_year,omitempty"`
	Faculty      *string   `json:"faculty,omitempty"`
	Status       Status    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Role         Role      `json:"role"`
}
