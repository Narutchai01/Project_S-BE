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
