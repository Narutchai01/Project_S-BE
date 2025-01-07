package entities

import (
	"gorm.io/gorm"
)

type Recovery struct {
	gorm.Model
	OTP string `json:"otp"`
	UserId    uint `json:"user_id"`
}