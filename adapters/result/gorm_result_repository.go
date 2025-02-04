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
	err := repo.db.Preload("Skincare").Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

// func (repo *GormResultRepository) GetResults() ([]presentation.Result, error) {
// 	var results []presentation.Result
// 	err := repo.db.Raw(`
//         SELECT r.*, json_agg(s.*) AS skincare
//         FROM results r
//         LEFT JOIN skincare s ON s.id = ANY(r.skincare)
//         GROUP BY r.id
//     `).Scan(&results).Error

// 	return results, err
// }

func (repo *GormResultRepository) GetResultById(id int) (entities.Result, error) {
	var result entities.Result
	err := repo.db.First(&result, id).Error
	return result, err
	// var result presentation.Result
	// err := repo.db.Raw(`SELECT r.*, s.* FROM results r JOIN skincare s ON s.id = ANY(r.skincare) WHERE r.id = `)
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

func (repo *GormResultRepository) GetLatestResultByUserIdFromToken(user_id int) (entities.Result, error) {
	var result entities.Result
	err := repo.db.Where("user_id = ?", user_id).Last(&result).Error
	return result, err
}
