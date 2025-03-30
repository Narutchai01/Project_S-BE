package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/user"
	"github.com/Narutchai01/Project_S-BE/middlewares"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(app fiber.Router, db *gorm.DB) {

	userRepo := adapters.NewGormUserRepository(db)
	userService := usecases.NewUserUseCase(userRepo)
	userHandler := adapters.NewHttpUserHandler(userService)

	app.Post("/register", userHandler.Register)
	app.Post("/login/", userHandler.LogIn)
	app.Put("/forget-password", userHandler.ForgetPassword)
	app.Post("/goolge-signin", userHandler.GoogleSignIn)
	app.Get("/me", middlewares.AuthorizationRequired(), userHandler.GetUser)
	app.Post("/follower/:follow_id", middlewares.AuthorizationRequired(), userHandler.Follower)
	app.Put("/", middlewares.AuthorizationRequired(), userHandler.UpdateUser)
	app.Get("/:id", middlewares.AuthorizationRequired(), userHandler.GetUserByID)
}
