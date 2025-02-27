package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/favorite"
	adaptersUser "github.com/Narutchai01/Project_S-BE/adapters/user"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FavoriteRoutes(app fiber.Router, db *gorm.DB) {
	favoriteRepo := adapters.NewGormFavoriteRepository(db)
	userRepo := adaptersUser.NewGormUserRepository(db)
	favoriteService := usecases.NewFavoriteUseCase(favoriteRepo, userRepo)
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
