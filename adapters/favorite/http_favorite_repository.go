package adapters

import (
	"errors"
	"strconv"

	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpFavoriteHandler struct {
	FavoriteUsecases usecases.FavoriteUseCase
}

func NewHttpFavoriteHandler(favoriteUcase usecases.FavoriteUseCase) *HttpFavoriteHandler {
	return &HttpFavoriteHandler{favoriteUcase}
}

func (handler *HttpFavoriteHandler) HandleFavoriteComment(c *fiber.Ctx) error {

	id := c.Params("id")

	comment_id, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("failed to favorite comment")))
	}

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(errors.New("token is required")))
	}

	result, err := handler.FavoriteUsecases.FavoriteComment(uint(comment_id), token)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(result)

}
