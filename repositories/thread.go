package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type ThreadRepository interface {
	CreateThread(user_id uint, title string, image string) (entities.Thread, error)
	CreateThreadDetail(threadDetail entities.ThreadDetail) (entities.ThreadDetail, error)
	GetThreads() ([]entities.Thread, error)
	GetThread(id uint) (entities.Thread, error)
	GetThreadDetails(thread_id uint) ([]entities.ThreadDetail, error)
	DeleteThread(thread_id uint) error
	// CreateBookmark(thread_id uint, user_id uint) (entities.Bookmark, error)
	// FindBookMark(thread_id uint, user_id uint) (entities.Bookmark, error)
	// UpdateBookMark(thread_id uint, user_id uint, status bool) (entities.Bookmark, error)
}
