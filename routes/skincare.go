package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/skincare"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SkincareRoutes(app fiber.Router, admin fiber.Router, db *gorm.DB) {

	skincareRepo := adapters.NewGormSkincareRepository(db)
	skincareService := usecases.NewSkincareUseCase(skincareRepo)
	skincareHandler := adapters.NewHttpSkincareHandler(skincareService)

	app.Get("/skincare", skincareHandler.GetSkincares)
	app.Get("/skincare/:id", skincareHandler.GetSkincareById)

	//admin
	admin.Post("/skincare", skincareHandler.CreateSkincare)
	admin.Put("/skincare/:id", skincareHandler.UpdateSkincareById)
	admin.Delete("/skincare/:id", skincareHandler.DeleteSkincareById)
}