package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpSkincareHandler struct {
	skincarenUcase usecases.SkincareUsecases
}

func NewHttpSkincareHandler(skincareUcase usecases.SkincareUsecases) *HttpSkincareHandler {
	return &HttpSkincareHandler{skincareUcase}
}

// CreateSkincare godoc
//
//	@Summary		Create a skincare
//	@Description	Create a skincare
//	@Tags			skincare
//	@Accept			json
//	@Produce		json
//	@Param			skincare	formData	entities.Skincare	true	"Skincare Object"
//	@Param			file		formData	file				true	"Skincare Image"
//
//	@Param			token		header		string				true	"Token"
//
//	@Success		201			{object}	presentation.Responses
//	@Failure		400			{object}	presentation.Responses
//	@Failure		404			{object}	presentation.Responses
//	@Router			/admin/skincare [post]
func (handler *HttpSkincareHandler) CreateSkincare(c *fiber.Ctx) error {
	var skincare entities.Skincare

	if err := c.BodyParser(&skincare); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}
	create_by_token := c.Get("token")

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	result, err := handler.skincarenUcase.CreateSkincare(skincare, *file, create_by_token, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(presentation.SkincareResponse(result))
}

// GetSkincares godoc
//
//	@Summary		Get skincares
//	@Description	Get skincares
//	@Tags			skincare
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	presentation.Responses
//	@Failure		400	{object}	presentation.Responses
//	@Failure		404	{object}	presentation.Responses
//	@Router			/skincare [get]
func (handler *HttpSkincareHandler) GetSkincares(c *fiber.Ctx) error {
	skincares, err := handler.skincarenUcase.GetSkincares()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.SkincaresResponse(skincares))
}

// GetSkincareById godoc
//
//	@Summary		Get a skincare
//	@Description	Get a skincare
//	@Tags			skincare
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Skincare ID"
//	@Success		200	{object}	presentation.Responses
//	@Failure		400	{object}	presentation.Responses
//	@Failure		404	{object}	presentation.Responses
//	@Router			/skincare/{id} [get]
func (handler *HttpSkincareHandler) GetSkincareById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	skincare, err := handler.skincarenUcase.GetSkincareById(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.SkincareResponse(skincare))
}

// UpdateSkincareById godoc
//
//	@Summary		Update a skincare
//	@Description	Update a skincare
//	@Tags			skincare
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int					true	"Skincare ID"
//	@Param			skincare	formData	entities.Skincare	true	"Skincare Object"
//	@Param			file		formData	file				false	"Skincare Image"
//	@Success		200			{object}	presentation.Responses
//	@Failure		400			{object}	presentation.Responses
//	@Failure		404			{object}	presentation.Responses
//	@Router			/admin/skincare/{id} [put]
func (handler *HttpSkincareHandler) UpdateSkincareById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	var skincare entities.Skincare

	if err := c.BodyParser(&skincare); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	file, err := c.FormFile("file")
	if err != nil {
		file = nil
	}

	result, err := handler.skincarenUcase.UpdateSkincareById(id, skincare, file, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.SkincareResponse(result))
}

// DeleteSkincareById godoc
//
//	@Summary		Delete a skincare
//	@Description	Delete a skincare
//	@Tags			skincare
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Skincare ID"
//	@Success		204	{object}	presentation.Responses
//	@Failure		400	{object}	presentation.Responses
//	@Failure		404	{object}	presentation.Responses
//	@Router			/admin/skincare/{id} [delete]
func (handler *HttpSkincareHandler) DeleteSkincareById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	_, err = handler.skincarenUcase.DeleteSkincareById(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.DeleteResponse(id))
}
