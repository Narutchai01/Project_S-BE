package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Birthday    string `json:"birthday"`
	SensitiveSkin  string `json:"sensitive_skin"`
	Password string `json:"password"`
	Image    string `json:"image"`
}