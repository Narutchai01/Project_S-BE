package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (repo *GormRecoveryRepository) DeleteRecoveryById(id int) (entities.Recovery, error) {
	err := repo.db.Where("id = ?", id).Delete(&entities.Recovery{}).Error
	return entities.Recovery{}, err
}

func (repo *GormRecoveryRepository) GetRecoveries() ([]entities.Recovery, error) {
	var recoveries []entities.Recovery
	err := repo.db.Find(&recoveries).Error
	if err != nil {
		return nil, err
	}
	return recoveries, nil
}

func (repo *GormRecoveryRepository) GetRecoveryById(id int) (entities.Recovery, error) {
	var recovery entities.Recovery
	err := repo.db.First(&recovery, id).Error
	return recovery, err
}

func (repo *GormRecoveryRepository) GetRecoveryByUserId(user_id int) (entities.Recovery, error) {
	var recovery entities.Recovery
	err := repo.db.Where("user_id = ?", user_id).First(&recovery).Error
	return recovery, err
}

func (repo *GormRecoveryRepository) UpdateRecoveryOtpById(id int, otp string) (entities.Recovery, error) {
	var recovery entities.Recovery
	err := repo.db.Model(&recovery).Clauses(clause.Returning{}).Where("id = ?", id).Update("otp", otp).Error
	return recovery, err
}