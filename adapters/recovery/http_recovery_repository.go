package adapters

import (
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
	var request struct {
		Email string `json:"email"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	recovery, err := handler.recoveryUcase.CreateRecovery(request.Email)
	if err != nil {
		if err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.RecoveryResponse(recovery))

}

func (handler *HttpRecoveryHandler) ValidateRecovery(c *fiber.Ctx) error {
	var request struct {
		OTP    string `json:"otp"`
		UserID uint   `json:"user_id"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	recovery, err := handler.recoveryUcase.ValidateRecovery(request.OTP, request.UserID)
	if err != nil {
		if err.Error() == "invalid OTP" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.RecoveryResponse(recovery))
}

func (handler *HttpRecoveryHandler) ResetPassword(c *fiber.Ctx) error {
	var request struct {
		NewPassword string `json:"new_password"`
		UserID      uint   `json:"user_id"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	user, err := handler.recoveryUcase.ResetPassword(request.NewPassword, request.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.UserResponse(user))
}
