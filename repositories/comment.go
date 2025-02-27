package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type CommentRepository interface {
	CreateCommentThread(comment entities.CommentThread) (entities.CommentThread, error)
	GetCommentsThread(thread_id uint) ([]entities.CommentThread, error)
	CreateCommentReviewSkicnare(comment entities.FavoriteCommentReview) (entities.FavoriteCommentReview, error)
}
