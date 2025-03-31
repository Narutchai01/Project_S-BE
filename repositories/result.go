package repositories

import (
	"github.com/Narutchai01/Project_S-BE/entities"
)

type ResultsRepository interface {
	CreateResult(entities.Result) (entities.Result, error)
	CreateSkincareResult(entities.SkincareResult) (entities.SkincareResult, error)
	GetResult(id uint) (entities.Result, error)
	GetResults(user_id uint) ([]entities.Result, error)
}
