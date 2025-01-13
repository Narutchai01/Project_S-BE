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

func (repo *GormAcneRepository) GetAcnes() ([]entities.Acne, error) {
	var acnes []entities.Acne
	err := repo.db.Find(&acnes).Error
	return acnes, err
}

func (repo *GormAcneRepository) GetAcne(id int) (entities.Acne, error) {
	var acne entities.Acne
	err := repo.db.First(&acne, id).Error
	return acne, err
}

func (repo *GormAcneRepository) DeleteAcne(id int) error {
	err := repo.db.Delete(&entities.Acne{}, id).Error
	return err
}
