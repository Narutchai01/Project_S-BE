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

func (handler *HttpFavoriteHandler) HandleFavoriteCommentThread(c *fiber.Ctx) error {

	id := c.Params("id")

	comment_id, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("failed to favorite comment")))
	}

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(errors.New("token is required")))
	}

	result, err := handler.FavoriteUsecases.FavoriteCommentThread(uint(comment_id), token)
	if err != nil {
		if err.Error() == "comment not found" || err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
		}
	}

	return c.Status(fiber.StatusOK).JSON(result)

}

func (handler *HttpFavoriteHandler) HandleFavoriteThread(c *fiber.Ctx) error {

	id := c.Params("id")

	thread_id, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("failed to favorite thread")))
	}

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(errors.New("token is required")))
	}

	result, err := handler.FavoriteUsecases.FavoriteThread(uint(thread_id), token)
	if err != nil {
		if err.Error() == "thread not found" || err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
		}
	}

	return c.Status(fiber.StatusOK).JSON(result)

}

func (handler *HttpFavoriteHandler) HandleFavoriteReviewSkincare(c *fiber.Ctx) error {

	id := c.Params("id")

	review_id, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("failed to favorite review skincare")))
	}

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(errors.New("token is required")))
	}

	result, err := handler.FavoriteUsecases.FavoriteReviewSkincare(uint(review_id), token)
	if err != nil {
		if err.Error() == "review not found" || err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
		}
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (handler *HttpFavoriteHandler) HandleFavoriteCommentReviewSkincare(c *fiber.Ctx) error {

	id := c.Params("id")

	comment_id, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("failed to favorite comment review skincare")))
	}

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(errors.New("token is required")))
	}

	result, err := handler.FavoriteUsecases.FavoriteCommnetReviewSkincare(uint(comment_id), token)
	if err != nil {
		if err.Error() == "comment not found" || err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
		}
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
