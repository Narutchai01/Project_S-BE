package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpRecoveryHandler struct {
	recoveryUcase usecases.RecoveryUsecases
}

func NewHttpRecoveryHandler(recoveryUcase usecases.RecoveryUsecases) *HttpRecoveryHandler {
	return &HttpRecoveryHandler{recoveryUcase}
}

func (handler *HttpRecoveryHandler) CreateRecovery(c *fiber.Ctx) error {
	type RequestBody struct {
		Email  string `json:"email"`
		UserId uint   `json:"user_id"`
		OTP    string `json:"otp"`
	}

	var requestBody RequestBody

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	recovery := entities.Recovery{
		UserId: requestBody.UserId,
		OTP:    requestBody.OTP,
	}

	result, err := handler.recoveryUcase.CreateRecovery(recovery, requestBody.Email, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.RecoveryErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}
