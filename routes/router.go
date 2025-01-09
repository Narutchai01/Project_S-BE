package routes

import (
	_ "github.com/Narutchai01/Project_S-BE/docs"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	admin := api.Group("/admin")
	AdminRoutes(admin, db)

	SkincareRoutes(api, admin, db)

}
