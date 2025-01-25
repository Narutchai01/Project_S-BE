package entities

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model `swaggerignore:"true"`
	FullName   string `json:"fullname" gorm:"not null"`
	Email      string `json:"email" gorm:"unique not null"`
	Password   string `json:"password"`
	Image      string `json:"image" swaggerignore:"true"`
}
