package usecases

import (
	"errors"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
	"gorm.io/gorm"
)

type BookmarkUseCase interface {
	BookmarkCommunity(community_id uint, token string, type_community string) (entities.Bookmark, error)
	GetCommunitiesBookmark(user_id uint, token string) ([]entities.Bookmark, error)
}
type bookmarkService struct {
	repo          repositories.BookmarkRepository
	userRepo      repositories.UserRepository
	communityRepo repositories.CommunityRepository
	favoriteRepo  repositories.FavoriteRepository
}

func NewBookmarkUseCase(repo repositories.BookmarkRepository, userRepo repositories.UserRepository, communityRepo repositories.CommunityRepository, favoriteRepo repositories.FavoriteRepository) BookmarkUseCase {
	return &bookmarkService{repo, userRepo, communityRepo, favoriteRepo}
}

func (service *bookmarkService) BookmarkCommunity(community_id uint, token string, type_community string) (entities.Bookmark, error) {
	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.Bookmark{}, errors.New("invalid token")
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.Bookmark{}, errors.New("user not found")
	}

	community_type, err := service.communityRepo.GetCommunityType(type_community)
	if err != nil {
		return entities.Bookmark{}, errors.New("community type not found")
	}

	community, err := service.communityRepo.GetCommunity(community_id, uint64(community_type.ID))
	if err != nil {
		return entities.Bookmark{}, errors.New("community not found")
	}

	isBookmarked, bookmark, err := service.repo.FindBookmark(community.ID, user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			bookmark, err = service.repo.Bookmark(community.ID, user.ID)
			if err != nil {
				return entities.Bookmark{}, err
			}
			bookmark.Status = true
		} else {
			return entities.Bookmark{}, err
		}
	} else if isBookmarked {
		err = service.repo.DeleteBookmark(community.ID, user.ID)
		if err != nil {
			return entities.Bookmark{}, err
		}
		bookmark.Status = false
	}

	return bookmark, nil
}

func (service *bookmarkService) GetCommunitiesBookmark(user_id uint, token string) ([]entities.Bookmark, error) {

	usered_id, err := utils.ExtractToken(token)
	if err != nil {
		return []entities.Bookmark{}, err
	}

	user_bookmark, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return []entities.Bookmark{}, err
	}

	user, err := service.userRepo.GetUser(usered_id)
	if err != nil {
		return []entities.Bookmark{}, err
	}

	bookmarks, err := service.repo.GetCommunitiesBookmark(int(user_bookmark.ID))
	if err != nil {
		return nil, err
	}

	for i, bookmark := range bookmarks {
		community := bookmark.Community
		isFavorted, _, _ := service.favoriteRepo.FindFavorite(community.ID, "community_id", user.ID)
		bookmarks[i].Community.Favorite = isFavorted
	}

	return bookmarks, nil
}
