package usecases

import (
	"errors"

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
	repo        repositories.FavoriteRepository
	userRepo    repositories.UserRepository
	threadRepo  repositories.ThreadRepository
	reviewRepo  repositories.ReviewRepository
	commentRepo repositories.CommentRepository
}

func NewFavoriteUseCase(repo repositories.FavoriteRepository, userRepo repositories.UserRepository, threadRepo repositories.ThreadRepository, reviewRepo repositories.ReviewRepository, commemntRepo repositories.CommentRepository) FavoriteUseCase {
	return &favoriteService{repo, userRepo, threadRepo, reviewRepo, commemntRepo}
}

func (service *favoriteService) FavoriteCommentThread(comment_id uint, token string) (entities.FavoriteCommentThread, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.FavoriteCommentThread{}, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.FavoriteCommentThread{}, errors.New("user not found")
	}

	comment, err := service.commentRepo.GetCommentThread(comment_id)
	if err != nil {
		return entities.FavoriteCommentThread{}, errors.New("comment not found")
	}

	favorite, err := service.repo.FindFavoriteCommentThread(comment.ID, user.ID)
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
		return entities.FavoriteThread{}, errors.New("user not found")
	}

	_, err = service.threadRepo.GetThread(thread_id)
	if err != nil {
		return entities.FavoriteThread{}, errors.New("thread not found")
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
		return entities.FavoriteReviewSkincare{}, errors.New("user not found")
	}

	_, err = service.reviewRepo.GetReviewSkincare(review_id)
	if err != nil {
		return entities.FavoriteReviewSkincare{}, errors.New("review not found")
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
		return entities.FavoriteCommentReviewSkincare{}, errors.New("user not found")
	}

	comment, err := service.commentRepo.GetCommentThread(comment_id)
	if err != nil {
		return entities.FavoriteCommentReviewSkincare{}, errors.New("comment not found")
	}

	favorite, err := service.repo.FindFavoriteCommentReviewSkincare(comment.ID, user.ID)
	if err != nil {
		return service.repo.FavoriteCommentReviewSkincare(comment_id, user.ID)
	}

	favorite.Status = !favorite.Status
	return service.repo.UpdateFavoriteCommentReviewSkincare(favorite)
}
