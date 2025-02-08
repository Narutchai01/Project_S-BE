package repositories

import (
	"github.com/Narutchai01/Project_S-BE/entities"
)

type ResultsRepository interface {
	CreateResult(entities.Result) (entities.Result, error)
	GetResults(id uint) ([]entities.Result, error)
	GetResult(uint) (entities.Result, error)
}
