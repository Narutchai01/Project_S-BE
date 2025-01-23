package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/admin"
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
	app.Put("/manage/", adminHandler.UpdateAdmin)
	app.Delete("/manage/:id", adminHandler.DeleteAdmin)
	app.Post("/login/", adminHandler.LogIn)
	app.Get("/profile/", adminHandler.GetAdminByToken)
}
