package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type ResultRepository interface {
	CreateResult(result entities.Result) (entities.Result, error)
	GetResults() ([]entities.Result, error)
	GetResultById(id int) (entities.Result, error)
}