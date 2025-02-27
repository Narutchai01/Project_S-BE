package usecases

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
)

type BookmarkUseCase interface {
	BookmarkThread(thread_id uint, token string) (entities.BookmarkThread, error)
	BookmarkReviewSkincare(review_id uint, token string) (entities.BookmarkReviewSkincare, error)
}
type bookmarkService struct {
	repo repositories.BookmarkRepository
}

func NewBookmarkUseCase(repo repositories.BookmarkRepository) BookmarkUseCase {
	return &bookmarkService{repo}
}

func (service *bookmarkService) BookmarkThread(thread_id uint, token string) (entities.BookmarkThread, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.BookmarkThread{}, err
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

	bookmark, err := service.repo.FindBookMarkReviewSkincare(review_id, user_id)
	if err != nil {
		return service.repo.BookmarkReviewSkincare(review_id, user_id)
	}

	status := !bookmark.Status

	return service.repo.UpdateBookMarkReviewSkincare(review_id, user_id, status)
}
