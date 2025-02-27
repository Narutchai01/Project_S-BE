package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
)

type GormSkincareRepository struct {
	db *gorm.DB
}

func NewGormSkincareRepository(db *gorm.DB) repositories.SkincareRepository {
	return &GormSkincareRepository{db: db}
}

func (repo *GormSkincareRepository) CreateSkincare(skincare entities.Skincare) (entities.Skincare, error) {
	err := repo.db.Create(&skincare).Error
	return skincare, err
}

func (repo *GormSkincareRepository) GetSkincares() ([]entities.Skincare, error) {
	var skincares []entities.Skincare
	err := repo.db.Find(&skincares).Error
	if err != nil {
		return nil, err
	}
	return skincares, nil
}

func (repo *GormSkincareRepository) GetSkincareById(id int) (entities.Skincare, error) {
	var skincare entities.Skincare
	err := repo.db.First(&skincare, id).Error
	return skincare, err
}

func (repo *GormSkincareRepository) UpdateSkincareById(id int, skincare entities.Skincare) (entities.Skincare, error) {
	err := repo.db.Model(&entities.Skincare{}).Where("id = ?", id).Updates(&skincare).Error
	return skincare, err
}

func (repo *GormSkincareRepository) DeleteSkincareById(id int) (entities.Skincare, error) {
	err := repo.db.Where("id = ?", id).Delete(&entities.Skincare{}).Error
	return entities.Skincare{}, err
}

func (repo *GormSkincareRepository) GetSkincareByIds(ids []int) ([]entities.Skincare, error) {
	var skincares []entities.Skincare

	err := repo.db.Find(&skincares, ids).Error

	return skincares, err
}
