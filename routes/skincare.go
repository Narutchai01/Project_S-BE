package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/skincare"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SkincareRoutes(app fiber.Router, db *gorm.DB) {

	skincareRepo := adapters.NewGormSkincareRepository(db)
	skincareService := usecases.NewSkincareUseCase(skincareRepo)
	skincareHandler := adapters.NewHttpSkincareHandler(skincareService)

	app.Get("/manage/", skincareHandler.GetSkincares)
	app.Get("/manage/:id", skincareHandler.GetSkincare)
}