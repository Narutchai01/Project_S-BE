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
		log.Fatal("Could not connect to the database")
	}

	routes.Router(app, db)

	port := config.GetEnv("PORT")

	log.Fatal(app.Listen(":" + port))
}
