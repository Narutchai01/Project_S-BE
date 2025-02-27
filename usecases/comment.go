package usecases

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
)

type CommentUsecase interface {
	CreateCommentThread(comment entities.CommentThread, token string) (entities.CommentThread, error)
	GetCommentsThread(thread_id uint, token string) ([]entities.CommentThread, error)
	CreateCommentReviewSkicnare(comment entities.CommentReviewSkicare, token string) (entities.CommentReviewSkicare, error)
	GetCommentsReviewSkincare(review_id uint, token string) ([]entities.CommentReviewSkicare, error)
}

type commentService struct {
	repo         repositories.CommentRepository
	favoriteRepo repositories.FavoriteRepository
	userRepo     repositories.UserRepository
}

func NewCommentUseCase(repo repositories.CommentRepository, favoriteRepo repositories.FavoriteRepository, userRepo repositories.UserRepository) CommentUsecase {
	return &commentService{repo, favoriteRepo, userRepo}
}

func (service *commentService) CreateCommentThread(comment entities.CommentThread, token string) (entities.CommentThread, error) {

	user_id, err := utils.ExtractToken(token)

	if err != nil {
		return entities.CommentThread{}, err
	}

	comment.UserID = user_id

	return service.repo.CreateCommentThread(comment)
}

func (service *commentService) GetCommentsThread(thread_id uint, token string) ([]entities.CommentThread, error) {

	user_id, err := utils.ExtractToken(token)

	if err != nil {
		return []entities.CommentThread{}, err
	}

	result, err := service.repo.GetCommentsThread(thread_id)

	if err != nil {
		return []entities.CommentThread{}, err
	}

	for i, comment := range result {
		favorite, err := service.favoriteRepo.FindFavoriteCommentThread(comment.ID, user_id)
		if err != nil {
			result[i].Favorite = false
		} else {
			result[i].Favorite = favorite.Status
		}
	}

	return result, nil
}

func (service *commentService) CreateCommentReviewSkicnare(comment entities.CommentReviewSkicare, token string) (entities.CommentReviewSkicare, error) {

	user_id, err := utils.ExtractToken(token)

	if err != nil {
		return entities.CommentReviewSkicare{}, err
	}

	user, err := service.userRepo.GetUser(user_id)

	if err != nil {
		return entities.CommentReviewSkicare{}, err
	}

	comment.UserID = user.ID

	return service.repo.CreateCommentReviewSkicnare(comment)
}

func (service *commentService) GetCommentsReviewSkincare(review_id uint, token string) ([]entities.CommentReviewSkicare, error) {

	user_id, err := utils.ExtractToken(token)

	if err != nil {
		return []entities.CommentReviewSkicare{}, err
	}

	result, err := service.repo.GetCommentsReviewSkincare(review_id)

	if err != nil {
		return []entities.CommentReviewSkicare{}, err
	}

	for i, comment := range result {
		favorite, err := service.favoriteRepo.FindFavoriteCommentReviewSkincare(comment.ID, user_id)
		if err != nil {
			result[i].Favorite = false
		} else {
			result[i].Favorite = favorite.Status
		}

		favorite_count, err := service.favoriteRepo.CountFavoriteCommentReviewSkincare(comment.ID)
		if err != nil {
			result[i].FavoriteCount = 0
		} else {
			result[i].FavoriteCount = int(favorite_count)
		}
	}

	return result, nil
}
