package routes

import (
	adaptersBookmark "github.com/Narutchai01/Project_S-BE/adapters/bookmark"
	adaptersCommunity "github.com/Narutchai01/Project_S-BE/adapters/community"
	adaptersFavorite "github.com/Narutchai01/Project_S-BE/adapters/favorite"
	adaptersReview "github.com/Narutchai01/Project_S-BE/adapters/review"
	adaptersUser "github.com/Narutchai01/Project_S-BE/adapters/user"
	"github.com/Narutchai01/Project_S-BE/middlewares"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ReviewRoutes(app fiber.Router, db *gorm.DB) {
	userRepo := adaptersUser.NewGormUserRepository(db)
	favoriteRepo := adaptersFavorite.NewGormFavoriteRepository(db)
	bookmarkRepo := adaptersBookmark.NewGormBookmarkRepository(db)
	communityRepo := adaptersCommunity.NewGormCommunityRepository(db)
	communityService := usecases.NewCommunityUseCase(communityRepo, userRepo, favoriteRepo, bookmarkRepo)
	reviewHandler := adaptersReview.NewHttpReviewRepository(communityService)

	reviewGroup := app.Group("/reviews").Use(middlewares.AuthorizationRequired())
	reviewGroup.Post("/", reviewHandler.CreateReviewSkincare)
	reviewGroup.Get("/", reviewHandler.GetReviewSkincares)
	reviewGroup.Get("/user/:id", reviewHandler.GetReviewSkincareByUserID)
	reviewGroup.Get("/:id", reviewHandler.GetReviewSkincare)

}
