package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName     string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}
