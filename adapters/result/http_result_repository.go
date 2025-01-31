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

func (handler *HttpResultHandler) GetResultById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	result, err := handler.resultUcase.GetResultById(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToResultResponse(result))
}

func (handler *HttpResultHandler) UpdateResultById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	var new_result entities.Result

	if err := c.BodyParser(&new_result); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	result, err := handler.resultUcase.UpdateResultById(id, new_result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToResultResponse(result))
}

func (handler *HttpResultHandler) DeleteResultById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	err = handler.resultUcase.DeleteResultById(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.DeleteResponse(id))
}

