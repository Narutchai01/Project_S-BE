package usecases

import (
	"errors"
	"mime/multipart"
	"os"
	"strings"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CommunityUseCase interface {
	CreateCommunityThread(community entities.Community, token string, files []*multipart.FileHeader, c *fiber.Ctx, type_community string) (entities.Community, error)
	GetCommunity(id uint, type_community string, token string) (entities.Community, error)
	GetCommunities(type_community string, token string) ([]entities.Community, error)
}

type communityService struct {
	communityRepo repositories.CommunityRepository
	userRepo      repositories.UserRepository
}

func NewCommunityUseCase(communityRepo repositories.CommunityRepository, userRepo repositories.UserRepository) CommunityUseCase {
	return &communityService{communityRepo, userRepo}
}

func (service *communityService) CreateCommunityThread(community entities.Community, token string, files []*multipart.FileHeader, c *fiber.Ctx, community_type string) (entities.Community, error) {
	var ImageURLs []string

	type_community, err := service.communityRepo.GetCommunityType(strings.ToLower(community_type))
	if err != nil {
		return entities.Community{}, err
	}

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.Community{}, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.Community{}, err
	}

	community.UserID = uint64(user.ID)
	community.TypeID = uint64(type_community.ID)

	for _, file := range files {
		fileName := uuid.New().String() + ".jpg"

		if err := utils.CheckDirectoryExist(); err != nil {
			return entities.Community{}, err
		}

		if err := c.SaveFile(file, "./uploads/"+fileName); err != nil {
			return entities.Community{}, err
		}

		imageUrl, err := utils.UploadImage(fileName, "/thread")

		if err != nil {
			return entities.Community{}, err
		}

		err = os.Remove("./uploads/" + fileName)

		if err != nil {
			return entities.Community{}, err
		}

		ImageURLs = append(ImageURLs, imageUrl)
	}

	if len(ImageURLs) != len(files) {
		return entities.Community{}, err
	}

	community, err = service.communityRepo.CreateCommunity(community)
	if err != nil {
		return entities.Community{}, err
	}

	for _, imageUrl := range ImageURLs {
		image := entities.CommunityImage{
			CommunityID: uint64(community.ID),
			Image:       imageUrl,
		}
		_, err = service.communityRepo.CreateCommunityImage(image)
		if err != nil {
			return entities.Community{}, err
		}
	}

	if len(community.SkincareID) > 1 {
		for _, skincare_id := range community.SkincareID {
			err = service.communityRepo.CreateSkincareCommunity(community.ID, uint(skincare_id))
			if err != nil {
				return entities.Community{}, err
			}
		}
	}

	community, err = service.communityRepo.GetCommunity(uint(community.ID), community.TypeID)
	if err != nil {
		return entities.Community{}, err
	}
	community.Owner = (user.ID == community.User.ID)

	return community, nil
}

func (service *communityService) GetCommunity(id uint, type_community string, token string) (entities.Community, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.Community{}, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.Community{}, errors.New("user not found")
	}

	community_type, err := service.communityRepo.GetCommunityType(strings.ToLower(type_community))
	if err != nil {
		return entities.Community{}, err
	}

	community, err := service.communityRepo.GetCommunity(id, uint64(community_type.ID))
	if err != nil {
		return entities.Community{}, errors.New("community not found")
	}

	community.Owner = (user.ID == community.User.ID)

	return community, nil
}

func (service *communityService) GetCommunities(type_community string, token string) ([]entities.Community, error) {
	community_type, err := service.communityRepo.GetCommunityType(strings.ToLower(type_community))
	if err != nil {
		return []entities.Community{}, err
	}
	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return []entities.Community{}, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return []entities.Community{}, err
	}

	communities, err := service.communityRepo.GetCommunities(uint64(community_type.ID))
	if err != nil {
		return []entities.Community{}, err
	}

	for i, community := range communities {
		communities[i].Owner = (user.ID == community.User.ID)
	}

	return communities, nil

}
