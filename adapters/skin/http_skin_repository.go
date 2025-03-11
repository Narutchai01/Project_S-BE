package adapters

import (
	"fmt"
	"strconv"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpSkinHandler struct {
	skinUsecase usecases.SkinUsecases
}

func NewHttpSkinHandler(skinUcase usecases.SkinUsecases) *HttpSkinHandler {
	return &HttpSkinHandler{skinUcase}
}

// CreateSkin godoc
//
//	@Summary		Create an skin
//	@Description	Create an skin
//	@Tags			skin
//	@Accept			json
//	@Produce		json
//	@Param			skin	formData	entities.Skin	true	"Skin Object"
//	@Param			file	formData	file			true	"Skin Image"
//	@Param			token	header		string			true	"Token"
//	@Success		201		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		404		{object}	presentation.Responses
//
//	@Router			/admin/skin [post]
func (handler *HttpSkinHandler) CreateSkin(c *fiber.Ctx) error {
	var skin entities.Skin

	if err := c.BodyParser(&skin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	create_by_token := c.Get("token")

	if skin.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fmt.Errorf("name is required")))
	}

	result, err := handler.skinUsecase.CreateSkin(skin, *file, c, create_by_token)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(presentation.ToSkinResponse(result))
}

// GetSkins godoc
//
//	@Summary		Get skins
//	@Description	Get skins
//	@Tags			skin
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	presentation.Responses
//	@Failure		400	{object}	presentation.Responses
//	@Failure		404	{object}	presentation.Responses
//
//	@Router			/skin [get]
func (handler *HttpSkinHandler) GetSkins(c *fiber.Ctx) error {
	skins, err := handler.skinUsecase.GetSkins()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToSkinsResponse(skins))
}

// GetSkin godoc
//
//	@Summary		Get skin
//	@Description	Get skin
//	@Tags			skin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Skin ID"
//	@Success		200	{object}	presentation.Responses
//	@Failure		400	{object}	presentation.Responses
//	@Failure		404	{object}	presentation.Responses
//
//	@Router			/skin/{id} [get]
func (handler *HttpSkinHandler) GetSkin(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	skin, err := handler.skinUsecase.GetSkin(intID)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToSkinResponse(skin))
}

// UpdateSkin godoc
//
//	@Summary		Update skin
//	@Description	Update skin
//	@Tags			skin
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string			true	"Skin ID"
//	@Param			skin	formData	entities.Skin	true	"Skin Object"
//	@Param			file	formData	file			false	"Skin Image"
//	@Success		200		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		404		{object}	presentation.Responses
//
//	@Router			/admin/skin/{id} [put]
func (handler *HttpSkinHandler) UpdateSkin(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	var skin entities.Skin

	if err := c.BodyParser(&skin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	file, _ := c.FormFile("file")

	result, err := handler.skinUsecase.UpdateSkin(intID, skin, file, c)
	if err != nil {
		if err.Error() == "facial not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
		}
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToSkinResponse(result))
}

// DeleteSkin godoc
//
//	@Summary		Delete skin
//	@Description	Delete skin
//	@Tags			skin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Skin ID"
//	@Success		204	{object}	presentation.Responses
//	@Failure		400	{object}	presentation.Responses
//	@Failure		404	{object}	presentation.Responses
//
//	@Router			/admin/skin/{id} [delete]
func (handler *HttpSkinHandler) DeleteSkin(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	err = handler.skinUsecase.DeleteSkin(intID)

	if err != nil {
		if err.Error() == "facial not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))

		}
	}

	return c.Status(fiber.StatusOK).JSON(presentation.DeleteResponse(intID))
}
