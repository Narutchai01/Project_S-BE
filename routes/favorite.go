package routes

import (
	adaptersComment "github.com/Narutchai01/Project_S-BE/adapters/comment"
	adaptersCommunity "github.com/Narutchai01/Project_S-BE/adapters/community"
	adapters "github.com/Narutchai01/Project_S-BE/adapters/favorite"
	adaptersUser "github.com/Narutchai01/Project_S-BE/adapters/user"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FavoriteRoutes(app fiber.Router, db *gorm.DB) {
	favoriteRepo := adapters.NewGormFavoriteRepository(db)
	userRepo := adaptersUser.NewGormUserRepository(db)
	commentRepo := adaptersComment.NewGormCommentRepository(db)
	communityRepo := adaptersCommunity.NewGormCommunityRepository(db)
	favoriteService := usecases.NewFavoriteUseCase(favoriteRepo, userRepo, commentRepo, communityRepo)
	favoriteHandler := adapters.NewHttpFavoriteHandler(favoriteService)

	favorite := app.Group("/favorite")
	favorite.Post("/thread/:id", favoriteHandler.HandleFavoriteThread)

	reviewGroup := favorite.Group("/review")
	reviewGroup.Post("/skincare/:id", favoriteHandler.HandleFavoriteReviewSkincare)

	commentGroup := favorite.Group("/comment")
	commentGroup.Post("/thread/:id", favoriteHandler.HandleFavoriteCommentThread)

	reviewCommentGroup := commentGroup.Group("/review")
	reviewCommentGroup.Post("/skincare/:id", favoriteHandler.HandleFavoriteCommentReviewSkincare)

}
