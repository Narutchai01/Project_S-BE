package entities

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Image    string `json:"image"`
}
