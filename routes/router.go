package routes

import (
	_ "github.com/Narutchai01/Project_S-BE/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

//	@title			Project S API
//	@version		1.0
//	@description	This is a sample server for Project S.
//	@host			localhost:8080
//	@BasePath		/api

func Router(app *fiber.App, db *gorm.DB) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	admin := api.Group("/admin")
	AdminRoutes(admin, db)

	SkincareRoutes(api, admin, db)

}
