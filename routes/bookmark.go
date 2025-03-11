package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/bookmark"
	adaptersReview "github.com/Narutchai01/Project_S-BE/adapters/review"
	adaptersThread "github.com/Narutchai01/Project_S-BE/adapters/thread"
	adaptersUser "github.com/Narutchai01/Project_S-BE/adapters/user"
	"github.com/Narutchai01/Project_S-BE/middlewares"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func BookMarkRouters(app fiber.Router, db *gorm.DB) {

	bookmarkRepo := adapters.NewGormBookmarkRepository(db)
	threadRepo := adaptersThread.NewGormThreadRepository(db)
	reviewRepo := adaptersReview.NewGormReviewRepository(db)
	userRepo := adaptersUser.NewGormUserRepository(db)
	bookmarkService := usecases.NewBookmarkUseCase(bookmarkRepo, threadRepo, reviewRepo, userRepo)
	bookmarkHandler := adapters.NewHttpBookmarkHandler(bookmarkService)

	BookmarkGroup := app.Group("/bookmark").Use(middlewares.AuthorizationRequired())

	BookmarkGroup.Post("/thread/:id", bookmarkHandler.BookMarkThread)
	BookmarkGroup.Post("/review/:id", bookmarkHandler.BookMarkReviewSkincare)
}
