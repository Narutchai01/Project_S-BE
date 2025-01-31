package usecases

import (
	"reflect"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
)

type ResultUsecases interface {
	CreateResult(result entities.Result) (entities.Result, error)
	GetResults() ([]entities.Result, error)
	GetResultById(id int) (entities.Result, error)
	UpdateResultById(id int, result entities.Result) (entities.Result, error)
	DeleteResultById(id int) error
	GetResultsByUserIdFromToken(token string) ([]entities.Result, error)
	GetResultsByUserId(id int) ([]entities.Result, error)
}
type resultService struct {
	repo repositories.ResultRepository
}

func NewResultUsecase(repo repositories.ResultRepository) ResultUsecases {
	return &resultService{repo}
}

func (service *resultService) CreateResult(result entities.Result) (entities.Result, error) {
	return service.repo.CreateResult(result)
}

func (service *resultService) GetResults() ([]entities.Result, error) {
	return service.repo.GetResults()
}

func (service *resultService) GetResultById(id int) (entities.Result, error) {
	return service.repo.GetResultById(id)
}

func (service *resultService) UpdateResultById(id int, result entities.Result) (entities.Result, error) {

	old_result, err := service.repo.GetResultById(id)
	if err != nil {
		return entities.Result{}, err
	}

	old_result.Image = utils.CheckEmptyValueBeforeUpdate(result.Image, old_result.Image)
	if result.UserId != old_result.UserId && result.UserId != 0 {
		old_result.UserId = result.UserId
	}
	if len(result.AcneType) > 0 && !reflect.DeepEqual(result.AcneType, old_result.AcneType) {
		old_result.AcneType = result.AcneType
	}
	if len(result.FacialType) > 0 && !reflect.DeepEqual(result.FacialType, old_result.FacialType) {
		old_result.FacialType = result.FacialType
	}
	if result.SkinType != old_result.SkinType {
		old_result.SkinType = result.SkinType
	}
	if len(result.Skincare) > 0 && !reflect.DeepEqual(result.Skincare, old_result.Skincare) {
		old_result.Skincare = result.Skincare
	}

	return service.repo.UpdateResultById(id, old_result)
}

func (service *resultService) DeleteResultById(id int) error {
	return service.repo.DeleteResultById(id)
}

func (service *resultService) GetResultsByUserIdFromToken(token string) ([]entities.Result, error) {
	user_id, err := utils.ExtractToken(token)

	if err != nil {
		return []entities.Result{}, err
	}
	return service.repo.GetResultsByUserId(int(user_id))
}

func (service *resultService) GetResultsByUserId(user_id int) ([]entities.Result, error) {
	return service.repo.GetResultsByUserId(user_id)
}
