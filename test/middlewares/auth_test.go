package middlewares_test

import (
	"net/http/httptest"
	"testing"

	"github.com/Narutchai01/Project_S-BE/middlewares"
	"github.com/Narutchai01/Project_S-BE/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	t.Setenv("JWT_SECRET_KEY", "test-secret-key")

	app := fiber.New()
	app.Post("/admin/skincare", middlewares.AuthorizationRequired())

	t.Run("Token is not Provided", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/admin/skincare", nil)
		res, err := app.Test(req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, fiber.StatusUnauthorized, res.StatusCode)
	})

	t.Run("Token is provided", func(t *testing.T) {
		token, _ := utils.GenerateToken(1)
		req := httptest.NewRequest("POST", "/admin/skincare", nil)
		req.Header.Set("token", "Bearer "+token)

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEqual(t, fiber.StatusUnauthorized, resp.StatusCode)
	})
}
