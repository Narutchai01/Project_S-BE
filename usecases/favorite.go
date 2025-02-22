package usecases

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
)

type FavoriteUseCase interface {
	FavoriteComment(thread_id uint, token string) (entities.FavoriteComment, error)
	FavoriteThread(thread_id uint, token string) (entities.FavoriteThread, error)
}

type favoriteService struct {
	repo repositories.FavoriteRepository
}

func NewFavoriteUseCase(repo repositories.FavoriteRepository) FavoriteUseCase {
	return &favoriteService{repo}
}

func (service *favoriteService) FavoriteComment(comment_id uint, token string) (entities.FavoriteComment, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.FavoriteComment{}, err
	}

	favorite, err := service.repo.FindFavoriteComment(comment_id, user_id)
	if err != nil {
		return service.repo.FavoriteComment(comment_id, user_id)
	}

	favorite.Status = !favorite.Status
	return service.repo.UpdateFavoriteComment(favorite)

}

func (service *favoriteService) FavoriteThread(thread_id uint, token string) (entities.FavoriteThread, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.FavoriteThread{}, err
	}

	favorite, err := service.repo.FindFavoriteThread(thread_id, user_id)
	if err != nil {
		return service.repo.FavoriteThread(thread_id, user_id)
	}

	favorite.Status = !favorite.Status
	return service.repo.UpdateFavoriteThread(favorite)
}
