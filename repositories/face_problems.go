package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type FaceProblemRepository interface {
	CreateFaceProblem(faceProblem entities.FaceProblem) (entities.FaceProblem, error)
	GetFaceProblemType(type_problem string) (entities.FaceProblemType, error)
	GetFaceProblem(id uint64) (entities.FaceProblem, error)
	GetFaceProblems(type_id uint64) ([]entities.FaceProblem, error)
	UpdateFaceProblem(id uint64, problem entities.FaceProblem) (entities.FaceProblem, error)
	DeleteFaceProblem(id int) error
}
