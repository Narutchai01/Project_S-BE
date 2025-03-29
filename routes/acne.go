package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/acne"
	adaptersFaceProblems "github.com/Narutchai01/Project_S-BE/adapters/face_problems"
	"github.com/Narutchai01/Project_S-BE/middlewares"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AcneRouters(app fiber.Router, admin fiber.Router, db *gorm.DB) {

	faceProblemRepo := adaptersFaceProblems.NewGormFaceProblemRepository(db)
	faceProblemService := usecases.NewFaceProblemUseCase(faceProblemRepo)
	acneHandler := adapters.NewHttpAcneHandler(faceProblemService)

	acneUser := app.Group("/acne")
	acneUser.Get("/", acneHandler.GetAcnes)
	acneUser.Get("/:id", acneHandler.GetAcne)

	acneAdmin := admin.Group("/acne")
	acneAdmin.Use(middlewares.AuthorizationRequired())
	acneAdmin.Post("/", acneHandler.CreateAcne)
	acneAdmin.Delete("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(fiber.ErrNotFound))
	})
	acneAdmin.Delete("/:id", acneHandler.DeleteAcne)
	acneAdmin.Put("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(fiber.ErrNotFound))
	})
	acneAdmin.Put("/:id", acneHandler.UpdateAcne)
}
