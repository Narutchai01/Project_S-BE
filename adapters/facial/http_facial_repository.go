package adapters

import (
	"strconv"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpFacialHandler struct {
	facialUsecase usecases.FacialUsecases
}

func NewHttpFacialHandler(facialUcase usecases.FacialUsecases) *HttpFacialHandler {
	return &HttpFacialHandler{facialUcase}
}

func (handler *HttpFacialHandler) CreateFacial(c *fiber.Ctx) error {
	var facial entities.Facial

	if err := c.BodyParser(&facial); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(facial)
	}

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.ErrBadGateway.Code).JSON(facial)
	}

	create_by_token := c.Get("token")

	result, err := handler.facialUsecase.CreateFacial(facial, *file, c, create_by_token)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(facial)
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func (handler *HttpFacialHandler) GetFacials(c *fiber.Ctx) error {
	facial, err := handler.facialUsecase.GetFacials()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(facial)
}

func (handler *HttpFacialHandler) GetFacial(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	facial, err := handler.facialUsecase.GetFacial(intID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(facial)
}
