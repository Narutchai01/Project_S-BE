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

// // GetResults godoc
// //
// //	@Summary		Get results
// //	@Description	Get results
// //	@Tags			results
// //	@Accept			json
// //	@Produce		json
// //	@Param			token	header		string	true	"Token"
// //	@Success		200		{object}	presentation.Responses
// //	@Failure		400		{object}	presentation.Responses
// //	@Failure		500		{object}	presentation.Responses
// //	@Router			/results/ [get]
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

// // GetResultLatest godoc
// //
// //	@Summary		Get the latest result
// //	@Description	Get the latest result
// //	@Tags			results
// //	@Accept			json
// //	@Produce		json
// //	@Param			token	header		string	true	"Token"
// //	@Success		200		{object}	presentation.Responses
// //	@Failure		400		{object}	presentation.Responses
// //	@Failure		500		{object}	presentation.Responses
// //	@Router			/results/latest [get]
func (handler *HttpResultHandler) GetResultLatest(c *fiber.Ctx) error {
	token := c.Get("token")

	result, err := handler.resultUsecase.GetResultLatest(token)

	if err != nil {
		if err.Error() == "result not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToResultResponse(result))
}

// // GetResult godoc
// //
// //	@Summary		Get a result
// //	@Description	Get a result
// //	@Tags			results
// //	@Accept			json
// //	@Produce		json
// //	@Param			id	path		string	true	"Result ID"
// //	@Success		200	{object}	presentation.Responses
// //	@Failure		400	{object}	presentation.Responses
// //	@Failure		500	{object}	presentation.Responses
// //	@Router			/results/{id} [get]
func (handler *HttpResultHandler) GetResult(c *fiber.Ctx) error {
	id := c.Params("id")

	// Convert id from string to uint
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	result, err := handler.resultUsecase.GetResult(uint(uintID))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToResultResponse(result))
}

// // GetResultByIDs godoc
// //
// //	@Summary		Get results by IDs
// //	@Description	Get results by IDs
// //	@Tags			results
// //	@Accept			json
// //	@Produce		json
// //	@Param			token	header		string	true	"Token"
// //	@Param			ids		body		object{IDs=[]uint}	true	"IDs"
// //	@Success		200		{object}	presentation.Responses
// //	@Failure		400		{object}	presentation.Responses
// //	@Failure		500		{object}	presentation.Responses
// //	@Router			/results/compare [post]
func (handler *HttpResultHandler) GetResultByIDs(c *fiber.Ctx) error {
	var ids struct{ IDs []uint }

	if err := c.BodyParser(&ids); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	if len(ids.IDs) <= 1 || ids.IDs == nil || len(ids.IDs) > 3 {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrBadRequest))
	}

	token := c.Get("token")

	results, err := handler.resultUsecase.GetResultByIDs(ids.IDs, token)
	if err != nil {
		if err.Error() == "results not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
		}
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToResultsResponse(results))
}
