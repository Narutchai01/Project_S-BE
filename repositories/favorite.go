package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type FavoriteRepository interface {
	FavoriteCommentThread(comment_id uint, user_id uint) (entities.FavoriteCommentThread, error)
	FindFavoriteCommentThread(comment_id uint, user_id uint) (entities.FavoriteCommentThread, error)
	UpdateFavoriteCommentThread(favorite_comment entities.FavoriteCommentThread) (entities.FavoriteCommentThread, error)
	CountFavoriteCommentThread(comment_id uint) (int64, error)

	FavoriteThread(thread_id uint, user_id uint) (entities.FavoriteThread, error)
	FindFavoriteThread(thread_id uint, user_id uint) (entities.FavoriteThread, error)
	UpdateFavoriteThread(favorite_thread entities.FavoriteThread) (entities.FavoriteThread, error)
	CountFavoriteThread(thread_id uint) (int64, error)

	FavoriteReviewSkincare(review_skincare_id uint, user_id uint) (entities.FavoriteReviewSkincare, error)
	FindFavoriteReviewSkincare(review_skincare_id uint, user_id uint) (entities.FavoriteReviewSkincare, error)
	UpdateFavoriteReviewSkincare(favorite_review_skincare entities.FavoriteReviewSkincare) (entities.FavoriteReviewSkincare, error)
	CountFavoriteReviewSkincare(review_skincare_id uint) (int64, error)

	FavoriteCommentReviewSkincare(comment_id uint, user_id uint) (entities.FavoriteCommentReviewSkincare, error)
	FindFavoriteCommentReviewSkincare(comment_id uint, user_id uint) (entities.FavoriteCommentReviewSkincare, error)
	UpdateFavoriteCommentReviewSkincare(favorite_comment_review_skincare entities.FavoriteCommentReviewSkincare) (entities.FavoriteCommentReviewSkincare, error)
	CountFavoriteCommentReviewSkincare(comment_id uint) (int64, error)
}
