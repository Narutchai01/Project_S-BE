package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/comment"
	adaptersCommunity "github.com/Narutchai01/Project_S-BE/adapters/community"
	adaptersFavorite "github.com/Narutchai01/Project_S-BE/adapters/favorite"
	adaptersUser "github.com/Narutchai01/Project_S-BE/adapters/user"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CommentRouters(app fiber.Router, db *gorm.DB) {
	commentRepo := adapters.NewGormCommentRepository(db)
	favoriteCommentRepo := adaptersFavorite.NewGormFavoriteRepository(db)
	userRepo := adaptersUser.NewGormUserRepository(db)
	communityRepo := adaptersCommunity.NewGormCommunityRepository(db)
	commentService := usecases.NewCommentUseCase(commentRepo, favoriteCommentRepo, userRepo, communityRepo)
	commentHandler := adapters.NewHttpCommentHandler(commentService)

	commentGroup := app.Group("/comment")

	threadGroup := commentGroup.Group("/thread")
	threadGroup.Post("/", commentHandler.CreateCommentThread)
	threadGroup.Get("/:thread_id", commentHandler.GetCommentsThread)

	reviewGroup := commentGroup.Group("/reviews")
	reviewSkincareGroup := reviewGroup.Group("/skincare")
	reviewSkincareGroup.Post("/", commentHandler.CreateCommentReviewSkicnare)
	reviewSkincareGroup.Get("/:review_id", commentHandler.HandleGetCommentReviewSkincare)

}
