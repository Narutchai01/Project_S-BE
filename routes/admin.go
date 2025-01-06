package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/admin"
	adapters_skincare "github.com/Narutchai01/Project_S-BE/adapters/skincare"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AdminRoutes(app fiber.Router, db *gorm.DB) {

	adminRepo := adapters.NewGormAdminRepository(db)
	adminService := usecases.NewAdminUseCase(adminRepo)
	adminHandler := adapters.NewHttpAdminHandler(adminService)

	app.Get("/manage/", adminHandler.GetAdmins)
	app.Post("/manage/", adminHandler.CreateAdmin)
	app.Get("/manage/:id", adminHandler.GetAdmin)
	app.Put("/manage/:id", adminHandler.UpdateAdmin)
	app.Delete("/manage/:id", adminHandler.DeleteAdmin)
	app.Post("/login/", adminHandler.LogIn)

	skincareRepo := adapters_skincare.NewGormSkincareRepository(db)
	skincareService := usecases.NewSkincareUseCase(skincareRepo)
	skincareHandler := adapters_skincare.NewHttpSkincareHandler(skincareService)

	skincare := app.Group("/manage/skincare")
	skincare.Post("/", skincareHandler.CreateSkincare)
	skincare.Get("/", skincareHandler.GetSkincares)
	skincare.Get("/:id", skincareHandler.GetSkincareById)
	skincare.Put("/:id", skincareHandler.UpdateSkincareById)
	skincare.Delete("/:id", skincareHandler.DeleteSkincareById)

}
