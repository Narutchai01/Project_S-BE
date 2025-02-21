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

func (handler *HtppCommentHandler) CreateComment(c *fiber.Ctx) error {

	token := c.Get("token")
	if token == "" {
		return c.Status(400).JSON(presentation.ErrorResponse(errors.New("token is required")))
	}

	var comment entities.Comment
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(400).JSON(presentation.ErrorResponse(err))
	}

	result, err := handler.comment.CreateComment(comment, token)
	if err != nil {
		return c.Status(400).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(200).JSON(result)

}

func (handler *HtppCommentHandler) GetComment(c *fiber.Ctx) error {
	id := c.Params("thread_id")

	thread_id, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(presentation.ErrorResponse(errors.New("failed to get comment")))
	}

	token := c.Get("token")

	if token == "" {
		return c.Status(401).JSON(presentation.ErrorResponse(errors.New("token is required")))
	}

	result, err := handler.comment.GetComments(uint(thread_id), token)
	if err != nil {
		return c.Status(400).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(200).JSON(result)

}
