package adapters_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"mime/multipart"
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

func parseDate(dateStr string) *time.Time {
	if dateStr == "" {
		return nil
	}
	parsedDate, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		panic("Failed to parse date")
	}
	return &parsedDate
}

type MockUserService struct {
	mock.Mock
}

// Follower implements usecases.UserUsecases.
func (m *MockUserService) Follower(follow_id uint, token string) (entities.Follower, error) {
	args := m.Called(follow_id, token)
	return args.Get(0).(entities.Follower), args.Error(1)
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

func (m *MockUserService) GoogleSignIn(user entities.User) (string, error) {
	args := m.Called(user)
	return args.Get(0).(string), args.Error(1)
}

// Test
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
		SensitiveSkin: func(b bool) *bool { return &b }(true),
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
func TestGoogleSignInHandler(t *testing.T) {
	setup := func() (*MockUserService, *adapters.HttpUserHandler, *fiber.App) {
		mockService := new(MockUserService)
		handler := adapters.NewHttpUserHandler(mockService)

		app := fiber.New()
		app.Post("/user/google-signin", handler.GoogleSignIn)

		return mockService, handler, app
	}

	expectData := entities.User{
		Email:    "aut@gmail.com",
		FullName: "Aut",
		Image:    "dasdasdasd.png",
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GoogleSignIn",
			mock.Anything,
		).Return("some token", nil)

		body, _ := json.Marshal(expectData)

		req := httptest.NewRequest("POST", "/user/google-signin", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed in body parser", func(t *testing.T) {
		_, _, app := setup()
		req := httptest.NewRequest("POST", "/user/google-signin", bytes.NewBuffer([]byte("invalid body")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("failed to sign in with Google", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GoogleSignIn",
			mock.Anything,
		).Return("", errors.New("service error"))

		body, _ := json.Marshal(expectData)

		req := httptest.NewRequest("POST", "/user/google-signin", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func (m *MockUserService) GetUser(token string) (entities.User, error) {
	args := m.Called(token)
	return args.Get(0).(entities.User), args.Error(1)
}

func TestGetUserHandler(t *testing.T) {
	setup := func() (*MockUserService, *adapters.HttpUserHandler, *fiber.App) {
		mockService := new(MockUserService)
		handler := adapters.NewHttpUserHandler(mockService)

		app := fiber.New()
		app.Get("/user/me", handler.GetUser)

		return mockService, handler, app
	}

	expectData := entities.User{
		FullName:      "aut",
		Email:         "aut@gmail.com",
		Birthday:      parseDate("12-09-2003"),
		SensitiveSkin: func(b bool) *bool { return &b }(true),
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetUser",
			mock.Anything,
		).Return(expectData, nil)

		req := httptest.NewRequest("GET", "/user/me", nil)
		req.Header.Set("token", "some token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get user", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GetUser",
			mock.Anything,
		).Return(entities.User{}, errors.New("service error"))

		req := httptest.NewRequest("GET", "/user/me", nil)
		req.Header.Set("token", "some token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func (m *MockUserService) UpdateUser(user entities.User, token string, file *multipart.FileHeader, c *fiber.Ctx) (entities.User, error) {
	args := m.Called(user, token, file, c)
	return args.Get(0).(entities.User), args.Error(1)
}

func TestUpdateUserHandler(t *testing.T) {
	setup := func() (*MockUserService, *adapters.HttpUserHandler, *fiber.App) {
		mockService := new(MockUserService)
		handler := adapters.NewHttpUserHandler(mockService)

		app := fiber.New()
		app.Put("/user/", handler.UpdateUser)

		return mockService, handler, app
	}

	dataTest := []struct {
		Title        string
		Data         entities.User
		ExpectStatus int
	}{
		{
			Title: "Success all Param",
			Data: entities.User{
				FullName:      "aut",
				Email:         "aut@gmail.com",
				Birthday:      parseDate("12-09-2003"),
				SensitiveSkin: func(b bool) *bool { return &b }(true),
			},
			ExpectStatus: fiber.StatusOK,
		},
		{
			Title: "Success without email",
			Data: entities.User{
				FullName:      "aut",
				Birthday:      parseDate("12-09-2003"),
				SensitiveSkin: func(b bool) *bool { return &b }(true),
			},
			ExpectStatus: fiber.StatusOK,
		},
		{
			Title: "Success without fullname",
			Data: entities.User{
				Email:         "aut@gmail.com",
				Birthday:      parseDate("12-09-2003"),
				SensitiveSkin: func(b bool) *bool { return &b }(true),
			},
			ExpectStatus: fiber.StatusOK,
		},
		{
			Title: "Success without birthday",
			Data: entities.User{
				FullName:      "aut",
				Email:         "aut@gmail.com",
				SensitiveSkin: func(b bool) *bool { return &b }(true),
			},
			ExpectStatus: fiber.StatusOK,
		}, {
			Title: "Success without Sensitiv skin",
			Data: entities.User{
				FullName: "aut",
				Email:    "aut@gmail.com",
				Birthday: parseDate("12-09-2003"),
			},
			ExpectStatus: fiber.StatusOK,
		},
	}

	for _, data := range dataTest {
		data := data
		t.Run(data.Title, func(t *testing.T) {
			mockService, _, app := setup()
			mockService.On("UpdateUser",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).Return(data.Data, nil)

			body, _ := json.Marshal(data.Data)

			req := httptest.NewRequest("PUT", "/user/", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("token", "some token")
			resp, err := app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, data.ExpectStatus, resp.StatusCode)
			mockService.AssertExpectations(t)
		})
	}

}
