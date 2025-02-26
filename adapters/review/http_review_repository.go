package adapters

import (
	"encoding/json"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HtttpReviewRepository struct {
	reviewUsecase usecases.ReviewThreadUseCase
}

func NewHttpReviewRepository(reviewUsecase usecases.ReviewThreadUseCase) *HtttpReviewRepository {
	return &HtttpReviewRepository{reviewUsecase}
}

func (repo *HtttpReviewRepository) CreateReviewSkincare(c *fiber.Ctx) error {
	var review entities.ReviewSkincare
	var skincare_id []int

	review.Title = c.FormValue("title")
	review.Content = c.FormValue("content")

	if err := json.Unmarshal([]byte(c.FormValue("skincare_id")), &skincare_id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(fiber.ErrUnauthorized))
	}

	review.SkincareID = skincare_id

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	result, err := repo.reviewUsecase.CreateReviewThread(review, token, *file, c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(result)

}
