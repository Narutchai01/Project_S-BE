package adapters

import (
	"mime/multipart"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpThreadRepository struct {
	threadUsecase usecases.ThreadUseCase
}

func NewHttpThreadRepository(threadUsecase usecases.ThreadUseCase) *HttpThreadRepository {
	return &HttpThreadRepository{threadUsecase}
}

func validateThread(thread entities.Thread) error {
	if thread.Title == "" || thread.Caption == "" {
		return fiber.ErrBadRequest
	}
	return nil
}

func (repo *HttpThreadRepository) CreateThread(c *fiber.Ctx) error {
	var thread entities.Thread

	if err := c.BodyParser(&thread); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrBadRequest))
	}

	if err := validateThread(thread); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrBadRequest))
	}

	token := c.Get("token")

	form, err := c.MultipartForm()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrBadRequest))
	}

	var files []*multipart.FileHeader
	for _, fh := range form.File {
		files = append(files, fh...)
	}

	result, err := repo.threadUsecase.CreateThread(thread, token, files, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(presentation.ToThreadResponse(result))
}

func (repo *HttpThreadRepository) GetThread(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrBadRequest))
	}

	token := c.Get("token")

	result, err := repo.threadUsecase.GetThread(uint(id), token)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToThreadResponse(result))
}

func (repo *HttpThreadRepository) GetThreads(c *fiber.Ctx) error {
	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(fiber.ErrUnauthorized))

	}

	result, err := repo.threadUsecase.GetThreads(token)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToThreadsResponse(result))
}
