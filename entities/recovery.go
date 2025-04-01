package entities

import (
	"gorm.io/gorm"
)

type Recovery struct {
	gorm.Model
	OTP    string `json:"otp"`
	UserID uint   `json:"user_id"`
	User   User   `json:"user" gorm:"foreignKey:UserID"`
}
