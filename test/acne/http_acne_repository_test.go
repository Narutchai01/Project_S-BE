package adapters_test

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http/httptest"

	"testing"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/acne"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockAcneService struct {
	mock.Mock
}

// CreateAcne implements usecases.AcneUseCase.
func (m *MockAcneService) CreateAcne(acne entities.Acne, file multipart.FileHeader, c *fiber.Ctx, token string) (entities.Acne, error) {
	args := m.Called(acne, file, c, token)
	return args.Get(0).(entities.Acne), args.Error(1)
}

func (m *MockAcneService) GetAcnes() ([]entities.Acne, error) {
	args := m.Called()
	return args.Get(0).([]entities.Acne), args.Error(1)
}

func (m *MockAcneService) GetAcne(id int) (entities.Acne, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Acne), args.Error(1)
}

func (m *MockAcneService) UpdateAcne(id int, acne entities.Acne, file *multipart.FileHeader, c *fiber.Ctx) (entities.Acne, error) {
	args := m.Called(id, acne, file, c)
	return args.Get(0).(entities.Acne), args.Error(1)
}

func (m *MockAcneService) DeleteAcne(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Test
func TestCreateAcneHandler(t *testing.T) {
	setup := func() (*MockAcneService, *adapters.HttpAcneHandler, *fiber.App) {
		mockService := new(MockAcneService)
		handler := adapters.NewHttpAcneHandler(mockService)

		app := fiber.New()
		app.Post("/admin/acne", handler.CreateAcne)

		return mockService, handler, app
	}

	expectData := entities.Acne{
		Name:     "innisfree",
		Image:    "innisfree/image/path",
		CreateBY: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("CreateAcne",
			mock.Anything,
			mock.AnythingOfType("multipart.FileHeader"),
			mock.Anything,
			mock.Anything,
		).Return(expectData, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("POST", "/admin/acne", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed in body parser", func(t *testing.T) {
		_, _, app := setup()
		req := httptest.NewRequest("POST", "/admin/acne", bytes.NewBuffer([]byte("invalid body")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("failed to get image file", func(t *testing.T) {
		mockService, _, app := setup()
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		writer.Close()

		req := httptest.NewRequest("POST", "/admin/acne", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to create acne", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("CreateAcne",
			mock.Anything,
			mock.AnythingOfType("multipart.FileHeader"),
			mock.Anything,
			mock.Anything,
		).Return(entities.Acne{}, errors.New("service error"))

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("POST", "/admin/acne", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetAcnesHandler(t *testing.T) {
	setup := func() (*MockAcneService, *adapters.HttpAcneHandler, *fiber.App) {
		mockService := new(MockAcneService)
		handler := adapters.NewHttpAcneHandler(mockService)

		app := fiber.New()
		app.Get("/acne", handler.GetAcnes)

		return mockService, handler, app
	}

	expectData := []entities.Acne{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name:     "innisfree",
			Image:    "innisfree/image/path",
			CreateBY: 1,
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name:     "Dr.Pong",
			Image:    "drpong/image/path",
			CreateBY: 1,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetAcnes").Return(expectData, nil)

		req := httptest.NewRequest("GET", "/acne", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get admins", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GetAcnes").Return([]entities.Acne{}, errors.New("service error"))

		req := httptest.NewRequest("GET", "/acne", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetAcneHandler(t *testing.T) {
	setup := func() (*MockAcneService, *adapters.HttpAcneHandler, *fiber.App) {
		mockService := new(MockAcneService)
		handler := adapters.NewHttpAcneHandler(mockService)

		app := fiber.New()
		app.Get("/acne/:id", handler.GetAcne)

		return mockService, handler, app
	}

	expectData := entities.Acne{
		Model: gorm.Model{
			ID: 1,
		},
		Name:     "innisfree",
		Image:    "innisfree/image/path",
		CreateBY: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetAcne", int(expectData.ID)).Return(expectData, nil)

		req := httptest.NewRequest("GET", "/acne/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to convert id to int", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("GET", "/acne/error-id", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get acne", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GetAcne", int(expectData.ID)).Return(entities.Acne{}, errors.New("service error"))

		req := httptest.NewRequest("GET", "/acne/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateAcneHandler(t *testing.T) {
	setup := func() (*MockAcneService, *adapters.HttpAcneHandler, *fiber.App) {
		mockService := new(MockAcneService)
		handler := adapters.NewHttpAcneHandler(mockService)

		app := fiber.New()
		app.Put("/admin/acne/:id", handler.UpdateAcne)

		return mockService, handler, app
	}

	expectData := entities.Acne{
		Model: gorm.Model{
			ID: 1,
		},
		Name:     "innisfree",
		Image:    "innisfree/image/path",
		CreateBY: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateAcne",
			mock.Anything,
			mock.Anything,
			mock.AnythingOfType("*multipart.FileHeader"),
			mock.Anything,
		).Return(expectData, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/acne/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to convert id to int", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("PUT", "/admin/acne/error-id", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed in body parser", func(t *testing.T) {
		_, _, app := setup()
		req := httptest.NewRequest("PUT", "/admin/acne/1", bytes.NewBuffer([]byte("invalid body")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("can upload if not provide image", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateAcne",
			mock.Anything,
			mock.Anything,
			mock.MatchedBy(func(file *multipart.FileHeader) bool {
				return file == nil
			}),
			mock.Anything,
		).Return(expectData, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/acne/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to update admin", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateAcne",
			mock.Anything,
			mock.Anything,
			mock.AnythingOfType("*multipart.FileHeader"),
			mock.Anything,
		).Return(entities.Acne{}, errors.New("service error"))

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/acne/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test-token")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestDeleteAdminHandler(t *testing.T) {
	setup := func() (*MockAcneService, *adapters.HttpAcneHandler, *fiber.App) {
		mockService := new(MockAcneService)
		handler := adapters.NewHttpAcneHandler(mockService)

		app := fiber.New()
		app.Delete("/admin/acne/:id", handler.DeleteAcne)

		return mockService, handler, app
	}

	expectData := entities.Acne{
		Model: gorm.Model{
			ID: 1,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("DeleteAcne", int(expectData.ID)).Return(nil)

		req := httptest.NewRequest("DELETE", "/admin/acne/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to convert id to int", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("DELETE", "/admin/acne/error-id", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to delete acne", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("DeleteAcne", int(expectData.ID)).Return(errors.New("service error"))

		req := httptest.NewRequest("DELETE", "/admin/acne/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
