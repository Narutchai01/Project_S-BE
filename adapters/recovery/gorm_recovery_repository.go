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
	return &GormRecoveryRepository{db: db}
}

func (repo *GormRecoveryRepository) CreateRecovery(recovery entities.Recovery) (entities.Recovery, error) {
	err := repo.db.Create(&recovery).Error
	return recovery, err
}

func (repo *GormRecoveryRepository) FindRecoveryByOTP(otp string, user_id uint) (entities.Recovery, error) {
	var recovery entities.Recovery
	err := repo.db.Where("otp = ? and user_id = ?", otp, user_id).First(&recovery).Error
	if err != nil {
		return entities.Recovery{}, err
	}
	return recovery, nil
}

func (repo *GormRecoveryRepository) DeleteRecoveryById(id uint) error {
	err := repo.db.Delete(&entities.Recovery{}, id).Error
	return err
}
