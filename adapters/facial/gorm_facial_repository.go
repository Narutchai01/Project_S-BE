package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
)

type GormFacialRepository struct {
	db *gorm.DB
}

func NewGormFacialRepository(db *gorm.DB) repositories.FacialRepository {
	return &GormFacialRepository{db: db}
}

func (repo *GormFacialRepository) CreateFacial(facial entities.Facial) (entities.Facial, error) {
	err := repo.db.Create(&facial).Error
	return facial, err
}

func (repo *GormFacialRepository) GetFacials() ([]entities.Facial, error) {
	var facials []entities.Facial
	err := repo.db.Find(&facials).Error
	return facials, err
}
