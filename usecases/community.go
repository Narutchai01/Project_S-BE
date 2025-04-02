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
	GetCommunitiesByUserID(user_id uint, type_community string, token string) ([]entities.Community, error)
	DeleteCommunity(id uint, token string, type_community string) error
}

type communityService struct {
	communityRepo repositories.CommunityRepository
	userRepo      repositories.UserRepository
	favoriteRepo  repositories.FavoriteRepository
	bookmarkRepo  repositories.BookmarkRepository
}

func NewCommunityUseCase(communityRepo repositories.CommunityRepository, userRepo repositories.UserRepository, favoriteRepo repositories.FavoriteRepository, bookmarkRepo repositories.BookmarkRepository) CommunityUseCase {
	return &communityService{communityRepo, userRepo, favoriteRepo, bookmarkRepo}
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

	// isFavorted, _, _ := service.favoriteRepo.FindFavorite(uint(community.ID), "community_id", user.ID)

	// community.Favorite = isFavorted

	// community.Likes = uint64(service.favoriteRepo.CountFavorite(community.ID, "community_id"))

	// isBookmark, _, _ := service.bookmarkRepo.FindBookmark(community.ID, user.ID)

	// community.Bookmark = isBookmark

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

	isFavorted, _, err := service.favoriteRepo.FindFavorite(uint(community.ID), "community_id", user.ID)
	if err != nil {
		return entities.Community{}, err
	}

	community.Favorite = isFavorted

	community.Likes = uint64(service.favoriteRepo.CountFavorite(community.ID, "community_id"))

	_, err = service.userRepo.FindFollower(community.User.ID, user.ID)
	if err != nil {
		user.Follow = false
	} else {
		user.Follow = true
	}

	isBookmarked, _, _ := service.bookmarkRepo.FindBookmark(community.ID, user.ID)

	community.Bookmark = isBookmarked

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

		isFavorted, _, err := service.favoriteRepo.FindFavorite(uint(community.ID), "community_id", user.ID)
		if err != nil {
			return []entities.Community{}, err
		}

		communities[i].Favorite = isFavorted

		communities[i].Likes = uint64(service.favoriteRepo.CountFavorite(community.ID, "community_id"))
	}

	return communities, nil

}

func (service *communityService) GetCommunitiesByUserID(target_user_id uint, type_community string, token string) ([]entities.Community, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return []entities.Community{}, err
	}

	user_target, err := service.userRepo.GetUser(target_user_id)
	if err != nil {
		return []entities.Community{}, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return []entities.Community{}, err
	}

	community_type, err := service.communityRepo.GetCommunityType(strings.ToLower(type_community))
	if err != nil {
		return []entities.Community{}, err
	}

	communities, err := service.communityRepo.GetCommunitiesByUserID(user_target.ID, uint(community_type.ID))
	if err != nil {
		return []entities.Community{}, err
	}

	for i, community := range communities {
		communities[i].Owner = (user.ID == community.User.ID)

		isFavorted, _, err := service.favoriteRepo.FindFavorite(uint(community.ID), "community_id", user.ID)
		if err != nil {
			return []entities.Community{}, err
		}

		communities[i].Favorite = isFavorted

		communities[i].Likes = uint64(service.favoriteRepo.CountFavorite(community.ID, "community_id"))
	}

	return communities, nil
}

func (service *communityService) DeleteCommunity(id uint, token string, type_community string) error {
	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return err
	}

	community_type, err := service.communityRepo.GetCommunityType(strings.ToLower(type_community))
	if err != nil {
		return err
	}

	community, err := service.communityRepo.GetCommunity(id, uint64(community_type.ID))
	if err != nil {
		return errors.New("community not found")
	}

	if community.UserID != uint64(user.ID) {
		return errors.New("you are not the owner of this community")
	}

	err = service.communityRepo.DeleteCommunity(community.ID)
	if err != nil {
		return err
	}

	return nil

}
