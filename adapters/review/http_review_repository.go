package adapters

import (
	"encoding/json"
	"errors"
	"mime/multipart"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HtttpReviewRepository struct {
	communityUseccase usecases.CommunityUseCase
}

func NewHttpReviewRepository(communityUsecase usecases.CommunityUseCase) *HtttpReviewRepository {
	return &HtttpReviewRepository{communityUsecase}
}

func (repo *HtttpReviewRepository) CreateReviewSkincare(c *fiber.Ctx) error {
	var review entities.Community
	var skincare_id []int

	review.Title = c.FormValue("title")
	review.Caption = c.FormValue("content")

	if review.Title == "" || review.Caption == "" {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrBadRequest))
	}

	if err := json.Unmarshal([]byte(c.FormValue("skincare_id")), &skincare_id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	if len(skincare_id) < 1 || len(skincare_id) > 10 {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrBadRequest))
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

	files := []*multipart.FileHeader{file}

	result, err := repo.communityUseccase.CreateCommunityThread(review, token, files, c, "Review")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(presentation.ToReviewResponse(result))

}

func (repo *HtttpReviewRepository) GetReviewSkincare(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(errors.New("invalid ID")))
	}

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(fiber.ErrUnauthorized))
	}

	result, err := repo.communityUseccase.GetCommunity(uint(id), "Review", token)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(errors.New("review not found")))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToReviewResponse(result))
}

func (repo *HtttpReviewRepository) GetReviewSkincares(c *fiber.Ctx) error {

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(presentation.ErrorResponse(fiber.ErrUnauthorized))
	}

	results, err := repo.communityUseccase.GetCommunities("Review", token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToReviewsResponse(results))
}
