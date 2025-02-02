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

// CreateResult godoc
//
//	@Summary		Create  new result
//	@Description	Create  new result
//	@Tags			result
//	@Accept			json
//	@Produce		json
//	@Param			result	body		entities.Result	true	"Result information"
//	@Success		201		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		500		{object}	presentation.Responses
//	@Router			/result [post]
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

// GetResults godoc
//
//	@Summary		Get results
//	@Description	Get results
//	@Tags			result
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	presentation.Responses
//	@Failure		500	{object}	presentation.Responses
//	@Router			/result [get]
func (handler *HttpResultHandler) GetResults(c *fiber.Ctx) error {
	results, err := handler.resultUcase.GetResults()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ResultsResponse(results))
}

// GetResultById godoc
//
//	@Summary		Get a result by ID
//	@Description	Get a result by ID
//	@Tags			result
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Result ID"
//	@Success		200	{object}	presentation.Responses
//	@Failure		400	{object}	presentation.Responses
//	@Failure		404	{object}	presentation.Responses
//	@Router			/result/{id} [get]
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

// UpdateResultById godoc
//
//	@Summary		Update a result by ID
//	@Description	Update a result by ID
//	@Tags			result
//	@Accept			json
//	@Produce		json
//	@Param			result	body		entities.Result	true	"Result information"
//	@Param			id	path		int	true	"Result ID"
//	@Success		200		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		500		{object}	presentation.Responses
//	@Router			/admin/result/ [put]
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

// DeleteResultById godoc
//
//	@Summary		Delete a result by ID
//	@Description	Delete a result by ID
//	@Tags			result
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Result ID"
//	@Success		204	{object}	presentation.Responses
//	@Failure		400	{object}	presentation.Responses
//	@Failure		500	{object}	presentation.Responses
//	@Router			/result/{id} [delete]
func (handler *HttpResultHandler) DeleteResultById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	err = handler.resultUcase.DeleteResultById(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusNoContent).JSON(presentation.DeleteResponse(id))
}

// GetResultsByUserIdFromToken godoc
//
//		@Summary		Get results by user_id from token
//		@Description	Get results by user_id from token
//		@Tags			result
//		@Accept			json
//		@Produce		json
//		@Param			token	header		string	true	"User Bearer Token"
//		@Param			result	body		entities.Result	true	"Result information"
//		@Success		200		{object}	presentation.Responses
//		@Failure		500		{object}	presentation.Responses
//	   @Failure		401		{object}	presentation.Responses
//		@Router			/user/result [get]
func (handler *HttpResultHandler) GetResultsByUserIdFromToken(c *fiber.Ctx) error {
	token := c.Get("token")

	results, err := handler.resultUcase.GetResultsByUserIdFromToken(token)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ResultsResponse(results))
}

// GetResultsByUserId godoc
//
//	@Summary		Get a result by user ID
//	@Description	Get a result by user ID
//	@Tags			result
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		int	true	"User ID"
//	@Success		200	{object}	presentation.Responses
//	@Failure		400	{object}	presentation.Responses
//	@Failure		404	{object}	presentation.Responses
//	@Router			/result/user/{userId} [get]
func (handler *HttpResultHandler) GetResultsByUserId(c *fiber.Ctx) error {

	user_id, err := c.ParamsInt("userId")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	results, err := handler.resultUcase.GetResultsByUserId(user_id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ResultsResponse(results))
}

// GetLatestResultByUserIdFromToken godoc
//
//		@Summary		Get latest result by user_id from token
//		@Description	Get latest result by user_id from token
//		@Tags			result
//		@Accept			json
//		@Produce		json
//		@Param			token	header		string	true	"User Bearer Token"
//		@Success		200		{object}	presentation.Responses
//		@Failure		500		{object}	presentation.Responses
//	   @Failure		401		{object}	presentation.Responses
//		@Router			/user/result [get]
func (handler *HttpResultHandler) GetLatestResultByUserIdFromToken(c *fiber.Ctx) error {
	token := c.Get("token")

	result, err := handler.resultUcase.GetLatestResultByUserIdFromToken(token)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToResultResponse(result))
}
