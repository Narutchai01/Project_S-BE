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

func (handler *HttpThreadHandler) GetThreads(c *fiber.Ctx) error {
	result, err := handler.threadUsecase.GetThreads()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToThreadListResponse(result))
}

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

// DeleteThread is a function to delete thread by ID
//
// @Summary Delete thread by ID
// @Tags Thread
// @Description Delete thread by ID
// @ID delete-thread
// @Accept json
// @Produce json
// @Param id path int true "Thread ID"
// @Success 200 {object} DeleteResponse
// @Failure 400 {object} ErrorResponse
// @Router /thread/{id} [delete]
func (handler *HttpThreadHandler) DeleteThread(c *fiber.Ctx) error {
	id := c.Params("id")
	threadID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("invalid thread ID")))
	}

	err = handler.threadUsecase.DeleteThread(uint(threadID))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.DeleteResponse(threadID))
}
