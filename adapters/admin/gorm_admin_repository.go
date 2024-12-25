package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
)

type GormAdminRepository struct {
	db *gorm.DB
}

func NewGormAdminRepository(db *gorm.DB) repositories.AdminRepository {
	return &GormAdminRepository{db: db}
}

func (repo *GormAdminRepository) CreateAdmin(admin entities.Admin) (entities.Admin, error) {
	err := repo.db.Create(&admin).Error
	return admin, err
}

func (repo *GormAdminRepository) GetAdmins() ([]entities.Admin, error) {
	var admins []entities.Admin
	err := repo.db.Find(&admins).Error
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (repo *GormAdminRepository) GetAdmin(id int) (entities.Admin, error) {
	var admin entities.Admin
	err := repo.db.First(&admin, id).Error
	return admin, err
}

func (repo *GormAdminRepository) UpdateAdmin(id int, admin entities.Admin) (entities.Admin, error) {
	err := repo.db.Model(&entities.Admin{}).Where("id = ?", id).Updates(&admin).Error
	return admin, err
}

func (repo *GormAdminRepository) DeleteAdmin(id int) (entities.Admin, error) {
	err := repo.db.Where("id = ?", id).Delete(&entities.Admin{}).Error
	return entities.Admin{}, err
}

func (repo *GormAdminRepository) GetAdminByEmail(email string) (entities.Admin, error) {
	var admin entities.Admin
	err := repo.db.Where("email = ?", email).First(&admin).Error
	return admin, err
}
