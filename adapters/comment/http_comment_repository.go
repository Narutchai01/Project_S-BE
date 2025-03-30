package adapters

import (
	"errors"
	"strconv"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpCommentHandler struct {
	comment usecases.CommentUsecase
}

func NewHttpCommentHandler(commentUsecase usecases.CommentUsecase) *HtppCommentHandler {
	return &HtppCommentHandler{commentUsecase}
}

func (handler *HtppCommentHandler) CreateCommentThread(c *fiber.Ctx) error {

	token := c.Get("token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(errors.New("token is required")))
	}

	var comment entities.CommentThread
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	if comment.Text == "" || comment.ThreadID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("text and thread_id is required")))
	}

	newComment := entities.Comment{
		CommunityID: comment.ThreadID,
		Content:     comment.Text,
	}

	result, err := handler.comment.CreateComment(newComment, token, "thread")
	if err != nil {
		if err.Error() == "thread not found" || err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(presentation.ToCommentThread(result))

}

func (handler *HtppCommentHandler) GetCommentsThread(c *fiber.Ctx) error {
	id := c.Params("thread_id")

	thread_id, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("failed to get comment")))
	}

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(errors.New("token is required")))
	}

	result, err := handler.comment.GetComments(uint(thread_id), "thread", token)

	if err != nil {
		if err.Error() == "thread not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToCommentsThread(result))

}

func (handler *HtppCommentHandler) CreateCommentReviewSkicnare(c *fiber.Ctx) error {

	token := c.Get("token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(fiber.ErrUnauthorized))
	}

	var comment entities.CommentReviewSkicare
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	if comment.Content == "" || comment.ReviewSkincareID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("content and review_id is required")))
	}

	newComment := entities.Comment{
		CommunityID: comment.ReviewSkincareID,
		Content:     comment.Content,
	}

	result, err := handler.comment.CreateComment(newComment, token, "review")
	if err != nil {
		if err.Error() == "review not found" || err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(presentation.ToCommentReviewSkincare(result))

}

func (handler *HtppCommentHandler) HandleGetCommentReviewSkincare(c *fiber.Ctx) error {

	id := c.Params("review_id")

	review_id, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("failed to get comment review skincare")))
	}

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(errors.New("token is required")))
	}

	result, err := handler.comment.GetComments(uint(review_id), "review", token)
	if err != nil {
		if err.Error() == "review not found" {
			return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(err))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToCommentsReviewSkincare(result))
}
