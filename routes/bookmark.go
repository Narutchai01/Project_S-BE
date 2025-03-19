package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/bookmark"
	adaptersCommunity "github.com/Narutchai01/Project_S-BE/adapters/community"
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
	communityRepo := adaptersCommunity.NewGormCommunityRepository(db)
	bookmarkService := usecases.NewBookmarkUseCase(bookmarkRepo, threadRepo, reviewRepo, userRepo, communityRepo)
	bookmarkHandler := adapters.NewHttpBookmarkHandler(bookmarkService)

	BookmarkGroup := app.Group("/bookmark").Use(middlewares.AuthorizationRequired())

	BookmarkGroup.Post("/thread/:id", bookmarkHandler.BookMarkThread)
	BookmarkGroup.Post("/review/:id", bookmarkHandler.BookMarkReviewSkincare)
}
