package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
)

type GormResultRepository struct {
	db *gorm.DB
}

func NewGormResultRepository(db *gorm.DB) repositories.ResultRepository {
	return &GormResultRepository{db: db}
}

func (repo *GormResultRepository) CreateResult(result entities.Result) (entities.Result, error) {
	err := repo.db.Create(&result).Error
	return result, err
}

func (repo *GormResultRepository) GetResults() ([]entities.Result, error) {
	var results []entities.Result
	err := repo.db.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (repo *GormResultRepository) GetResultById(id int) (entities.Result, error) {
	var result entities.Result
	err := repo.db.First(&result, id).Error
	return result, err
}

func (repo *GormResultRepository) UpdateResultById(id int, result entities.Result) (entities.Result, error) {
	err := repo.db.Model(&entities.Result{}).Where("id = ?", id).Updates(&result).Error
	return result, err
}

func (repo *GormResultRepository) DeleteResultById(id int) error {
	err := repo.db.Delete(&entities.Result{}, id).Error
	return err
}

func (repo *GormResultRepository) GetResultsByUserId(user_id int) ([]entities.Result, error) {
	var results []entities.Result
	err := repo.db.Where("user_id = ?", user_id).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}