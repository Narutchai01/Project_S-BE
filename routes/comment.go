package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/comment"
	adaptersFavorite "github.com/Narutchai01/Project_S-BE/adapters/favorite"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CommentRouters(app fiber.Router, db *gorm.DB) {
	commentRepo := adapters.NewGormCommentRepository(db)
	favoriteCommentRepo := adaptersFavorite.NewGormFavoriteRepository(db)
	commentService := usecases.NewCommentUseCase(commentRepo, favoriteCommentRepo)
	commentHandler := adapters.NewHttpCommentHandler(commentService)

	comment := app.Group("/comment")
	comment.Post("/", commentHandler.CreateComment)
	comment.Get("/:thread_id", commentHandler.GetComment)
}
