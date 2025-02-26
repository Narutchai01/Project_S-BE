package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type BookmarkRepository interface {
	CreateBookmarkThread(thread_id uint, user_id uint) (entities.BookmarkThread, error)
	FindBookMarkThread(thread_id uint, user_id uint) (entities.BookmarkThread, error)
	UpdateBookMarkThread(thread_id uint, user_id uint, status bool) (entities.BookmarkThread, error)
}
