package config

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		if testing.Testing() {
			log.Println("Skipping .env loading in tests")
			return ""
		}
		log.Fatalf("Error loading .env file for key: %s", key)
	}
	return os.Getenv(key)
}
