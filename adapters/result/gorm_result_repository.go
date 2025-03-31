package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
)

type GormResultRepository struct {
	db *gorm.DB
}

func NewGormResultRepository(db *gorm.DB) repositories.ResultsRepository {
	return &GormResultRepository{db: db}
}

func (repo *GormResultRepository) CreateResult(result entities.Result) (entities.Result, error) {
	err := repo.db.Create(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

func (repo *GormResultRepository) CreateSkincareResult(skincareResult entities.SkincareResult) (entities.SkincareResult, error) {
	err := repo.db.Create(&skincareResult).Error
	if err != nil {
		return entities.SkincareResult{}, err
	}

	return skincareResult, nil
}

func (repo *GormResultRepository) GetReuslt(id uint) (entities.Result, error) {
	var result entities.Result
	err := repo.db.Preload("User").Preload("Skincare.Skincare").Preload("Skin").Where("id = ?", id).First(&result).Error
	if err != nil {
		return entities.Result{}, err
	}

	return result, nil
}

func (repo *GormResultRepository) GetResults(user_id uint) ([]entities.Result, error) {
	var results []entities.Result
	err := repo.db.Preload("User").Preload("Skincare.Skincare").Preload("Skin").Where("user_id = ?", user_id).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}
