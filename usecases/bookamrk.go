package usecases

import (
	"errors"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
	"gorm.io/gorm"
)

type BookmarkUseCase interface {
	BookmarkThread(thread_id uint, token string) (entities.BookmarkThread, error)
	BookmarkReviewSkincare(review_id uint, token string) (entities.BookmarkReviewSkincare, error)
	BookmarkCommunity(community_id uint, token string, type_community string) (entities.Bookmark, error)
}
type bookmarkService struct {
	repo          repositories.BookmarkRepository
	threadRepo    repositories.ThreadRepository
	reviewRepo    repositories.ReviewRepository
	userRepo      repositories.UserRepository
	communityRepo repositories.CommunityRepository
}

func NewBookmarkUseCase(repo repositories.BookmarkRepository, threadRepo repositories.ThreadRepository, reviewRepo repositories.ReviewRepository, userRepo repositories.UserRepository, communityRepo repositories.CommunityRepository) BookmarkUseCase {
	return &bookmarkService{repo, threadRepo, reviewRepo, userRepo, communityRepo}
}

func (service *bookmarkService) BookmarkThread(thread_id uint, token string) (entities.BookmarkThread, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.BookmarkThread{}, errors.New("token expired")
	}

	_, err = service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.BookmarkThread{}, errors.New("user not found")
	}

	_, err = service.threadRepo.GetThread(thread_id)
	if err != nil {
		return entities.BookmarkThread{}, errors.New("thread not found")
	}

	bookmark, err := service.repo.FindBookMarkThread(thread_id, user_id)
	if err != nil {
		return service.repo.CreateBookmarkThread(thread_id, user_id)
	}

	status := !bookmark.Status

	return service.repo.UpdateBookMarkThread(thread_id, user_id, status)
}

func (service *bookmarkService) BookmarkReviewSkincare(review_id uint, token string) (entities.BookmarkReviewSkincare, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.BookmarkReviewSkincare{}, err
	}

	_, err = service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.BookmarkReviewSkincare{}, errors.New("user not found")
	}

	_, err = service.reviewRepo.GetReviewSkincare(review_id)
	if err != nil {
		return entities.BookmarkReviewSkincare{}, errors.New("review not found")
	}

	bookmark, err := service.repo.FindBookMarkReviewSkincare(review_id, user_id)
	if err != nil {
		return service.repo.BookmarkReviewSkincare(review_id, user_id)
	}

	status := !bookmark.Status

	return service.repo.UpdateBookMarkReviewSkincare(review_id, user_id, status)
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
		} else {
			return entities.Bookmark{}, err
		}
	} else if isBookmarked {
		bookmark, err = service.repo.DeleteBookmark(community.ID, user.ID)
		if err != nil {
			return entities.Bookmark{}, err
		}
		bookmark.Status = false
	}

	return bookmark, nil
}
