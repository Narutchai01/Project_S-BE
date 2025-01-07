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
		UserId int   `json:"user_id"`
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

	oldRecovery, err := handler.recoveryUcase.GetRecoveryByUserId(recovery.UserId)
	if (err == nil) || (oldRecovery != entities.Recovery{}) {
		updateOtp, err := handler.recoveryUcase.UpdateRecoveryOtpById(oldRecovery, requestBody.Email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presentation.RecoveryErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(updateOtp)
	}

	result, err := handler.recoveryUcase.CreateRecovery(recovery, requestBody.Email, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.RecoveryErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func (handler *HttpRecoveryHandler) DeleteRecoveryById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	_, err = handler.recoveryUcase.DeleteRecoveryById(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.RecoveryErrorResponse(err))
	}

	return c.Status(fiber.StatusNoContent).JSON(presentation.DeleteRecoveryResponse(id))
}

func (handler *HttpRecoveryHandler) GetRecoveries(c *fiber.Ctx) error {
	recoveries, err := handler.recoveryUcase.GetRecoveries()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.RecoveryErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.RecoveriesResponse(recoveries))
}

func (handler *HttpRecoveryHandler) OtpValidation(c *fiber.Ctx) error {
	var recovery entities.Recovery
	if err := c.BodyParser(&recovery); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	isValid, err := handler.recoveryUcase.OtpValidation(int(recovery.ID), recovery.OTP)
	if err != nil || !isValid {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "invalid OTP",
		})
	}

	_, err = handler.recoveryUcase.DeleteRecoveryById(int(recovery.ID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.RecoveryErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.RecoveryResponse(recovery))
}