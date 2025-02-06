package repositories

import (
	"github.com/Narutchai01/Project_S-BE/entities"
)

type ResultRepository interface {
	CreateResult(result entities.Result) (entities.ResultWithSkincare, error)
	GetResults() ([]entities.ResultWithSkincare, error)
	GetResultById(id int) (entities.ResultWithSkincare, error)
	GetResultByIdWithoutSkincare(id int) (entities.Result, error)
	UpdateResultById(id int, result entities.Result) (entities.ResultWithSkincare, error)
	DeleteResultById(id int) error
	GetResultsByUserId(user_id int) ([]entities.ResultWithSkincare, error)
	GetLatestResultByUserIdFromToken(user_id int) (entities.ResultWithSkincare, error)
}