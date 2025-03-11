package usecases

import (
	"errors"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
)

type BookmarkUseCase interface {
	BookmarkThread(thread_id uint, token string) (entities.BookmarkThread, error)
	BookmarkReviewSkincare(review_id uint, token string) (entities.BookmarkReviewSkincare, error)
}
type bookmarkService struct {
	repo       repositories.BookmarkRepository
	threadRepo repositories.ThreadRepository
	reviewRepo repositories.ReviewRepository
	userRepo   repositories.UserRepository
}

func NewBookmarkUseCase(repo repositories.BookmarkRepository, threadRepo repositories.ThreadRepository, reviewRepo repositories.ReviewRepository, userRepo repositories.UserRepository) BookmarkUseCase {
	return &bookmarkService{repo, threadRepo, reviewRepo, userRepo}
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
