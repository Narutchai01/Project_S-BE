package adapters

import (
	"errors"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpSkinHandler struct {
	faceProblemsUsecase usecases.FaceProblemUseCase
}

func NewHttpSkinHandler(faceProblemsUsecase usecases.FaceProblemUseCase) *HttpSkinHandler {
	return &HttpSkinHandler{faceProblemsUsecase}
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
	var skin entities.FaceProblem

	if err := c.BodyParser(&skin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("file is required")))
	}

	create_by_token := c.Get("token")

	if create_by_token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(fiber.ErrUnauthorized))
	}

	if skin.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("name is required")))
	}

	result, err := handler.faceProblemsUsecase.CreateProblem(skin, *file, c, create_by_token, "skin")
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
	skins, err := handler.faceProblemsUsecase.GetProblems("skin")

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
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	skin, err := handler.faceProblemsUsecase.GetProblem(uint64(id))

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
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	var skin entities.FaceProblem

	if err := c.BodyParser(&skin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	file, _ := c.FormFile("file")

	result, err := handler.faceProblemsUsecase.UpdateFaceProblems(id, skin, file, c)
	if err != nil {
		if err.Error() == "skin not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
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
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrNotFound))
	}

	err = handler.faceProblemsUsecase.DeleteFaceProblem(id)

	if err != nil {
		if err.Error() == "skin not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(fiber.ErrNotFound))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.DeleteResponse(id))
}
