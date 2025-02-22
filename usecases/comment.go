package usecases

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
)

type CommentUsecase interface {
	CreateComment(comment entities.Comment, token string) (entities.Comment, error)
	GetComments(thread_id uint, token string) ([]entities.Comment, error)
}

type commentService struct {
	repo         repositories.CommentRepository
	favoriteRepo repositories.FavoriteRepository
}

func NewCommentUseCase(repo repositories.CommentRepository, favoriteRepo repositories.FavoriteRepository) CommentUsecase {
	return &commentService{repo, favoriteRepo}
}

func (service *commentService) CreateComment(comment entities.Comment, token string) (entities.Comment, error) {

	user_id, err := utils.ExtractToken(token)

	if err != nil {
		return entities.Comment{}, err
	}

	comment.UserID = user_id

	return service.repo.CreateComment(comment)
}

func (service *commentService) GetComments(thread_id uint, token string) ([]entities.Comment, error) {

	user_id, err := utils.ExtractToken(token)

	if err != nil {
		return []entities.Comment{}, err
	}

	result, err := service.repo.GetComments(thread_id)

	if err != nil {
		return []entities.Comment{}, err
	}

	for i, comment := range result {
		favorite, err := service.favoriteRepo.FindFavoriteComment(comment.ID, user_id)
		if err != nil {
			result[i].Favorite = false
		} else {
			result[i].Favorite = favorite.Status
		}
	}

	return result, nil
}
