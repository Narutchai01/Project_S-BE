package adapters_test

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http/httptest"

	"testing"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/facial"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockFacialService struct {
	mock.Mock
}

func (m *MockFacialService) CreateFacial(facial entities.Facial, file multipart.FileHeader, c *fiber.Ctx, token string) (entities.Facial, error) {
	args := m.Called(facial, file, c, token)
	return args.Get(0).(entities.Facial), args.Error(1)
}

func (m *MockFacialService) GetFacials() ([]entities.Facial, error) {
	args := m.Called()
	return args.Get(0).([]entities.Facial), args.Error(1)
}

func (m *MockFacialService) GetFacial(id int) (entities.Facial, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Facial), args.Error(1)
}

func (m *MockFacialService) UpdateFacial(id int, facial entities.Facial, file *multipart.FileHeader, c *fiber.Ctx) (entities.Facial, error) {
	args := m.Called(id, facial, file, c)
	return args.Get(0).(entities.Facial), args.Error(1)
}

func (m *MockFacialService) DeleteFacial(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Test
func TestCreateFacialHandler(t *testing.T) {
	setup := func() (*MockFacialService, *adapters.HttpFacialHandler, *fiber.App) {
		mockService := new(MockFacialService)
		handler := adapters.NewHttpFacialHandler(mockService)

		app := fiber.New()
		app.Post("/admin/facial", handler.CreateFacial)

		return mockService, handler, app
	}

	expectData := entities.Facial{
		Name:     "facial_type1",
		Image:    "facial/type1/path",
		CreateBY: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("CreateFacial",
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

		req := httptest.NewRequest("POST", "/admin/facial", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed in body parser", func(t *testing.T) {
		_, _, app := setup()
		req := httptest.NewRequest("POST", "/admin/facial", bytes.NewBuffer([]byte("invalid body")))
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

		req := httptest.NewRequest("POST", "/admin/facial", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to create facial", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("CreateFacial",
			mock.Anything,
			mock.AnythingOfType("multipart.FileHeader"),
			mock.Anything,
			mock.Anything,
		).Return(entities.Facial{}, errors.New("service error"))

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("POST", "/admin/facial", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetFacialsHandler(t *testing.T) {
	setup := func() (*MockFacialService, *adapters.HttpFacialHandler, *fiber.App) {
		mockService := new(MockFacialService)
		handler := adapters.NewHttpFacialHandler(mockService)

		app := fiber.New()
		app.Get("/facial", handler.GetFacials)

		return mockService, handler, app
	}

	expectData := []entities.Facial{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name:     "facial_type1",
			Image:    "facial/type1/path",
			CreateBY: 1,
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name:     "facial_type2",
			Image:    "facial/type2/path",
			CreateBY: 1,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetFacials").Return(expectData, nil)

		req := httptest.NewRequest("GET", "/facial", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get facials", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GetFacials").Return([]entities.Facial{}, errors.New("service error"))

		req := httptest.NewRequest("GET", "/facial", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetFacialHandler(t *testing.T) {
	setup := func() (*MockFacialService, *adapters.HttpFacialHandler, *fiber.App) {
		mockService := new(MockFacialService)
		handler := adapters.NewHttpFacialHandler(mockService)

		app := fiber.New()
		app.Get("/facial/:id", handler.GetFacial)

		return mockService, handler, app
	}

	expectData := entities.Facial{
		Model: gorm.Model{
			ID: 1,
		},
		Name:     "facial_type1",
		Image:    "facial/type1/path",
		CreateBY: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetFacial", int(expectData.ID)).Return(expectData, nil)

		req := httptest.NewRequest("GET", "/facial/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to convert id to int", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("GET", "/facial/error-id", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get facial", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GetFacial", 1).Return(entities.Facial{}, errors.New("service error"))

		req := httptest.NewRequest("GET", "/facial/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateFacialHandler(t *testing.T) {
	setup := func() (*MockFacialService, *adapters.HttpFacialHandler, *fiber.App) {
		mockService := new(MockFacialService)
		handler := adapters.NewHttpFacialHandler(mockService)

		app := fiber.New()
		app.Put("/admin/facial/:id", handler.UpdateFacial)

		return mockService, handler, app
	}

	expectData := entities.Facial{
		Model: gorm.Model{
			ID: 1,
		},
		Name:     "facial_type1",
		Image:    "facial/type1/path",
		CreateBY: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateFacial",
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

		req := httptest.NewRequest("PUT", "/admin/facial/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to convert id to int", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("PUT", "/admin/facial/error-id", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed in body parser", func(t *testing.T) {
		_, _, app := setup()
		req := httptest.NewRequest("PUT", "/admin/facial/1", bytes.NewBuffer([]byte("invalid body")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("can upload if not provide image", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateFacial",
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

		req := httptest.NewRequest("PUT", "/admin/facial/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to update admin", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateFacial",
			mock.Anything,
			mock.Anything,
			mock.AnythingOfType("*multipart.FileHeader"),
			mock.Anything,
		).Return(entities.Facial{}, errors.New("service error"))

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/facial/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test-token")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestDeleteAdminHandler(t *testing.T) {
	setup := func() (*MockFacialService, *adapters.HttpFacialHandler, *fiber.App) {
		mockService := new(MockFacialService)
		handler := adapters.NewHttpFacialHandler(mockService)

		app := fiber.New()
		app.Delete("/admin/facial/:id", handler.DeleteFacial)

		return mockService, handler, app
	}

	expectData := entities.Facial{
		Model: gorm.Model{
			ID: 1,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("DeleteFacial", int(expectData.ID)).Return(nil)

		req := httptest.NewRequest("DELETE", "/admin/facial/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to convert id to int", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("DELETE", "/admin/facial/error-id", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to delete facial", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("DeleteFacial", int(expectData.ID)).Return(errors.New("service error"))

		req := httptest.NewRequest("DELETE", "/admin/facial/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
