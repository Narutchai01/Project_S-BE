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
