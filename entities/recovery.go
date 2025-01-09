package entities

import (
	"gorm.io/gorm"
)

type Recovery struct {
	gorm.Model
	OTP string `json:"otp"`
	UserId    int `json:"user_id"`
}