package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpAdminHandler struct {
	adminUcase usecases.AdminUsecases
}

func NewHttpAdminHandler(adminUcase usecases.AdminUsecases) *HttpAdminHandler {
	return &HttpAdminHandler{adminUcase}
}

func (handler *HttpAdminHandler) CreateAdmin(c *fiber.Ctx) error {
	var admin entities.Admin

	if err := c.BodyParser(&admin); err != nil {
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

	result, err := handler.adminUcase.CreateAdmin(admin, *file, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.AdminErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func (handler *HttpAdminHandler) GetAdmins(c *fiber.Ctx) error {
	admins, err := handler.adminUcase.GetAdmins()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.AdminErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToAdminsResponse(admins))
}

func (handler *HttpAdminHandler) GetAdmin(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.AdminErrorResponse(err))
	}

	admin, err := handler.adminUcase.GetAdmin(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presentation.AdminErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToAdminResponse(admin))
}

func (handler *HttpAdminHandler) UpdateAdmin(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.AdminErrorResponse(err))
	}

	var admin entities.Admin

	if err := c.BodyParser(&admin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.AdminErrorResponse(err))
	}

	result, err := handler.adminUcase.UpdateAdmin(id, admin)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.AdminErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToAdminResponse(result))
}

func (handler *HttpAdminHandler) DeleteAdmin(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	_, err = handler.adminUcase.DeleteAdmin(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.AdminErrorResponse(err))
	}

	return c.Status(fiber.StatusNoContent).JSON(presentation.DeleteAdminResponse(id))
}

func (handler *HttpAdminHandler) LogIn(c *fiber.Ctx) error {
	var admin entities.Admin

	if err := c.BodyParser(&admin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result, err := handler.adminUcase.LogIn(admin.Email, admin.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.AdminErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToAdminResponse(result))
}
