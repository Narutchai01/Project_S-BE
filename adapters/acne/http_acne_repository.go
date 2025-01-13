package adapters

import (
	"strconv"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpAcneHandler struct {
	acneUsecase usecases.AcneUseCase
}

func NewHttpAcneHandler(acneUcase usecases.AcneUseCase) *HttpAcneHandler {
	return &HttpAcneHandler{acneUcase}
}

func (handler *HttpAcneHandler) CreateAcne(c *fiber.Ctx) error {
	var acne entities.Acne

	if err := c.BodyParser(&acne); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(acne)
	}

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.ErrBadGateway.Code).JSON(acne)
	}

	create_by_token := c.Get("token")

	result, err := handler.acneUsecase.CreateAcne(acne, *file, c, create_by_token)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(acne)
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func (handler *HttpAcneHandler) GetAcnes(c *fiber.Ctx) error {
	result, err := handler.acneUsecase.GetAcnes()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (handler *HttpAcneHandler) GetAcne(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	result, err := handler.acneUsecase.GetAcne(intID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (handler *HttpAcneHandler) DeleteAcne(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	err = handler.acneUsecase.DeleteAcne(intID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusNoContent).JSON(nil)
}
