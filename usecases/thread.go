package usecases

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
)

type ThreadUseCase interface {
	CreateThread(threadDetails entities.ThreadRequest, token string) (entities.Thread, error)
	GetThreads() ([]entities.Thread, error)
	GetThread(id uint) (entities.Thread, error)
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

	thread, err := service.repo.CreateThread(user_id)
	if err != nil {
		return entities.Thread{}, err
	}

	for _, threadDetail := range threadDetails.ThreadDetail {
		threadDetail.ThreadID = thread.ID
		_, err := service.repo.CreateThreadDetail(threadDetail)
		if err != nil {
			return entities.Thread{}, err
		}
	}

	result, err := service.GetThread(thread.ID)
	if err != nil {
		return entities.Thread{}, err
	}

	thread_details, err := service.repo.GetThreadDetails(result.ID)
	if err != nil {
		return entities.Thread{}, err
	}

	result.Threads = thread_details

	return result, nil
}

func (service *threadService) GetThreads() ([]entities.Thread, error) {
	result, err := service.repo.GetThreads()
	if err != nil {
		return []entities.Thread{}, err
	}

	for i, thread := range result {
		thread_details, err := service.repo.GetThreadDetails(thread.ID)
		if err != nil {
			return []entities.Thread{}, err
		}

		result[i].Threads = thread_details
	}

	return result, nil
}

func (service *threadService) GetThread(id uint) (entities.Thread, error) {
	result, err := service.repo.GetThread(id)
	if err != nil {
		return entities.Thread{}, err
	}

	thread_details, err := service.repo.GetThreadDetails(result.ID)
	if err != nil {
		return entities.Thread{}, err
	}

	result.Threads = thread_details

	return result, nil
}
