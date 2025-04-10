package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/result"
	adaptersUser "github.com/Narutchai01/Project_S-BE/adapters/user"
	"github.com/Narutchai01/Project_S-BE/middlewares"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ResultRoutes(app fiber.Router, db *gorm.DB) {
	resultRepo := adapters.NewGormResultRepository(db)
	userRepo := adaptersUser.NewGormUserRepository(db)
	resultService := usecases.NewResultsUsecase(resultRepo, userRepo)
	resultHandler := adapters.NewHttpResultHandler(resultService)

	resultGroup := app.Group("/results")
	resultGroup.Use(middlewares.AuthorizationRequired())
	resultGroup.Post("/", resultHandler.CreateResult)
	resultGroup.Get("/", resultHandler.GetResults)
	resultGroup.Get("/latest", resultHandler.GetResultLatest)
	resultGroup.Post("/compare", resultHandler.GetResultByIDs)
	resultGroup.Get("/:id", resultHandler.GetResult)
	// resultGroup.Put("/:id", resultHandler.UpdateResult)
	// resultGroup.Delete("/:id", resultHandler.DeleteResult)
}
