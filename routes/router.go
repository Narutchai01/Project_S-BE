package routes

import (
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

	user := api.Group("/user")
	UserRoutes(user, db)

	recovery := api.Group("/recovery")
	RecoveryRoutes(recovery, db)

	SkincareRoutes(api, admin, db)

	FacialRouters(api, admin, db)
	AcneRouters(api, admin, db)
	SkinRouters(api, admin, db)
	ResultRoutes(api, db)

	ThreadRouters(api, db)
	BookMarkRouters(api, db)

}
