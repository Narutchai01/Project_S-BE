package adapters

import (
	"encoding/json"
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
// @Param			token	header	string	true	"Token"
// @Param			title	formData 	string	true	"Thread Title"
// @Param			thread_details	formData 	string	true	"Thread Details"
// @Param			file	formData 	file	true	"Thread Image"
// @Success		201		{object}	presentation.Responses
// @Failure		400		{object}	presentation.Responses
// @Failure		404		{object}	presentation.Responses
// @Router			/thread/ [post]
func (handler *HttpThreadHandler) CreateThread(c *fiber.Ctx) error {

	var threadDetails []entities.ThreadDetail
	if err := json.Unmarshal([]byte(c.FormValue("thread_details")), &threadDetails); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("invalid thread details format")))
	}
	title := c.FormValue("title")

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(fiber.ErrUnauthorized))
	}

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	if len(threadDetails) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("ThreadDetail is required")))
	}

	for _, detail := range threadDetails {
		if detail.SkincareID == 0 || detail.Caption == "" {
			return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("SkincareID and Caption is required")))

		}
	}

	result, err := handler.threadUsecase.CreateThread(threadDetails, title, token, *file, c)

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

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(fiber.ErrUnauthorized))
	}

	result, err := handler.threadUsecase.GetThreads(token)

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

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrUnauthorized))
	}

	result, err := handler.threadUsecase.GetThread(uint(threadID), token)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToThreadResponse(result))
}

// Delete Thread godoc
//
// @Summary		Delete a thread
// @Description	Delete a thread
// @Tags			thread
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"Thread ID"
// @Param			token	header	string	true	"Token"
// @Success		200		{object}	presentation.Responses
// @Failure		400		{object}	presentation.Responses
// @Router			/thread/{id} [delete]
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

// Update Thread godoc
//
// @Summary		Update a thread
// @Description	Update a thread
// @Tags			thread
// @Accept			json
// @Produce		json
// @Param			token	header	string	true	"Token"
// @Param			id	path	int	true	"Thread ID"
// @Param			title	formData 	string	false	"Thread Title"
// @Param			thread_details	formData 	string	false	"Thread Details"
// @Param			file	formData 	file	false	"Thread Image"
// @Success		201		{object}	presentation.Responses
// @Failure		400		{object}	presentation.Responses
// @Failure		404		{object}	presentation.Responses
// @Router			/thread/{id} [put]
func (handler *HttpThreadHandler) UpdateThread(c *fiber.Ctx) error {
	id := c.Params("id")
	thread_id, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("invalid thread ID")))
	}

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(fiber.ErrUnauthorized))
	}

	title := c.FormValue("title")

	file, err := c.FormFile("file")

	if err != nil {
		file = nil
	}

	var threadDetails []entities.ThreadDetail
	_ = json.Unmarshal([]byte(c.FormValue("thread_details")), &threadDetails)

	result, err := handler.threadUsecase.UpdateThread(uint(thread_id), token, title, threadDetails, file, c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("invalid thread details format")))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToThreadResponse(result))
}
