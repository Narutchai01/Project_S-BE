package adapters

import (
	"strconv"

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

	return c.Status(fiber.StatusCreated).JSON(presentation.ToResultResponse(result))
}

func (handler *HttpResultHandler) GetResults(c *fiber.Ctx) error {

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrBadRequest))
	}

	results, err := handler.resultUsecase.GetResults(token)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToResultsResponse(results))
}

func (handler *HttpResultHandler) GetResult(c *fiber.Ctx) error {
	id := c.Params("id")

	// Convert id from string to uint
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	result, err := handler.resultUsecase.GetResult(uint(uintID))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToResultResponse(result))
}
