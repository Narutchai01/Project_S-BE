package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type CommunityRepository interface {
	CreateCommunity(community entities.Community) (entities.Community, error)
	CreateCommunityImage(communityImage entities.CommunityImage) (entities.CommunityImage, error)
	GetCommunity(community_id uint, typeID uint64) (entities.Community, error)
	GetCommunityType(type_community string) (entities.CommunityType, error)
	GetCommunities(typeID uint64) ([]entities.Community, error)
	CreateSkincareCommunity(community_id uint, skincare_id uint) error
	GetCommunitiesByUserID(user_id uint, type_id uint) ([]entities.Community, error)
	DeleteCommunity(community_id uint) error
	UpdateCommunity(community_id uint, community *entities.Community) error
	DeleteCommunityImage(image_id uint, community_id uint) error
	DeleteSkincareCommunity(community_id uint, skincare_id uint) error
	FindSkincareCommunity(community_id uint, skincare_id uint) error
}
