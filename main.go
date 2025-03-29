package main

import (
	"log"

	"github.com/Narutchai01/Project_S-BE/config"
	"github.com/Narutchai01/Project_S-BE/db"
	_ "github.com/Narutchai01/Project_S-BE/docs"
	"github.com/Narutchai01/Project_S-BE/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// @title			Project S API
// @version		1.0
// @description	This is a sample server for Project S.
// @host			localhost:8080
// @BasePath		/api

func main() {
	app := fiber.New()
	log.Println("Fiber initialized.")
	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
		},
	))

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	dbc, err := db.ConnectDB()

	db.Seeds(dbc)

	if err != nil {
		log.Fatalf("Could not connect to the database")
	}
	routes.Router(app, dbc)

	port := config.GetEnv("PORT")

	log.Printf("Starting the server on port %s...", port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
