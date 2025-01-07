package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
)

type GormRecoveryRepository struct {
	db *gorm.DB
}

func NewGormRecoveryRepository(db *gorm.DB) repositories.RecoveryRepository {
	return GormRecoveryRepository{db: db}
}

func (repo GormRecoveryRepository) CreateRecovery(recovery entities.Recovery) (entities.Recovery, error) {
	err := repo.db.Create(&recovery).Error
	return recovery, err
}

