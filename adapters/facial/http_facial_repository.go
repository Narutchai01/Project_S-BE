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

// CreateFacial godoc
//	@Summary		Create facial
//	@Description	Create facial
//	@Tags			facial
//	@Accept			json
//	@Produce		json
//	@Param			file	formData	file			true	"Facial image"
//	@Param			facial	formData	entities.Facial	true	"Facial information"
//	@Param			token	header		string			true	"Token"
//	@Router			/admin/facial [post]
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

// GetFacials godoc
//	@Summary		Get all facials
//	@Description	Get all facials
//	@Tags			facial
//	@Accept			json
//	@Produce		json
//	@Router			/facial [get]
func (handler *HttpFacialHandler) GetFacials(c *fiber.Ctx) error {
	facial, err := handler.facialUsecase.GetFacials()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(facial)
}

// GetFacial godoc
//	@Summary		Get facial by ID
//	@Description	Get facial by ID
//	@Tags			facial
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"Facial ID"
//	@Router			/facial/{id} [get]
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

// UpdateFacial godoc
//	@Summary		Update facial by ID
//	@Description	Update facial by ID
//	@Tags			facial
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int				true	"Facial ID"
//	@Param			facial	body	entities.Facial	true	"Facial information"
//	@Router			/admin/facial/{id} [put]
func (handler *HttpFacialHandler) UpdateFacial(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var facial entities.Facial

	if err := c.BodyParser(&facial); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(facial)
	}

	result, err := handler.facialUsecase.UpdateFacial(intID, facial)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// DeleteFacial godoc
//	@Summary		Delete facial by ID
//	@Description	Delete facial by ID
//	@Tags			facial
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"Facial ID"
//	@Router			/admin/facial/{id} [delete]
func (handler *HttpFacialHandler) DeleteFacial(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	err = handler.facialUsecase.DeleteFacial(intID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusNoContent).JSON(nil)
}
