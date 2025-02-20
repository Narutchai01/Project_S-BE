package adapters

import (
	"errors"
	"strconv"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpThreadHandler struct {
	threadUsecase usecases.ThreadUseCase
}

func NewHttpThreadHandler(threadUcase usecases.ThreadUseCase) *HttpThreadHandler {
	return &HttpThreadHandler{threadUcase}
}

// Create Thread godoc
//
// @Summary		Create a thread
// @Description	Create a thread
// @Tags			thread
// @Accept			json
// @Produce		json
// @Param			thread	body	entities.ThreadRequest	true	"Thread Object"
// @Param			token	header	string	true	"Token"
// @Success		201		{object}	presentation.Responses
// @Failure		400		{object}	presentation.Responses
// @Failure		404		{object}	presentation.Responses
// @Router			/thread/ [post]
func (handler *HttpThreadHandler) CreateThread(c *fiber.Ctx) error {
	var thread entities.ThreadRequest

	if err := c.BodyParser(&thread); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("token is required")))
	}

	if len(thread.ThreadDetail) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("ThreadDetail is required")))
	}

	result, err := handler.threadUsecase.CreateThread(thread, token)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(presentation.ToThreadResponse(result))
}

// Get Threads godoc
//
// @Summary		Get all threads
// @Description	Get all threads
// @Tags			thread
// @Accept			json
// @Produce		json
// @Param			token	header	string	true	"Token"
// @Success		200		{object}	presentation.Responses
// @Failure		400		{object}	presentation.Responses
// @Router			/thread/ [get]
func (handler *HttpThreadHandler) GetThreads(c *fiber.Ctx) error {
	result, err := handler.threadUsecase.GetThreads()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToThreadListResponse(result))
}

// Get Thread godoc
//
// @Summary		Get a thread
// @Description	Get a thread
// @Tags			thread
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"Thread ID"
// @Param			token	header	string	true	"Token"
// @Success		200		{object}	presentation.Responses
// @Failure		400		{object}	presentation.Responses
// @Router			/thread/{id} [get]
func (handler *HttpThreadHandler) GetThread(c *fiber.Ctx) error {
	id := c.Params("id")
	threadID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("invalid thread ID")))
	}

	result, err := handler.threadUsecase.GetThread(uint(threadID))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToThreadResponse(result))
}
