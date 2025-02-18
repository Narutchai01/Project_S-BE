package usecases

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
)

type ThreadUseCase interface {
	CreateThread(threadDetails entities.ThreadRequest, token string) (entities.Thread, error)
}

type threadService struct {
	repo repositories.ThreadRepository
}

func NewThreadUseCase(repo repositories.ThreadRepository) ThreadUseCase {
	return &threadService{repo}
}

func (service *threadService) CreateThread(threadDetails entities.ThreadRequest, token string) (entities.Thread, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.Thread{}, err
	}

	result, err := service.repo.CreateThread(threadDetails.ThreadDetail, user_id)

	if err != nil {
		return entities.Thread{}, err
	}

	return result, nil
}
