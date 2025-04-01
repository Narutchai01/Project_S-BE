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
	faceProblemsUsecase usecases.FaceProblemUseCase
}

func NewHttpFacialHandler(faceProblemsUsecase usecases.FaceProblemUseCase) *HttpFacialHandler {
	return &HttpFacialHandler{faceProblemsUsecase}
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
	var facial entities.FaceProblem

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

	result, err := handler.faceProblemsUsecase.CreateProblem(facial, *file, c, create_by_token, "facial")

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
	facial, err := handler.faceProblemsUsecase.GetProblems("facial")

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
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
	}

	facial, err := handler.faceProblemsUsecase.GetProblem(uint64(id))

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

	var facial entities.FaceProblem

	if err := c.BodyParser(&facial); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	file, _ := c.FormFile("file")

	result, err := handler.faceProblemsUsecase.UpdateFaceProblems(intID, facial, file, c)
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
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	err = handler.faceProblemsUsecase.DeleteFaceProblem(id)

	if err != nil {
		if err.Error() == "facial not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.DeleteResponse(id))
}
