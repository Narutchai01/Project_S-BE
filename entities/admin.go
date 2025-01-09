package entities

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model `swaggerignore:"true"`
	FullName   string `json:"full_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Image      string `json:"image" swaggerignore:"true"`
}
