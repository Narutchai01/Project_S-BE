package adapters_test

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http/httptest"
	"testing"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/admin"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAdminService is a mock implementation of the usecases.AdminUsecases interface
type MockAdminService struct {
	mock.Mock
}

func (m *MockAdminService) CreateAdmin(admin entities.Admin, file multipart.FileHeader, c *fiber.Ctx) (entities.Admin, error) {
	args := m.Called(admin, file, c)
	return args.Get(0).(entities.Admin), args.Error(1)
}

func (m *MockAdminService) GetAdmins() ([]entities.Admin, error) {
	args := m.Called()
	return args.Get(0).([]entities.Admin), args.Error(1)
}

func (m *MockAdminService) GetAdmin(id int) (entities.Admin, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Admin), args.Error(1)
}

func (m *MockAdminService) UpdateAdmin(token string, admin entities.Admin, file *multipart.FileHeader, c *fiber.Ctx) (entities.Admin, error) {
	args := m.Called(token, admin, file, c) // Pass `file` directly
	return args.Get(0).(entities.Admin), args.Error(1)
}

func (m *MockAdminService) DeleteAdmin(id int) (entities.Admin, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Admin), args.Error(1)
}

func (m *MockAdminService) LogIn(email string, password string) (string, error) {
	args := m.Called(email, password)
	return args.Get(0).(string), args.Error(1)
}

func (m *MockAdminService) GetAdminByToken(token string) (entities.Admin, error) {
	args := m.Called(token)
	return args.Get(0).(entities.Admin), args.Error(1)
}

// Test
func TestCreateAdminHandler(t *testing.T) {
	setup := func() (*MockAdminService, *adapters.HttpAdminHandler, *fiber.App) {
		mockService := new(MockAdminService)
		handler := adapters.NewHttpAdminHandler(mockService)

		app := fiber.New()
		app.Post("/admin/manage", handler.CreateAdmin)

		return mockService, handler, app
	}

	expectData := entities.Admin{
		FullName: "aut",
		Email:    "aut@gmail.com",
		Password: "aut1234hashed",
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("CreateAdmin",
			mock.Anything,
			mock.AnythingOfType("multipart.FileHeader"),
			mock.Anything,
		).Return(expectData, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("full_name", expectData.FullName)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("POST", "/admin/manage", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed in body parser", func(t *testing.T) {
		_, _, app := setup()
		req := httptest.NewRequest("POST", "/admin/manage", bytes.NewBuffer([]byte("invalid body")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("failed to get image file", func(t *testing.T) {
		mockService, _, app := setup()
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("full_name", expectData.FullName)
		writer.Close()

		req := httptest.NewRequest("POST", "/admin/manage", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to create admin", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("CreateAdmin", mock.Anything, mock.AnythingOfType("multipart.FileHeader"), mock.Anything).Return(entities.Admin{}, errors.New("service error"))

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("full_name", expectData.FullName)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("POST", "/admin/manage", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetAdminsHandler(t *testing.T) {
	setup := func() (*MockAdminService, *adapters.HttpAdminHandler, *fiber.App) {
		mockService := new(MockAdminService)
		handler := adapters.NewHttpAdminHandler(mockService)

		app := fiber.New()
		app.Get("/admin/manage", handler.GetAdmins)

		return mockService, handler, app
	}

	expectData := []entities.Admin{
		{FullName: "aut", Email: "aut@gmail.com", Image: "autimage"},
		{FullName: "bee", Email: "bee@gmail.com", Image: "beeimage"},
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetAdmins").Return(expectData, nil)

		req := httptest.NewRequest("GET", "/admin/manage", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get admins", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GetAdmins").Return([]entities.Admin{}, errors.New("service error"))

		req := httptest.NewRequest("GET", "/admin/manage", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetAdminHandler(t *testing.T) {
	setup := func() (*MockAdminService, *adapters.HttpAdminHandler, *fiber.App) {
		mockService := new(MockAdminService)
		handler := adapters.NewHttpAdminHandler(mockService)

		app := fiber.New()
		app.Get("/admin/manage/:id", handler.GetAdmin)

		return mockService, handler, app
	}

	expectData := entities.Admin{
		FullName: "aut",
		Email:    "aut@gmail.com",
		Image:    "autimage",
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetAdmin", 1).Return(expectData, nil)

		req := httptest.NewRequest("GET", "/admin/manage/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get id from parameter", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("GET", "/admin/manage/error", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get admin", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GetAdmin", 1).Return(entities.Admin{}, errors.New("service error"))

		req := httptest.NewRequest("GET", "/admin/manage/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateAdminHandler(t *testing.T) {
	setup := func() (*MockAdminService, *adapters.HttpAdminHandler, *fiber.App) {
		mockService := new(MockAdminService)
		handler := adapters.NewHttpAdminHandler(mockService)

		app := fiber.New()
		app.Put("/admin/manage", handler.UpdateAdmin)

		return mockService, handler, app
	}

	expectData := entities.Admin{
		FullName: "aut",
		Email:    "aut@gmail.com",
		Password: "aut1234hashed",
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateAdmin", mock.Anything,
			mock.Anything,
			mock.AnythingOfType("*multipart.FileHeader"),
			mock.Anything,
		).Return(expectData, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("full_name", expectData.FullName)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/manage", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed in body parser", func(t *testing.T) {
		_, _, app := setup()
		req := httptest.NewRequest("PUT", "/admin/manage", bytes.NewBuffer([]byte("invalid body")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("can upload if not provide image", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateAdmin",
			mock.Anything,
			mock.Anything,
			mock.MatchedBy(func(file *multipart.FileHeader) bool {
				return file == nil
			}),
			mock.Anything,
		).Return(expectData, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("full_name", expectData.FullName) // Add form data
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/manage", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to update admin", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateAdmin",
			mock.Anything,
			mock.Anything,
			mock.AnythingOfType("*multipart.FileHeader"),
			mock.Anything,
		).Return(entities.Admin{}, errors.New("service error"))

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("full_name", expectData.FullName)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/manage", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test-token")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestDeleteAdminHandler(t *testing.T) {
	setup := func() (*MockAdminService, *adapters.HttpAdminHandler, *fiber.App) {
		mockService := new(MockAdminService)
		handler := adapters.NewHttpAdminHandler(mockService)

		app := fiber.New()
		app.Delete("/admin/manage/:id", handler.DeleteAdmin)

		return mockService, handler, app
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("DeleteAdmin", 1).Return(entities.Admin{}, nil)

		req := httptest.NewRequest("DELETE", "/admin/manage/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get id from parameter", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("DELETE", "/admin/manage/error", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to delete admin", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("DeleteAdmin", 1).Return(entities.Admin{}, errors.New("service error"))

		req := httptest.NewRequest("DELETE", "/admin/manage/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestLoginHandler(t *testing.T) {
	setup := func() (*MockAdminService, *adapters.HttpAdminHandler, *fiber.App) {
		mockService := new(MockAdminService)
		handler := adapters.NewHttpAdminHandler(mockService)

		app := fiber.New()
		app.Post("/admin/login/", handler.LogIn)

		return mockService, handler, app
	}

	expectData := entities.Admin{
		FullName: "aut",
		Email:    "aut@gmail.com",
		Image:    "autimage",
		Password: "1234",
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("LogIn",
			mock.Anything, mock.Anything,
		).Return("some token", nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("email", expectData.FullName)
		_ = writer.WriteField("password", expectData.Password)
		writer.Close()

		req := httptest.NewRequest("POST", "/admin/login/", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed in body parser", func(t *testing.T) {
		_, _, app := setup()
		req := httptest.NewRequest("POST", "/admin/login/", bytes.NewBuffer([]byte("invalid body")))
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

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("email", expectData.FullName)
		_ = writer.WriteField("password", expectData.Password)
		writer.Close()

		req := httptest.NewRequest("POST", "/admin/login/", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetAdminByTokenHandler(t *testing.T) {
	setup := func() (*MockAdminService, *adapters.HttpAdminHandler, *fiber.App) {
		mockService := new(MockAdminService)
		handler := adapters.NewHttpAdminHandler(mockService)

		app := fiber.New()
		app.Get("/admin/profile/", handler.GetAdminByToken)

		return mockService, handler, app
	}

	expectData := entities.Admin{
		FullName: "aut",
		Email:    "aut@gmail.com",
		Image:    "autimage",
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetAdminByToken", mock.Anything).Return(expectData, nil)

		req := httptest.NewRequest("GET", "/admin/profile/", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get admin", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GetAdminByToken", mock.Anything).Return(entities.Admin{}, errors.New("service error"))

		req := httptest.NewRequest("GET", "/admin/profile/", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
