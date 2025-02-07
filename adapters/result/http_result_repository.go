package adapters

import (
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpResultHandler struct {
	resultUsecase usecases.ResultsUsecase
}

func NewHttpResultHandler(resultUcase usecases.ResultsUsecase) *HttpResultHandler {
	return &HttpResultHandler{resultUcase}
}

func (handler *HttpResultHandler) CreateResult(c *fiber.Ctx) error {

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	create_by_token := c.Get("token")

	result, err := handler.resultUsecase.CreateResult(*file, create_by_token, c)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}
