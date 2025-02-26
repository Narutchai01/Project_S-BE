package usecases

import (
	"mime/multipart"
	"os"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ReviewThreadUseCase interface {
	CreateReviewThread(reviewThread entities.ReviewSkincare, token string, file multipart.FileHeader, c *fiber.Ctx) (entities.ReviewSkincare, error)
}

type reviewThreadService struct {
	reviewRepo repositories.ReviewRepository
	userRepo   repositories.UserRepository
}

func NewReviewThreadUseCase(reviewRepo repositories.ReviewRepository, userRepo repositories.UserRepository) ReviewThreadUseCase {
	return &reviewThreadService{reviewRepo, userRepo}
}

func (service *reviewThreadService) CreateReviewThread(review entities.ReviewSkincare, token string, file multipart.FileHeader, c *fiber.Ctx) (entities.ReviewSkincare, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.ReviewSkincare{}, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.ReviewSkincare{}, err
	}

	review.UserID = user.ID

	fileName := uuid.New().String() + ".jpg"

	if err := utils.CheckDirectoryExist(); err != nil {
		return entities.ReviewSkincare{}, err
	}

	if err := c.SaveFile(&file, "./uploads/"+fileName); err != nil {
		return entities.ReviewSkincare{}, err
	}

	imageUrl, err := utils.UploadImage(fileName, "/review")
	if err != nil {
		return entities.ReviewSkincare{}, err
	}

	err = os.Remove("./uploads/" + fileName)
	if err != nil {
		return entities.ReviewSkincare{}, err
	}

	review.Image = imageUrl

	return service.reviewRepo.CreateReviewSkincare(review)
}
