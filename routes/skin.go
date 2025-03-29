package routes

import (
	adaptersFaceProblems "github.com/Narutchai01/Project_S-BE/adapters/face_problems"
	adapters "github.com/Narutchai01/Project_S-BE/adapters/skin"
	"github.com/Narutchai01/Project_S-BE/middlewares"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SkinRouters(app fiber.Router, admin fiber.Router, db *gorm.DB) {

	faceProblemRepo := adaptersFaceProblems.NewGormFaceProblemRepository(db)
	faceProblemService := usecases.NewFaceProblemUseCase(faceProblemRepo)
	skinHandler := adapters.NewHttpSkinHandler(faceProblemService)

	skinAdmin := admin.Group("/skin")
	skinAdmin.Use(middlewares.AuthorizationRequired())
	skinAdmin.Post("/", skinHandler.CreateSkin)
	skinAdmin.Delete("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(fiber.ErrNotFound))
	})
	skinAdmin.Delete("/:id", skinHandler.DeleteSkin)
	skinAdmin.Put("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(fiber.ErrNotFound))
	})
	skinAdmin.Put("/:id", skinHandler.UpdateSkin)

	skinUser := app.Group("/skin")
	skinUser.Get("/", skinHandler.GetSkins)
	skinUser.Get("/:id", skinHandler.GetSkin)

}
