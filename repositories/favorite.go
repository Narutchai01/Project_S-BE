package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type FavoriteRepository interface {
	FavoriteComment(comment_id uint, user_id uint) (entities.FavoriteComment, error)
	FindFavoriteComment(comment_id uint, user_id uint) (entities.FavoriteComment, error)
	UpdateFavoriteComment(favorite_comment entities.FavoriteComment) (entities.FavoriteComment, error)
	FavoriteThread(thread_id uint, user_id uint) (entities.FavoriteThread, error)
	FindFavoriteThread(thread_id uint, user_id uint) (entities.FavoriteThread, error)
	UpdateFavoriteThread(favorite_thread entities.FavoriteThread) (entities.FavoriteThread, error)
	CountFavoriteThread(thread_id uint) (int64, error)
	FavoriteReviewSkincare(review_skincare_id uint, user_id uint) (entities.FavoriteReviewSkincare, error)
	FindFavoriteReviewSkincare(review_skincare_id uint, user_id uint) (entities.FavoriteReviewSkincare, error)
	UpdateFavoriteReviewSkincare(favorite_review_skincare entities.FavoriteReviewSkincare) (entities.FavoriteReviewSkincare, error)
	CountFavoriteReviewSkincare(review_skincare_id uint) (int64, error)
}
