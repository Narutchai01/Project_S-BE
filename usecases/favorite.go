package usecases

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
)

type FavoriteUseCase interface {
	FavoriteCommentThread(thread_id uint, token string) (entities.FavoriteCommentThread, error)
	FavoriteThread(thread_id uint, token string) (entities.FavoriteThread, error)
	FavoriteReviewSkincare(review_id uint, token string) (entities.FavoriteReviewSkincare, error)
	FavoriteCommnetReviewSkincare(comment_id uint, token string) (entities.FavoriteCommentReviewSkincare, error)
}

type favoriteService struct {
	repo     repositories.FavoriteRepository
	userRepo repositories.UserRepository
}

func NewFavoriteUseCase(repo repositories.FavoriteRepository, userRepo repositories.UserRepository) FavoriteUseCase {
	return &favoriteService{repo, userRepo}
}

func (service *favoriteService) FavoriteCommentThread(comment_id uint, token string) (entities.FavoriteCommentThread, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.FavoriteCommentThread{}, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.FavoriteCommentThread{}, err
	}

	favorite, err := service.repo.FindFavoriteCommentThread(comment_id, user.ID)
	if err != nil {
		return service.repo.FavoriteCommentThread(comment_id, user.ID)
	}

	favorite.Status = !favorite.Status
	return service.repo.UpdateFavoriteCommentThread(favorite)

}

func (service *favoriteService) FavoriteThread(thread_id uint, token string) (entities.FavoriteThread, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.FavoriteThread{}, err
	}
	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.FavoriteThread{}, err
	}
	favorite, err := service.repo.FindFavoriteThread(thread_id, user.ID)
	if err != nil {
		return service.repo.FavoriteThread(thread_id, user.ID)
	}

	favorite.Status = !favorite.Status
	return service.repo.UpdateFavoriteThread(favorite)
}

func (service *favoriteService) FavoriteReviewSkincare(review_id uint, token string) (entities.FavoriteReviewSkincare, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.FavoriteReviewSkincare{}, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.FavoriteReviewSkincare{}, err
	}

	favorite, err := service.repo.FindFavoriteReviewSkincare(review_id, user.ID)
	if err != nil {
		return service.repo.FavoriteReviewSkincare(review_id, user.ID)
	}

	favorite.Status = !favorite.Status
	return service.repo.UpdateFavoriteReviewSkincare(favorite)
}

func (service *favoriteService) FavoriteCommnetReviewSkincare(comment_id uint, token string) (entities.FavoriteCommentReviewSkincare, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.FavoriteCommentReviewSkincare{}, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.FavoriteCommentReviewSkincare{}, err
	}

	favorite, err := service.repo.FindFavoriteCommentReviewSkincare(comment_id, user.ID)
	if err != nil {
		return service.repo.FavoriteCommentReviewSkincare(comment_id, user.ID)
	}

	favorite.Status = !favorite.Status
	return service.repo.UpdateFavoriteCommentReviewSkincare(favorite)
}
