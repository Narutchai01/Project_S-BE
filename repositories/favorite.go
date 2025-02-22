package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type FavoriteRepository interface {
	FavoriteComment(comment_id uint, user_id uint) (entities.FavoriteComment, error)
	FindFavoriteComment(comment_id uint, user_id uint) (entities.FavoriteComment, error)
	UpdateFavoriteComment(favorite_comment entities.FavoriteComment) (entities.FavoriteComment, error)
	FavoriteThread(thread_id uint, user_id uint) (entities.FavoriteThread, error)
	FindFavoriteThread(thread_id uint, user_id uint) (entities.FavoriteThread, error)
	UpdateFavoriteThread(favorite_thread entities.FavoriteThread) (entities.FavoriteThread, error)
}
