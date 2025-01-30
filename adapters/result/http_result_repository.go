package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpResultHandler struct {
	resultUcase usecases.ResultUsecases
}

func NewHttpResultHandler(resultUcase usecases.ResultUsecases) *HttpResultHandler {
	return &HttpResultHandler{resultUcase}
}

func (handler *HttpResultHandler) CreateResult(c *fiber.Ctx) error {

	var result entities.Result

	if err := c.BodyParser(&result); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	createdResult, err := handler.resultUcase.CreateResult(result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	response := presentation.ToResultResponse(createdResult)

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (handler *HttpResultHandler) GetResults(c *fiber.Ctx) error {
	results, err := handler.resultUcase.GetResults()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ResultsResponse(results))
}
