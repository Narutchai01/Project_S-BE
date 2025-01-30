package usecases

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
)

type ResultUsecases interface {
	CreateResult(result entities.Result) (entities.Result, error)
	GetResults() ([]entities.Result, error)
}
type resultService struct {
	repo repositories.ResultRepository
}

func NewResultUseCase(repo repositories.ResultRepository) ResultUsecases {
	return &resultService{repo}
}

func (service *resultService) CreateResult(result entities.Result) (entities.Result, error) {
	return service.repo.CreateResult(result)
}

func (service *resultService) GetResults() ([]entities.Result, error) {
	return service.repo.GetResults()
}