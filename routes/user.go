package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/user"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(app fiber.Router, db *gorm.DB) {

	userRepo := adapters.NewGormUserRepository(db)
	userService := usecases.NewUserUseCase(userRepo)
	userHandler := adapters.NewHttpUserHandler(userService)

	app.Post("/register", userHandler.Register)
	

}