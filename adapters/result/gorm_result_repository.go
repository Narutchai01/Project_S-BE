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
	if err := repo.db.Create(&result).Error; err != nil {
		return entities.Result{}, err
	}
	err := repo.db.Preload("Skincare", 
		func(db *gorm.DB) *gorm.DB {
			return db.Omit("Admin")
	    	}).First(&result, result.ID).Error
	return result, err
}

func (repo *GormResultRepository) GetResults() ([]entities.Result, error) {
	var results []entities.Result
	err := repo.db.Preload("Skincare", 
		func(db *gorm.DB) *gorm.DB {
      		return db.Omit("Admin")
    		}).Find(&results).Error

	return results, err
}

func (repo *GormResultRepository) GetResultById(id int) (entities.Result, error) {
	var result entities.Result
	err := repo.db.Preload("Skincare", 
		func(db *gorm.DB) *gorm.DB {
			return db.Omit("Admin")
	    	}).First(&result, id).Error
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
	err := repo.db.Where("user_id = ?", user_id).Preload("Skincare", 
		func(db *gorm.DB) *gorm.DB {
			return db.Omit("Admin")
	    	}).Find(&results).Error
	return results, err
}

func (repo *GormResultRepository) GetLatestResultByUserIdFromToken(user_id int) (entities.Result, error) {
	var result entities.Result
	err := repo.db.Where("user_id = ?", user_id).Preload("Skincare", 
		func(db *gorm.DB) *gorm.DB {
			return db.Omit("Admin")
	    	}).Last(&result).Error
	return result, err
}
