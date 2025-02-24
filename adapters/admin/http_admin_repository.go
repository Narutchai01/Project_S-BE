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

// create func for validate admin
func validateAdmin(admin entities.Admin) error {
	if admin.Email == "" || admin.Password == "" || admin.FullName == "" {
		return fiber.ErrBadRequest
	}
	return nil
}

// CreateAdmin godoc
//
//	@Summary		Create an admin
//	@Description	Create an admin
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			admin	formData	entities.Admin	true	"Admin Object"
//	@Param			file	formData	file			true	"Admin Image"
//	@Success		201		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		404		{object}	presentation.Responses
//	@Router			/admin/manage [post]
func (handler *HttpAdminHandler) CreateAdmin(c *fiber.Ctx) error {
	var admin entities.Admin

	if err := c.BodyParser(&admin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	if err := validateAdmin(admin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	result, err := handler.adminUcase.CreateAdmin(admin, *file, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(presentation.ToAdminResponse(result))
}

// GetAdmin godoc
//
//	@Summary		Get admins
//	@Description	Get admins
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	presentation.Responses
//	@Failure		400	{object}	presentation.Responses
//	@Failure		404	{object}	presentation.Responses
//	@Router			/admin/manage [get]
func (handler *HttpAdminHandler) GetAdmins(c *fiber.Ctx) error {
	admins, err := handler.adminUcase.GetAdmins()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToAdminsResponse(admins))
}

// GetAdmin godoc
//
//	@Summary		Get an admin by ID
//	@Description	Get an admin by ID
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Admin ID"
//	@Success		200	{object}	presentation.Responses
//	@Failure		400	{object}	presentation.Responses
//	@Failure		404	{object}	presentation.Responses
//	@Router			/admin/manage/{id} [get]
func (handler *HttpAdminHandler) GetAdmin(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	admin, err := handler.adminUcase.GetAdmin(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToAdminResponse(admin))
}

// UpdateAdmin godoc
//
//	@Summary		Update an admin by ID
//	@Description	Update an admin by ID
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			admin	formData	entities.Admin	true	"Admin Object"
//	@Param			file	formData	file			false	"Admin Image"
//	@Param			token	header		string			true	"Admin Token"
//	@Success		200		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		404		{object}	presentation.Responses
//	@Router			/admin/manage/ [put]
func (handler *HttpAdminHandler) UpdateAdmin(c *fiber.Ctx) error {
	var admin entities.Admin

	if err := c.BodyParser(&admin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	adminToken := c.Get("token")

	file, err := c.FormFile("file")
	if err != nil {
		file = nil
	}

	result, err := handler.adminUcase.UpdateAdmin(adminToken, admin, file, c)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToAdminResponse(result))
}

// DeleteAdmin godoc
//
//	@Summary		Delete an admin by ID
//	@Description	Delete an admin by ID
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Admin ID"
//	@Success		204	{object}	presentation.Responses
//	@Failure		400	{object}	presentation.Responses
//	@Failure		404	{object}	presentation.Responses
//	@Router			/admin/manage/{id} [delete]
func (handler *HttpAdminHandler) DeleteAdmin(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	_, err = handler.adminUcase.DeleteAdmin(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.DeleteResponse(id))
}

// LogIn godoc

// @Summary		Log in
// @Description	Log in
// @Tags			admin
// @Accept			json
// @Produce		json
// @Param			admin	body		object{email=string,password=string}	true	"Admin Object"
// @Success		200		{object}	presentation.Responses
// @Failure		400		{object}	presentation.Responses
// @Failure		404		{object}	presentation.Responses
// @Router			/admin/login [post]
func (handler *HttpAdminHandler) LogIn(c *fiber.Ctx) error {
	var admin entities.Admin

	if err := c.BodyParser(&admin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result, err := handler.adminUcase.LogIn(admin.Email, admin.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.TokenResponse(result))
}

// GetAdminByToken godoc
//
//	@Summary		Get an admin by token
//	@Description	Get an admin by token
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			token	header		string	true	"Admin Token"
//	@Success		200		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		404		{object}	presentation.Responses
//	@Router			/admin/profile [get]
func (handler *HttpAdminHandler) GetAdminByToken(c *fiber.Ctx) error {
	token := c.Get("token")

	result, err := handler.adminUcase.GetAdminByToken(token)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToAdminResponse(result))
}
