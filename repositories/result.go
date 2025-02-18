package repositories

import (
	"github.com/Narutchai01/Project_S-BE/entities"
)

type ResultsRepository interface {
	CreateResult(entities.Result) (entities.Result, error)
	GetResults(id uint) ([]entities.Result, error)
	GetResult(id uint) (entities.Result, error)
	GetResultLatest(id uint) (entities.Result, error)
	GetResultByIDs(ids []uint) ([]entities.Result, error)
	// UpdateResult(result entities.Result, id uint) (entities.Result, error)
	// DeleteResult(uint) error
}
