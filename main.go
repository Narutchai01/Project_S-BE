package main

import (
	"log"

	"github.com/Narutchai01/Project_S-BE/config"
	"github.com/Narutchai01/Project_S-BE/db"
	"github.com/Narutchai01/Project_S-BE/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db, err := db.ConnectDB()

	if err != nil {
		log.Fatalf("Could not connect to the database")
	}
	routes.Router(app, db)

	port := config.GetEnv("PORT")

	log.Printf("Starting the server on port %s...", port)
	
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
