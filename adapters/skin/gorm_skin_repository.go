package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
)

type GormSkinRepository struct {
	db *gorm.DB
}

func NewGormSkinRepository(db *gorm.DB) repositories.SkinRepository {
	return &GormSkinRepository{db: db}
}

func (repo *GormSkinRepository) CreateSkin(skin entities.Skin) (entities.Skin, error) {
	err := repo.db.Create(&skin).Error
	return skin, err
}

func (repo *GormSkinRepository) GetSkins() ([]entities.Skin, error) {
	var skins []entities.Skin
	err := repo.db.Find(&skins).Error
	return skins, err
}

func (repo *GormSkinRepository) GetSkin(id int) (entities.Skin, error) {
	var skin entities.Skin
	err := repo.db.First(&skin, id).Error
	return skin, err
}

func (repo *GormSkinRepository) UpdateSkin(id int, skin entities.Skin) (entities.Skin, error) {
	err := repo.db.Model(&entities.Skin{}).Where("id = ?", id).Updates(skin).Error
	return skin, err
}

func (repo *GormSkinRepository) DeleteSkin(id int) error {
	err := repo.db.Delete(&entities.Skin{}, id).Error
	return err
}
