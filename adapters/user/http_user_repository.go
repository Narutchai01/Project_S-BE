package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpUserHandler struct {
	userUcase usecases.UserUsecases
}

func NewHttpUserHandler(userUcase usecases.UserUsecases) *HttpUserHandler {
	return &HttpUserHandler{userUcase}
}

func (handler *HttpUserHandler) Register(c *fiber.Ctx) error {
	var user entities.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result, err := handler.userUcase.Register(user, *file, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.UserErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}