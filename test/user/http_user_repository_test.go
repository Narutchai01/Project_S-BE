package adapters_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"

	"testing"
	"time"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/user"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func parseDate(dateStr string) time.Time {
	parsedDate, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		panic("Failed to parse date")
	}
	return parsedDate
}

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Register(user entities.User, c *fiber.Ctx) (entities.User, error) {
	args := m.Called(user, c)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockUserService) LogIn(email string, password string) (string, error) {
	args := m.Called(email, password)
	return args.Get(0).(string), args.Error(1)
}

func (m *MockUserService) ChangePassword(id int, ewPassword string, c *fiber.Ctx) (entities.User, error) {
	args := m.Called(id, ewPassword, c)
	return args.Get(0).(entities.User), args.Error(1)
}

//Test
func TestRegisterHandler(t *testing.T) {
	setup := func() (*MockUserService, *adapters.HttpUserHandler, *fiber.App) {
		mockService := new(MockUserService)
		handler := adapters.NewHttpUserHandler(mockService)

		app := fiber.New()
		app.Post("/user/register", handler.Register)

		return mockService, handler, app
	}

	expectData := entities.User{
		FullName:      "aut",
		Email:         "aut@gmail.com",
		Birthday:      parseDate("12-09-2003"),
		SensitiveSkin: true,
		Password:      "aut1234hashed",
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("Register",
			mock.Anything,
			mock.Anything,
		).Return(expectData, nil)

		body, _ := json.Marshal(expectData)

		req := httptest.NewRequest("POST", "/user/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed in body parser", func(t *testing.T) {
		_, _, app := setup()
		req := httptest.NewRequest("POST", "/user/register", bytes.NewBuffer([]byte("invalid body")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("failed to create admin", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("Register", 
			mock.Anything,
			mock.Anything,
		).Return(entities.User{}, errors.New("service error"))

		body, _ := json.Marshal(expectData)

		req := httptest.NewRequest("POST", "/user/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestLoginHandler(t *testing.T) {
	setup := func() (*MockUserService, *adapters.HttpUserHandler, *fiber.App) {
		mockService := new(MockUserService)
		handler := adapters.NewHttpUserHandler(mockService)

		app := fiber.New()
		app.Post("/user/login/", handler.LogIn)

		return mockService, handler, app
	}

	expectData := entities.User{
		Email:    "aut@gmail.com",
		Password: "1234",
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("LogIn",
			mock.Anything, 
			mock.Anything,
		).Return("some token", nil)

		body, _ := json.Marshal(expectData)

		req := httptest.NewRequest("POST", "/user/login/", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed in body parser", func(t *testing.T) {
		_, _, app := setup()
		req := httptest.NewRequest("POST", "/user/login", bytes.NewBuffer([]byte("invalid body")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("failed to login", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("LogIn",
			mock.Anything,
			mock.Anything,
		).Return("", errors.New("service error"))

		body, _ := json.Marshal(expectData)

		req := httptest.NewRequest("POST", "/user/login/", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestForgetPasswordHandler(t *testing.T) {
	setup := func() (*MockUserService, *adapters.HttpUserHandler, *fiber.App) {
		mockService := new(MockUserService)
		handler := adapters.NewHttpUserHandler(mockService)

		app := fiber.New()
		app.Put("/user/forget-password", handler.ForgetPassword)

		return mockService, handler, app
	}

	expectData := entities.User{
		Model: gorm.Model{
			ID: 1,
		},
		Password: "1234",
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("ChangePassword",
			mock.Anything, 
			mock.Anything,
			mock.Anything,
		).Return(expectData, nil)

		body, _ := json.Marshal(expectData)

		req := httptest.NewRequest("PUT", "/user/forget-password", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed in body parser", func(t *testing.T) {
		_, _, app := setup()
		req := httptest.NewRequest("PUT", "/user/forget-password", bytes.NewBuffer([]byte("invalid body")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("failed to change password", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("ChangePassword",
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(expectData, errors.New("service error"))

		body, _ := json.Marshal(expectData)

		req := httptest.NewRequest("PUT", "/user/forget-password", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}