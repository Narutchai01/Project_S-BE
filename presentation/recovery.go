package presentation

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
)

type Recovery struct {
	ID       uint   `json:"id"`
	UserId uint `json:"user_id"`
	OTP    string `json:"otp"`
}

func RecoveryResponse(data entities.Recovery) *fiber.Map {
	recovery := Recovery{
		ID:       data.ID,
		UserId: data.UserId,
		OTP :    data.OTP ,
	}

	return &fiber.Map{
		"status": true,
		"recovery":  recovery,
		"error":  nil,
	}
}

func RecoveryErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"recovery":  nil,
		"error":  err.Error(),
	}
}
