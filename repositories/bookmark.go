package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type BookmarkRepository interface {
	CreateBookmarkThread(thread_id uint, user_id uint) (entities.BookmarkThread, error)
	FindBookMarkThread(thread_id uint, user_id uint) (entities.BookmarkThread, error)
	UpdateBookMarkThread(thread_id uint, user_id uint, status bool) (entities.BookmarkThread, error)
	BookmarkReviewSkincare(review_id uint, user_id uint) (entities.BookmarkReviewSkincare, error)
	FindBookMarkReviewSkincare(review_id uint, user_id uint) (entities.BookmarkReviewSkincare, error)
	UpdateBookMarkReviewSkincare(review_id uint, user_id uint, status bool) (entities.BookmarkReviewSkincare, error)

	FindBookmark(community_id uint, user_id uint) (bool, entities.Bookmark, error)
	Bookmark(community_id uint, user_id uint) (entities.Bookmark, error)
	DeleteBookmark(community_id uint, user_id uint) (entities.Bookmark, error)
}
