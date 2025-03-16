package routes

import (
	adaptersBookmark "github.com/Narutchai01/Project_S-BE/adapters/bookmark"
	adaptersCommunity "github.com/Narutchai01/Project_S-BE/adapters/community"
	adaptersFavorite "github.com/Narutchai01/Project_S-BE/adapters/favorite"
	adaptersReview "github.com/Narutchai01/Project_S-BE/adapters/review"
	adaptersSkincare "github.com/Narutchai01/Project_S-BE/adapters/skincare"
	adaptersUser "github.com/Narutchai01/Project_S-BE/adapters/user"
	"github.com/Narutchai01/Project_S-BE/middlewares"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ReviewRoutes(app fiber.Router, db *gorm.DB) {
	reviewRepo := adaptersReview.NewGormReviewRepository(db)
	userRepo := adaptersUser.NewGormUserRepository(db)
	skincareRepo := adaptersSkincare.NewGormSkincareRepository(db)
	favoriteRepo := adaptersFavorite.NewGormFavoriteRepository(db)
	bookmarkRepo := adaptersBookmark.NewGormBookmarkRepository(db)
	communityRepo := adaptersCommunity.NewGormCommunityRepository(db)
	reviewService := usecases.NewReviewUseCase(reviewRepo, userRepo, skincareRepo, favoriteRepo, bookmarkRepo)
	communityService := usecases.NewCommunityUseCase(communityRepo, userRepo)
	reviewHandler := adaptersReview.NewHttpReviewRepository(reviewService, communityService)

	reviewGroup := app.Group("/reviews").Use(middlewares.AuthorizationRequired())
	reviewGroup.Post("/", reviewHandler.CreateReviewSkincare)
	reviewGroup.Get("/", reviewHandler.GetReviewSkincares)
	reviewGroup.Get("/:id", reviewHandler.GetReviewSkincare)

}
