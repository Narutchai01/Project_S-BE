package adapters_test

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http/httptest"

	"testing"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/skincare"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockSkincareService struct {
	mock.Mock
}

func (m *MockSkincareService) CreateSkincare(skincare entities.Skincare, file multipart.FileHeader, token string, c *fiber.Ctx) (entities.Skincare, error) {
	args := m.Called(skincare, file, token, c)
	return args.Get(0).(entities.Skincare), args.Error(1)
}

func (m *MockSkincareService) DeleteSkincareById(id int) (entities.Skincare, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Skincare), args.Error(1)
}

func (m *MockSkincareService) GetSkincareById(id int) (entities.Skincare, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Skincare), args.Error(1)
}

func (m *MockSkincareService) GetSkincares() ([]entities.Skincare, error) {
	args := m.Called()
	return args.Get(0).([]entities.Skincare), args.Error(1)
}

func (m *MockSkincareService) UpdateSkincareById(id int, skincare entities.Skincare, file *multipart.FileHeader, c *fiber.Ctx) (entities.Skincare, error) {
	args := m.Called(id, skincare, file, c)
	return args.Get(0).(entities.Skincare), args.Error(1)
}

func TestCreateSkincareHandler(t *testing.T) {
	setup := func() (*MockSkincareService, *adapters.HttpSkincareHandler, *fiber.App) {
		mockService := new(MockSkincareService)
		handler := adapters.NewHttpSkincareHandler(mockService)

		app := fiber.New()
		app.Post("/admin/skincare", handler.CreateSkincare)

		return mockService, handler, app
	}

	expectData := entities.Skincare{
		Image:       "innisfree/image/path",
		Name:        "innisfree",
		Description: "green tea seed serum",
		CreateBY:    1,
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("CreateSkincare",
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

		req := httptest.NewRequest("POST", "/admin/skincare", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed in body parser", func(t *testing.T) {
		_, _, app := setup()
		req := httptest.NewRequest("POST", "/admin/skincare", bytes.NewBuffer([]byte("invalid body")))
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

		req := httptest.NewRequest("POST", "/admin/skincare", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to create skincare", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("CreateSkincare",
			mock.Anything,
			mock.AnythingOfType("multipart.FileHeader"),
			mock.Anything,
			mock.Anything,
		).Return(entities.Skincare{}, errors.New("service error"))

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("POST", "/admin/skincare", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetSkinsHandler(t *testing.T) {
	setup := func() (*MockSkincareService, *adapters.HttpSkincareHandler, *fiber.App) {
		mockService := new(MockSkincareService)
		handler := adapters.NewHttpSkincareHandler(mockService)

		app := fiber.New()
		app.Get("/skincare", handler.GetSkincares)

		return mockService, handler, app
	}

	expectData := []entities.Skincare{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Image:       "innisfree/image/path",
			Name:        "innisfree",
			Description: "green tea seed serum",
			CreateBY:    1,
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Image:       "innisfree2/image/path",
			Name:        "innisfree2",
			Description: "green tea seed serum",
			CreateBY:    1,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetSkincares").Return(expectData, nil)

		req := httptest.NewRequest("GET", "/skincare", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get skins", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GetSkincares").Return([]entities.Skincare{}, errors.New("service error"))

		req := httptest.NewRequest("GET", "/skincare", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetSkinHandler(t *testing.T) {
	setup := func() (*MockSkincareService, *adapters.HttpSkincareHandler, *fiber.App) {
		mockService := new(MockSkincareService)
		handler := adapters.NewHttpSkincareHandler(mockService)

		app := fiber.New()
		app.Get("/skincare/:id", handler.GetSkincareById)

		return mockService, handler, app
	}

	expectData := entities.Skincare{
		Model: gorm.Model{
			ID: 1,
		},
		Image:       "innisfree/image/path",
		Name:        "innisfree",
		Description: "green tea seed serum",
		CreateBY:    1,
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetSkincareById", int(expectData.ID)).Return(expectData, nil)

		req := httptest.NewRequest("GET", "/skincare/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to convert id to int", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("GET", "/skincare/error-id", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get skin", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GetSkincareById", int(expectData.ID)).Return(entities.Skincare{}, errors.New("service error"))

		req := httptest.NewRequest("GET", "/skincare/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateSkinHandler(t *testing.T) {
	setup := func() (*MockSkincareService, *adapters.HttpSkincareHandler, *fiber.App) {
		mockService := new(MockSkincareService)
		handler := adapters.NewHttpSkincareHandler(mockService)

		app := fiber.New()
		app.Put("/admin/skincare/:id", handler.UpdateSkincareById)

		return mockService, handler, app
	}

	expectData := entities.Skincare{
		Model: gorm.Model{
			ID: 1,
		},
		Image:       "innisfree/image/path",
		Name:        "innisfree",
		Description: "green tea seed serum",
		CreateBY:    1,
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateSkincareById",
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

		req := httptest.NewRequest("PUT", "/admin/skincare/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to convert id to int", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("PUT", "/admin/skincare/error-id", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed in body parser", func(t *testing.T) {
		_, _, app := setup()
		req := httptest.NewRequest("PUT", "/admin/skincare/1", bytes.NewBuffer([]byte("invalid body")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("can upload if not provide image", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateSkincareById",
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

		req := httptest.NewRequest("PUT", "/admin/skincare/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to update skin", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateSkincareById",
			mock.Anything,
			mock.Anything,
			mock.AnythingOfType("*multipart.FileHeader"),
			mock.Anything,
		).Return(entities.Skincare{}, errors.New("service error"))

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/skincare/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test-token")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestDeleteAdminHandler(t *testing.T) {
	setup := func() (*MockSkincareService, *adapters.HttpSkincareHandler, *fiber.App) {
		mockService := new(MockSkincareService)
		handler := adapters.NewHttpSkincareHandler(mockService)

		app := fiber.New()
		app.Delete("/admin/skincare/:id", handler.DeleteSkincareById)

		return mockService, handler, app
	}

	expectData := entities.Skincare{
		Model: gorm.Model{
			ID: 1,
		},
		Image:       "innisfree/image/path",
		Name:        "innisfree",
		Description: "green tea seed serum",
		CreateBY:    1,
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("DeleteSkincareById", int(expectData.ID)).Return(entities.Skincare{}, nil)

		req := httptest.NewRequest("DELETE", "/admin/skincare/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to convert id to int", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("DELETE", "/admin/skincare/error-id", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to delete skin", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("DeleteSkincareById", int(expectData.ID)).Return(entities.Skincare{}, errors.New("service error"))

		req := httptest.NewRequest("DELETE", "/admin/skincare/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
