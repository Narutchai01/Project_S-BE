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
