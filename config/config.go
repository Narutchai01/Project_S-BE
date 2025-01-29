package config

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if testing.Testing() { 
		log.Println("Skipping .env loading in tests")
		return
	}

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system ENV")
	}
}


func GetEnv(key string) string {
	return os.Getenv(key)
}