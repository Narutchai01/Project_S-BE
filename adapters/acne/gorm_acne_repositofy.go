package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
)

type GormAcneRepository struct {
	db *gorm.DB
}

func NewGormAcneRepository(db *gorm.DB) repositories.AcneRepository {
	return &GormAcneRepository{db: db}
}

func (repo *GormAcneRepository) CreateAcne(acne entities.Acne) (entities.Acne, error) {
	err := repo.db.Create(&acne).Error
	return acne, err
}
