package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model    `swaggerignore:"true"`
	FullName      string    `json:"full_name"`
	Email         string    `json:"email"`
	Birthday      time.Time `json:"birthday"`
	SensitiveSkin bool      `json:"sensitive_skin"`
	Password      string    `json:"password"`
	Image         string    `json:"image" swaggerignore:"true"`
}
