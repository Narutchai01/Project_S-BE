package adapters

import (
	"errors"
	"strconv"

	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpBookmarkHandler struct {
	bookMark usecases.BookmarkUseCase
}

func NewHttpBookmarkHandler(bookmarkUsecase usecases.BookmarkUseCase) *HttpBookmarkHandler {
	return &HttpBookmarkHandler{bookmarkUsecase}
}

// BookMarkThread godoc
// @Summary Bookmark a thread
// @Description Bookmark a thread
// @Tags bookmark
// @Accept json
// @Produce json
// @Param id path int true "Thread ID"
// @Param token header string true "Token"
// @Success 200 {object} presentation.Responses
// @Failure 400 {object} presentation.Responses
// @Failure 401 {object} presentation.Responses
// @Router /bookmark/{id} [post]
func (handler *HttpBookmarkHandler) BookMarkThread(c *fiber.Ctx) error {
	id := c.Params("id")
	threadID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("failed to bookmark thread")))
	}

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(errors.New("token is required")))
	}

	result, err := handler.bookMark.BookmarkThread(uint(threadID), token)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("invalid thread ID")))

	}

	return c.Status(fiber.StatusOK).JSON(result)
}
