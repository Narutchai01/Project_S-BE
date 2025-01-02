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

func (handler *HttpSkincareHandler) CreateSkincare(c *fiber.Ctx) error {
	var skincare entities.Skincare

	if err := c.BodyParser(&skincare); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	create_by_token := c.Get("token")

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result, err := handler.skincarenUcase.CreateSkincare(skincare, *file, create_by_token, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.SkincareErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func (handler *HttpSkincareHandler) GetSkincares(c *fiber.Ctx) error {
	skincares, err := handler.skincarenUcase.GetSkincares()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.SkincareErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.SkincaresResponse(skincares))
}

func (handler *HttpSkincareHandler) GetSkincare(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.SkincareErrorResponse(err))
	}

	skincare, err := handler.skincarenUcase.GetSkincare(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presentation.SkincareErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.SkincareResponse(skincare))
}

func (handler *HttpSkincareHandler) UpdateSkincare(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.SkincareErrorResponse(err))
	}

	var skincare entities.Skincare

	if err := c.BodyParser(&skincare); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.SkincareErrorResponse(err))
	}

	result, err := handler.skincarenUcase.UpdateSkincare(id, skincare)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.SkincareErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.SkincareResponse(result))
}

func (handler *HttpSkincareHandler) DeleteSkincare(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	_, err = handler.skincarenUcase.DeleteSkincare(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.SkincareErrorResponse(err))
	}

	return c.Status(fiber.StatusNoContent).JSON(presentation.DeleteSkincareResponse(id))
}