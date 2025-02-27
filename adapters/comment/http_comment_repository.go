package adapters

import (
	"errors"
	"strconv"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HtppCommentHandler struct {
	comment usecases.CommentUsecase
}

func NewHttpCommentHandler(commentUsecase usecases.CommentUsecase) *HtppCommentHandler {
	return &HtppCommentHandler{commentUsecase}
}

// create swwager for CreateComment
// CreateComment godoc
// @Summary Create a comment
// @Description Create a comment
// @Tags comment
// @Accept json
// @Produce json
// @Param token header string true "Token"
// @Param comment body object{thread_id=uint,text=string} true "Comment"
// @Success 200 {object} presentation.Responses
// @Failure 400 {object} presentation.Responses
// @Router /comment [post]
func (handler *HtppCommentHandler) CreateCommentThread(c *fiber.Ctx) error {

	token := c.Get("token")
	if token == "" {
		return c.Status(400).JSON(presentation.ErrorResponse(errors.New("token is required")))
	}

	var comment entities.CommentThread
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(400).JSON(presentation.ErrorResponse(err))
	}

	result, err := handler.comment.CreateCommentThread(comment, token)
	if err != nil {
		return c.Status(400).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(presentation.ToCommentThread(result))

}

// create swwager for GetComment
// GetComment godoc
// @Summary Get a comment
// @Description Get a comment
// @Tags comment
// @Accept json
// @Produce json
// @Param thread_id path int true "Thread ID"
// @Param token header string true "Token"
// @Success 200 {object} presentation.Responses
// @Failure 400 {object} presentation.Responses
// @Failure 401 {object} presentation.Responses
// @Router /comment/{thread_id} [get]
func (handler *HtppCommentHandler) GetCommentsThread(c *fiber.Ctx) error {
	id := c.Params("thread_id")

	thread_id, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(presentation.ErrorResponse(errors.New("failed to get comment")))
	}

	token := c.Get("token")

	if token == "" {
		return c.Status(401).JSON(presentation.ErrorResponse(errors.New("token is required")))
	}

	result, err := handler.comment.GetCommentsThread(uint(thread_id), token)
	if err != nil {
		return c.Status(400).JSON(presentation.ErrorResponse(err))
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

	result, err := handler.comment.CreateCommentReviewSkicnare(comment, token)
	if err != nil {
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

	result, err := handler.comment.GetCommentsReviewSkincare(uint(review_id), token)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToCommentsReviewSkincare(result))

}
