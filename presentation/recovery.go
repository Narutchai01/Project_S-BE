package presentation

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
)

type Recovery struct {
	ID     uint   `json:"id"`
	UserId int   `json:"user_id"`
	OTP    string `json:"otp"`
}

func RecoveryResponse(data entities.Recovery) *fiber.Map {
	recovery := Recovery{
		ID:     data.ID,
		UserId: data.UserId,
		OTP:    data.OTP,
	}

	return &fiber.Map{
		"status":   true,
		"recovery": recovery,
		"error":    nil,
	}
}

func RecoveriesResponse(data []entities.Recovery) *fiber.Map {
	recoveries := []Recovery{}

	for _, recovery := range data {
		recoveries = append(recoveries, Recovery{
			ID:     recovery.ID,
			UserId: recovery.UserId,
			OTP:    recovery.OTP,
		})
	}
	return &fiber.Map{
		"status": true,
		"data":   recoveries,
		"error":  nil,
	}
}

func RecoveryErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status":   false,
		"recovery": nil,
		"error":    err.Error(),
	}
}

func DeleteRecoveryResponse(id int) *fiber.Map {
	return &fiber.Map{
		"status":    true,
		"delete_id": id,
		"error":     nil,
	}
}
