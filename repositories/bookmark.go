package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type BookmarkRepository interface {
	CreateBookmark(thread_id uint, user_id uint) (entities.Bookmark, error)
	FindBookMark(thread_id uint, user_id uint) (entities.Bookmark, error)
	UpdateBookMark(thread_id uint, user_id uint, status bool) (entities.Bookmark, error)
}
