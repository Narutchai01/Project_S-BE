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
	log.Println("Starting the application...")

	app := fiber.New()
	log.Println("Fiber initialized.")
	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
		},
	))

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	db, err := db.ConnectDB()
	log.Println("Connecting to the database...")
	if err != nil {
		log.Fatalf("Could not connect to the database")
	}
	log.Println("Connected to the database successfully.")

	log.Println("Setting up routes...")
	routes.Router(app, db)
	log.Println("Routes have been set up.")

	port := config.GetEnv("PORT")
	if port == "" {
		log.Fatal("Environment variable PORT is not set")
	}

	log.Printf("Starting the server on port %s...", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
