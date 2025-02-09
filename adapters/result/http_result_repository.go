package adapters

import (
	"strconv"

	"github.com/Narutchai01/Project_S-BE/entities"
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

// CreateResult godoc
//
//	@Summary		Create a result
//	@Description	Create a result
//	@Tags			results
//	@Accept			json
//	@Produce		json
//	@Param			file	formData	file	true	"File"
//	@Param			token	header		string	true	"Token"
//	@Success		201		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		500		{object}	presentation.Responses
//	@Router			/results/ [post]
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

// GetResults godoc
//
//	@Summary		Get results
//	@Description	Get results
//	@Tags			results
//	@Accept			json
//	@Produce		json
//	@Param			token	header		string	true	"Token"
//	@Success		200		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		500		{object}	presentation.Responses
//	@Router			/results/ [get]
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

// GetResult godoc
//
//	@Summary		Get a result
//	@Description	Get a result
//	@Tags			results
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Result ID"
//	@Success		200		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		500		{object}	presentation.Responses
//	@Router			/results/{id} [get]
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

// GetResultLatest godoc
//
//	@Summary		Get the latest result
//	@Description	Get the latest result
//	@Tags			results
//	@Accept			json
//	@Produce		json
//	@Param			token	header	string	true	"Token"
//	@Success		200		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		500		{object}	presentation.Responses
//	@Router			/results/latest [get]
func (handler *HttpResultHandler) GetResultLatest(c *fiber.Ctx) error {
	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrBadRequest))
	}

	result, err := handler.resultUsecase.GetResultLatest(token)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToResultResponse(result))
}

func (handler *HttpResultHandler) UpdateResult(c *fiber.Ctx) error {
	id := c.Params("id")
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	var result entities.Result

	if err := c.BodyParser(&result); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	updatedResult, err := handler.resultUsecase.UpdateResult(result, uint(uintID))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToResultResponse(updatedResult))
}

func (handler *HttpResultHandler) DeleteResult(c *fiber.Ctx) error {
	id := c.Params("id")

	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	err = handler.resultUsecase.DeleteResult(uint(uintID))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.DeleteResponse(int(uintID)))
}
