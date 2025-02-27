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

type ReviewUseCase interface {
	CreateReviewSkincare(reviewThread entities.ReviewSkincare, token string, file multipart.FileHeader, c *fiber.Ctx) (entities.ReviewSkincare, error)
	GetReviewSkincare(review_id uint, token string) (entities.ReviewSkincare, error)
	GetReviewSkincares(token string) ([]entities.ReviewSkincare, error)
}

type reviewService struct {
	reviewRepo   repositories.ReviewRepository
	userRepo     repositories.UserRepository
	skincareRepo repositories.SkincareRepository
	favoriteRepo repositories.FavoriteRepository
	bookmarkRepo repositories.BookmarkRepository
}

// GetReviewSkincares implements ReviewUseCase.

func NewReviewUseCase(reviewRepo repositories.ReviewRepository, userRepo repositories.UserRepository, skincareRepo repositories.SkincareRepository, favoriteRepo repositories.FavoriteRepository, bookmarkRepo repositories.BookmarkRepository) ReviewUseCase {
	return &reviewService{reviewRepo, userRepo, skincareRepo, favoriteRepo, bookmarkRepo}
}

func (service *reviewService) CreateReviewSkincare(review entities.ReviewSkincare, token string, file multipart.FileHeader, c *fiber.Ctx) (entities.ReviewSkincare, error) {

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

	review, err = service.reviewRepo.CreateReviewSkincare(review)

	if err != nil {
		return entities.ReviewSkincare{}, err
	}

	review, err = service.reviewRepo.GetReviewSkincare(review.ID)

	if err != nil {
		return entities.ReviewSkincare{}, err
	}

	skincares, err := service.skincareRepo.GetSkincareByIds(review.SkincareID)
	if err != nil {
		return entities.ReviewSkincare{}, err
	}

	review.Skincare = skincares

	favorite, err := service.favoriteRepo.FindFavoriteReviewSkincare(review.ID, user_id)
	if err != nil {
		review.Favorite = false
	} else {
		review.Favorite = favorite.Status
	}

	count_favorite, err := service.favoriteRepo.CountFavoriteReviewSkincare(review.ID)
	if err != nil {
		review.FavoriteCount = 0
	} else {
		review.FavoriteCount = count_favorite
	}

	if user_id != review.UserID {
		review.Owner = false
	} else {
		review.Owner = true
	}

	bookmark, err := service.bookmarkRepo.FindBookMarkReviewSkincare(review.ID, user_id)
	if err != nil {
		review.Bookmark = false
	} else {
		review.Bookmark = bookmark.Status
	}

	return review, nil
}

func (service *reviewService) GetReviewSkincare(review_id uint, token string) (entities.ReviewSkincare, error) {

	user_id, err := utils.ExtractToken(token)

	if err != nil {
		return entities.ReviewSkincare{}, err
	}

	_, err = service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.ReviewSkincare{}, err
	}

	review, err := service.reviewRepo.GetReviewSkincare(review_id)
	if err != nil {
		return entities.ReviewSkincare{}, err
	}

	skincares, err := service.skincareRepo.GetSkincareByIds(review.SkincareID)
	if err != nil {
		return entities.ReviewSkincare{}, err
	}

	review.Skincare = skincares

	favorite, err := service.favoriteRepo.FindFavoriteReviewSkincare(review.ID, user_id)
	if err != nil {
		review.Favorite = false
	} else {
		review.Favorite = favorite.Status
	}

	count_favorite, err := service.favoriteRepo.CountFavoriteReviewSkincare(review.ID)
	if err != nil {
		review.FavoriteCount = 0
	} else {
		review.FavoriteCount = count_favorite
	}

	if user_id != review.UserID {
		review.Owner = false
	} else {
		review.Owner = true
	}

	bookmark, err := service.bookmarkRepo.FindBookMarkReviewSkincare(review.ID, user_id)
	if err != nil {
		review.Bookmark = false
	} else {
		review.Bookmark = bookmark.Status
	}

	return review, nil
}

func (service *reviewService) GetReviewSkincares(token string) ([]entities.ReviewSkincare, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return []entities.ReviewSkincare{}, err
	}

	_, err = service.userRepo.GetUser(user_id)
	if err != nil {
		return []entities.ReviewSkincare{}, err
	}

	reviews, err := service.reviewRepo.GetReviewSkincares()

	if err != nil {
		return []entities.ReviewSkincare{}, err
	}

	for i, review := range reviews {
		skincares, err := service.skincareRepo.GetSkincareByIds(review.SkincareID)
		if err != nil {
			return []entities.ReviewSkincare{}, err
		}
		reviews[i].Skincare = skincares

		favorite, err := service.favoriteRepo.FindFavoriteReviewSkincare(review.ID, user_id)
		if err != nil {
			reviews[i].Favorite = false
		} else {
			reviews[i].Favorite = favorite.Status
		}

		count_favorite, err := service.favoriteRepo.CountFavoriteReviewSkincare(review.ID)
		if err != nil {
			reviews[i].FavoriteCount = 0
		} else {
			reviews[i].FavoriteCount = count_favorite
		}

		if user_id != review.UserID {
			reviews[i].Owner = false
		} else {
			reviews[i].Owner = true
		}

		bookmark, err := service.bookmarkRepo.FindBookMarkReviewSkincare(review.ID, user_id)
		if err != nil {
			reviews[i].Bookmark = false
		} else {
			reviews[i].Bookmark = bookmark.Status
		}
	}

	return reviews, nil
}
