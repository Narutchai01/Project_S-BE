package middlewares

import (
	// "fmt"

	"github.com/Narutchai01/Project_S-BE/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func AuthorizationRequired() fiber.Handler {
	secretKey := config.GetEnv("JWT_SECRET_KEY")
	return jwtware.New(jwtware.Config{
		SuccessHandler: AuthSuccess,
		ErrorHandler:   AuthError,
		SigningKey:     []byte(secretKey),
		SigningMethod:  "HS256",
		TokenLookup:    "header:token", 
		AuthScheme:     "",         
	})
}

  func AuthError(c *fiber.Ctx, e error) error {
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	    "message":   e.Error(),
	})
	return nil
  }
  
  func AuthSuccess(c *fiber.Ctx) error {
	c.Next()
	return nil
  }