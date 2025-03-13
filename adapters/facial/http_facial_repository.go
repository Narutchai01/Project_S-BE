package adapters

import (
	"errors"
	"strconv"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/presentation"
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
//
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
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	create_by_token := c.Get("token")

	if facial.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("name is required")))
	}

	result, err := handler.facialUsecase.CreateFacial(facial, *file, c, create_by_token)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(presentation.ToFacialResponse(result))
}

// GetFacials godoc
//
//	@Summary		Get all facials
//	@Description	Get all facials
//	@Tags			facial
//	@Accept			json
//	@Produce		json
//	@Router			/facial [get]
func (handler *HttpFacialHandler) GetFacials(c *fiber.Ctx) error {
	facial, err := handler.facialUsecase.GetFacials()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToFacialsResponse(facial))
}

// GetFacial godoc
//
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
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	facial, err := handler.facialUsecase.GetFacial(intID)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToFacialResponse(facial))
}

// UpdateFacial godoc
//
//	@Summary		Update facial by ID
//	@Description	Update facial by ID
//	@Tags			facial
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"Facial ID"
//	@Param			facial	formData	entities.Facial	true	"Facial information"
//	@Param			file	formData	file			false	"Facial image"
//	@Router			/admin/facial/{id} [put]
func (handler *HttpFacialHandler) UpdateFacial(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	var facial entities.Facial

	if err := c.BodyParser(&facial); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	file, _ := c.FormFile("file")

	result, err := handler.facialUsecase.UpdateFacial(intID, facial, file, c)
	if err != nil {
		if err.Error() == "facial not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToFacialResponse(result))
}

// DeleteFacial godoc
//
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
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	err = handler.facialUsecase.DeleteFacial(intID)

	if err != nil {
		if err.Error() == "facial not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.DeleteResponse(intID))
}
