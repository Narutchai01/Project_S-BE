package usecases

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
)

type BookmarkUseCase interface {
	BookmarkThread(thread_id uint, token string) (entities.Bookmark, error)
}
type bookmarkService struct {
	repo repositories.BookmarkRepository
}

func NewBookmarkUseCase(repo repositories.BookmarkRepository) BookmarkUseCase {
	return &bookmarkService{repo}
}

func (service *bookmarkService) BookmarkThread(thread_id uint, token string) (entities.Bookmark, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.Bookmark{}, err
	}

	bookmark, err := service.repo.FindBookMark(thread_id, user_id)
	if err != nil {
		return service.repo.CreateBookmark(thread_id, user_id)
	}

	status := !bookmark.Status

	return service.repo.UpdateBookMark(thread_id, user_id, status)
}
