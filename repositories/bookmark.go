package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type BookmarkRepository interface {
	FindBookmark(community_id uint, user_id uint) (bool, entities.Bookmark, error)
	Bookmark(community_id uint, user_id uint) (entities.Bookmark, error)
	DeleteBookmark(community_id uint, user_id uint) error
	GetCommunitiesBookmark(user_id int) ([]entities.Bookmark, error)
}
