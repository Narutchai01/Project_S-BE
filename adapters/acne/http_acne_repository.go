package adapters

import (
	"strconv"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpAcneHandler struct {
	faceProblemsUsecase usecases.FaceProblemUseCase
}

func NewHttpAcneHandler(faceProblemsUsecase usecases.FaceProblemUseCase) *HttpAcneHandler {
	return &HttpAcneHandler{faceProblemsUsecase}
}

// CreateAcne godoc
//
//	@Summary		Create an acne
//	@Description	Create an acne
//	@Tags			acne
//	@Accept			json
//	@Produce		json
//	@Param			acne	formData	entities.Acne	true	"Acne Object"
//	@Param			file	formData	file			true	"Acne Image"
//	@Param			token	header		string			true	"Token"
//	@Success		201		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		404		{object}	presentation.Responses
//	@Router			/admin/acne/ [post]
func (handler *HttpAcneHandler) CreateAcne(c *fiber.Ctx) error {
	var acne entities.FaceProblem

	if err := c.BodyParser(&acne); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	create_by_token := c.Get("token")

	if acne.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrBadRequest))
	}

	result, err := handler.faceProblemsUsecase.CreateProblem(acne, *file, c, create_by_token, "acne")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(presentation.ToAcneResponse(result))
}

// GetAcnes godoc
//
//	@Summary		Get acnes
//	@Description	Get acnes
//	@Tags			acne
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	presentation.Responses
//	@Failure		400	{object}	presentation.Responses
//	@Failure		404	{object}	presentation.Responses
//
//	@Router			/acne [get]
func (handler *HttpAcneHandler) GetAcnes(c *fiber.Ctx) error {
	result, err := handler.faceProblemsUsecase.GetProblems("acne")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToAcnesResponse(result))
}

// GetAcne godoc
//
//	@Summary		Get acne
//	@Description	Get acne
//	@Tags			acne
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Acne ID"
//	@Success		200	{object}	presentation.Responses
//	@Failure		400	{object}	presentation.Responses
//	@Failure		404	{object}	presentation.Responses
//	@Router			/acne/{id} [get]
func (handler *HttpAcneHandler) GetAcne(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	result, err := handler.faceProblemsUsecase.GetProblem(uint64(id))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToAcneResponse(result))
}

// UpdateAcne godoc
//
//	@Summary		Update an acne by ID
//	@Description	Update an acne by ID
//	@Tags			acne
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"Acne ID"
//	@Param			acne	formData	entities.Acne	true	"Acne Object"
//	@Param			file	formData	file			false	"Acne Image"
//	@Success		200		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		404		{object}	presentation.Responses
//	@Router			/admin/acne/{id} [put]
func (handler *HttpAcneHandler) UpdateAcne(c *fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	var acne entities.FaceProblem

	if err := c.BodyParser(&acne); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	file, err := c.FormFile("file")
	if err != nil {
		file = nil
	}

	result, err := handler.faceProblemsUsecase.UpdateFaceProblems(intID, acne, file, c)
	if err != nil {
		if err.Error() == "acne not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToAcneResponse(result))
}

// DeleteAcne godoc
//
//	@Summary		Delete an acne by ID
//	@Description	Delete an acne by ID
//	@Tags			acne
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Acne ID"
//	@Success		204	{object}	presentation.Responses
//	@Failure		400	{object}	presentation.Responses
//	@Failure		404	{object}	presentation.Responses
//	@Router			/admin/acne/{id} [delete]
func (handler *HttpAcneHandler) DeleteAcne(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	err = handler.faceProblemsUsecase.DeleteFaceProblem(id)

	if err != nil {
		if err.Error() == "acne not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.DeleteResponse(id))
}
