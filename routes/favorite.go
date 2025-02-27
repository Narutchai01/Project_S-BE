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
	favorite.Post("/comment/:id", favoriteHandler.HandleFavoriteComment)
	favorite.Post("/thread/:id", favoriteHandler.HandleFavoriteThread)
	favorite.Post("/reviewskincare/:id", favoriteHandler.HandleFavoriteReviewSkincare)

}
