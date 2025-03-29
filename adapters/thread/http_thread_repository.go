package adapters

import (
	"errors"
	"mime/multipart"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpThreadRepository struct {
	communityUseccase usecases.CommunityUseCase
}

func NewHttpThreadRepository(communityUsecase usecases.CommunityUseCase) *HttpThreadRepository {
	return &HttpThreadRepository{communityUsecase}
}

// CreateThread godoc
//
//	@Summary		Create a thread
//	@Description	Create a thread
//	@Tags			threads
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			token	header	string	true	"Token"
//	@Param			title	formData	string	false	"Title"
//	@Param			caption	formData	string	false	"Caption"
//	@Param			files	formData	file	false	"File"
//	@Success		201		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		500		{object}	presentation.Responses
//	@Router			/thread/ [post]
func (repo *HttpThreadRepository) CreateThread(c *fiber.Ctx) error {
	var thread entities.Community

	const communityType = "Thread"

	if err := c.BodyParser(&thread); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrBadRequest))
	}

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(fiber.ErrUnauthorized))
	}

	form, err := c.MultipartForm()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrBadRequest))
	}

	files := make([]*multipart.FileHeader, 0)
	for _, fh := range form.File {
		files = append(files, fh...)
	}

	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrBadRequest))
	}

	result, err := repo.communityUseccase.CreateCommunityThread(thread, token, files, c, communityType)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(presentation.ToThreadResponse(result))
}

// GetThread godoc
//
//	@Summary		Get a thread
//	@Description	Get a thread
//	@Tags			threads
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"ID"
//	@Param			token	header	string	true	"Token"
//	@Success		200		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		500		{object}	presentation.Responses
//	@Router			/thread/{id} [get]
func (repo *HttpThreadRepository) GetThread(c *fiber.Ctx) error {
	const communityType = "Thread"
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrBadRequest))
	}

	token := c.Get("token")

	result, err := repo.communityUseccase.GetCommunity(uint(id), communityType, token)
	if err != nil {
		if err.Error() == "community not found" || err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(errors.New("thread not found")))
		}
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToThreadResponse(result))
}

// GetThreads godoc
//
//	@Summary		Get threads
//	@Description	Get threads
//	@Tags			threads
//	@Accept			json
//	@Produce		json
//	@Param			token	header	string	true	"Token"
//	@Success		200		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		500		{object}	presentation.Responses
//	@Router			/thread/ [get]
func (repo *HttpThreadRepository) GetThreads(c *fiber.Ctx) error {
	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(fiber.ErrUnauthorized))

	}

	result, err := repo.communityUseccase.GetCommunities("Thread", token)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToThreadsResponse(result))
}
