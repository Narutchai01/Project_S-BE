package adapters

import (
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

func (handler *HttpSkinHandler) CreateSkin(c *fiber.Ctx) error {
	var skin entities.Skin

	if err := c.BodyParser(&skin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.SkinErrorResponse(err))
	}

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.ErrBadGateway.Code).JSON(presentation.SkinErrorResponse(err))
	}

	create_by_token := c.Get("token")

	result, err := handler.skinUsecase.CreateSkin(skin, *file, c, create_by_token)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.SkinErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(presentation.ToSkinResponse(result))
}

func (handler *HttpSkinHandler) GetSkins(c *fiber.Ctx) error {
	skins, err := handler.skinUsecase.GetSkins()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.SkinErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToSkinsResponse(skins))
}

func (handler *HttpSkinHandler) GetSkin(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	skin, err := handler.skinUsecase.GetSkin(intID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.SkinErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToSkinResponse(skin))
}

func (handler *HttpSkinHandler) UpdateSkin(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var skin entities.Skin

	if err := c.BodyParser(&skin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.SkinErrorResponse(err))
	}

	result, err := handler.skinUsecase.UpdateSkin(intID, skin)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.SkinErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToSkinResponse(result))
}

func (handler *HttpSkinHandler) DeleteSkin(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	err = handler.skinUsecase.DeleteSkin(intID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.SkinErrorResponse(err))
	}

	return c.Status(fiber.StatusNoContent).JSON(nil)
}
