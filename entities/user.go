package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model    `swaggerignore:"true"`
	FullName      string     `json:"full_name" gorm:"not null"`
	Email         string     `json:"email" gorm:"unique not null"`
	Birthday      *time.Time `json:"birthday" gorm:"default:null"`
	SensitiveSkin bool       `json:"sensitive_skin" gorm:"default:null"`
	Password      string     `json:"password"`
	Image         string     `json:"image" swaggerignore:"true"`
}
