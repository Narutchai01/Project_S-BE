package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
)

type GormCommunityRepository struct {
	db *gorm.DB
}

func NewGormCommunityRepository(db *gorm.DB) repositories.CommunityRepository {
	return &GormCommunityRepository{db: db}
}

func (repo *GormCommunityRepository) CreateCommunity(community entities.Community) (entities.Community, error) {
	err := repo.db.Create(&community).Error
	return community, err
}

func (repo *GormCommunityRepository) CreateCommunityImage(communityImage entities.CommunityImage) (entities.CommunityImage, error) {
	err := repo.db.Create(&communityImage).Error
	return communityImage, err
}

func (repo *GormCommunityRepository) GetCommunity(community_id uint, typeID uint64) (entities.Community, error) {
	var community entities.Community
	err := repo.db.Preload("Images").Preload("Skincares").Preload("Skincares.Skincare").Preload("User").Where("id = ? AND type_id = ?", community_id, typeID).First(&community).Error
	return community, err
}

func (repo *GormCommunityRepository) GetCommunityType(type_community string) (entities.CommunityType, error) {
	var community_type entities.CommunityType
	err := repo.db.Where("type = ?", type_community).First(&community_type).Error
	return community_type, err
}

func (repo *GormCommunityRepository) GetCommunities(typeID uint64) ([]entities.Community, error) {
	var communities []entities.Community
	err := repo.db.Preload("Images").Preload("Skincares").Preload("Skincares.Skincare").Preload("User").Where("type_id = ?", typeID).Find(&communities).Error
	return communities, err
}

func (repo *GormCommunityRepository) CreateSkincareCommunity(community_id uint, skincare_id uint) error {
	var community_skincare entities.SkincareCommunity

	community_skincare.CommunityID = uint64(community_id)
	community_skincare.SkincareID = uint64(skincare_id)

	return repo.db.Create(&community_skincare).Error
}

func (repo *GormCommunityRepository) GetCommunitiesByUserID(user_id uint, type_id uint) ([]entities.Community, error) {
	var communities []entities.Community
	err := repo.db.Preload("Images").Preload("Skincares").Preload("Skincares.Skincare").Preload("User").Where("user_id = ? AND type_id = ?", user_id, type_id).Find(&communities).Error
	return communities, err
}

func (repo *GormCommunityRepository) DeleteCommunity(community_id uint) error {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("community_id = ?", community_id).Delete(&entities.SkincareCommunity{}).Error; err != nil {
			return err
		}
		if err := tx.Where("community_id = ?", community_id).Delete(&entities.CommunityImage{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&entities.Community{}, community_id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (repo *GormCommunityRepository) UpdateCommunity(community_id uint, community *entities.Community) error {
	if err := repo.db.Model(&entities.Community{}).Where("id = ?", community_id).Updates(&community).Error; err != nil {
		return err
	}
	return nil
}

func (repo *GormCommunityRepository) DeleteCommunityImage(image_id uint, community_id uint) error {
	if err := repo.db.Where("community_id = ? AND id = ?", community_id, image_id).Delete(&entities.CommunityImage{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *GormCommunityRepository) DeleteSkincareCommunity(community_id uint, skincare_id uint) error {
	if err := repo.db.Where("community_id = ? AND skincare_id = ?", community_id, skincare_id).Delete(&entities.SkincareCommunity{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *GormCommunityRepository) FindSkincareCommunity(community_id uint, skincare_id uint) error {
	var community_skincare entities.SkincareCommunity
	err := repo.db.Where("community_id = ? AND skincare_id = ?", community_id, skincare_id).First(&community_skincare).Error
	if err != nil {
		return err
	}
	return nil
}
