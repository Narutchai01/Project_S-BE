package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/result"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ResultRoutes(app fiber.Router, user fiber.Router, db *gorm.DB) {

	resultRepo := adapters.NewGormResultRepository(db)
	resultService := usecases.NewResultUsecase(resultRepo)
	resultHandler := adapters.NewHttpResultHandler(resultService)

	result := app.Group("/result")
	result.Post("/", resultHandler.CreateResult)
	result.Get("/", resultHandler.GetResults)
	result.Get("/:id", resultHandler.GetResultById)
	result.Put("/:id", resultHandler.UpdateResultById)
	result.Delete("/:id", resultHandler.DeleteResultById)
	result.Get("/user/:userId", resultHandler.GetResultsByUserId)
	
	resultUser := user.Group("/result")
	resultUser.Get("/", resultHandler.GetResultsByUserIdFromToken)
}