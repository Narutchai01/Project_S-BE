package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
)

type GormFaceProblemRepository struct {
	db *gorm.DB
}

func NewGormFaceProblemRepository(db *gorm.DB) repositories.FaceProblemRepository {
	return &GormFaceProblemRepository{db: db}
}

func (repo *GormFaceProblemRepository) GetFaceProblemType(type_problem string) (entities.FaceProblemType, error) {
	var face_problem_type entities.FaceProblemType
	err := repo.db.Where("name = ?", type_problem).First(&face_problem_type).Error
	return face_problem_type, err
}

func (repo *GormFaceProblemRepository) CreateFaceProblem(faceProblem entities.FaceProblem) (entities.FaceProblem, error) {

	err := repo.db.Create(&faceProblem).Error
	if err != nil {
		return entities.FaceProblem{}, err
	}

	return faceProblem, nil
}

func (repo *GormFaceProblemRepository) GetFaceProblem(id uint64) (entities.FaceProblem, error) {
	var face_problem entities.FaceProblem
	err := repo.db.Preload("Type").Preload("Admin").First(&face_problem, id).Error
	return face_problem, err
}

func (repo *GormFaceProblemRepository) GetFaceProblems(type_id uint64) ([]entities.FaceProblem, error) {
	var face_problems []entities.FaceProblem
	err := repo.db.Preload("Type").Preload("Admin").Where("type_id = ?", type_id).Find(&face_problems).Error
	return face_problems, err
}

func (repo *GormFaceProblemRepository) UpdateFaceProblem(id uint64, problem entities.FaceProblem) (entities.FaceProblem, error) {
	err := repo.db.Model(&entities.FaceProblem{}).Where("id = ?", id).Updates(&problem).Error
	return problem, err
}

func (repo *GormFaceProblemRepository) DeleteFaceProblem(id int) error {
	err := repo.db.Delete(&entities.FaceProblem{}, id).Error
	return err
}
