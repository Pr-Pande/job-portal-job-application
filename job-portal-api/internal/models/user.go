package models

import "gorm.io/gorm"

//signup validation purpose
type NewUser struct {
	Name     string `json:"Name" validate:"required"`
	Email    string `json:"Email" validate:"required,email"`
	Password string `json:"Password" validate:"required"`
}

//store to databas
type User struct {
	gorm.Model
	Name         string `json:"Name"`
	Email        string `json:"Email"`
	PasswordHash string `json:"-"`
}
