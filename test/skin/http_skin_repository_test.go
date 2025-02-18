package adapters_test

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http/httptest"

	"testing"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/skin"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockSkinService struct {
	mock.Mock
}

func (m *MockSkinService) CreateSkin(skin entities.Skin, file multipart.FileHeader, c *fiber.Ctx, token string) (entities.Skin, error) {
	args := m.Called(skin, file, c, token)
	return args.Get(0).(entities.Skin), args.Error(1)
}

func (m *MockSkinService) GetSkins() ([]entities.Skin, error) {
	args := m.Called()
	return args.Get(0).([]entities.Skin), args.Error(1)
}

func (m *MockSkinService) GetSkin(id int) (entities.Skin, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Skin), args.Error(1)
}

func (m *MockSkinService) UpdateSkin(id int, skin entities.Skin, file *multipart.FileHeader, c *fiber.Ctx) (entities.Skin, error) {
	args := m.Called(id, skin, file, c)
	return args.Get(0).(entities.Skin), args.Error(1)
}

func (m *MockSkinService) DeleteSkin(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Test
func TestCreateSkinHandler(t *testing.T) {
	setup := func() (*MockSkinService, *adapters.HttpSkinHandler, *fiber.App) {
		mockService := new(MockSkinService)
		handler := adapters.NewHttpSkinHandler(mockService)

		app := fiber.New()
		app.Post("/admin/skin", handler.CreateSkin)

		return mockService, handler, app
	}

	expectData := entities.Skin{
		Name:     "skintype1",
		Image:    "skin/type1/path",
		CreateBY: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("CreateSkin",
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

		req := httptest.NewRequest("POST", "/admin/skin", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed in body parser", func(t *testing.T) {
		_, _, app := setup()
		req := httptest.NewRequest("POST", "/admin/skin", bytes.NewBuffer([]byte("invalid body")))
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

		req := httptest.NewRequest("POST", "/admin/skin", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to create skin", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("CreateSkin",
			mock.Anything,
			mock.AnythingOfType("multipart.FileHeader"),
			mock.Anything,
			mock.Anything,
		).Return(entities.Skin{}, errors.New("service error"))

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("POST", "/admin/skin", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetSkinsHandler(t *testing.T) {
	setup := func() (*MockSkinService, *adapters.HttpSkinHandler, *fiber.App) {
		mockService := new(MockSkinService)
		handler := adapters.NewHttpSkinHandler(mockService)

		app := fiber.New()
		app.Get("/skin", handler.GetSkins)

		return mockService, handler, app
	}

	expectData := []entities.Skin{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name:     "skintype1",
			Image:    "skin/type1/path",
			CreateBY: 1,
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name:     "skintype2",
			Image:    "skin/type2/path",
			CreateBY: 1,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetSkins").Return(expectData, nil)

		req := httptest.NewRequest("GET", "/skin", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get skins", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GetSkins").Return([]entities.Skin{}, errors.New("service error"))

		req := httptest.NewRequest("GET", "/skin", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetSkinHandler(t *testing.T) {
	setup := func() (*MockSkinService, *adapters.HttpSkinHandler, *fiber.App) {
		mockService := new(MockSkinService)
		handler := adapters.NewHttpSkinHandler(mockService)

		app := fiber.New()
		app.Get("/skin/:id", handler.GetSkin)

		return mockService, handler, app
	}

	expectData := entities.Skin{
		Model: gorm.Model{
			ID: 1,
		},
		Name:     "skintype1",
		Image:    "skin/type1/path",
		CreateBY: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetSkin", int(expectData.ID)).Return(expectData, nil)

		req := httptest.NewRequest("GET", "/skin/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to convert id to int", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("GET", "/skin/error-id", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get skin", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GetSkin", int(expectData.ID)).Return(entities.Skin{}, errors.New("service error"))

		req := httptest.NewRequest("GET", "/skin/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateSkinHandler(t *testing.T) {
	setup := func() (*MockSkinService, *adapters.HttpSkinHandler, *fiber.App) {
		mockService := new(MockSkinService)
		handler := adapters.NewHttpSkinHandler(mockService)

		app := fiber.New()
		app.Put("/admin/skin/:id", handler.UpdateSkin)

		return mockService, handler, app
	}

	expectData := entities.Skin{
		Model: gorm.Model{
			ID: 1,
		},
		Name:     "skintype1",
		Image:    "skin/type1/path",
		CreateBY: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateSkin",
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

		req := httptest.NewRequest("PUT", "/admin/skin/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to convert id to int", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("PUT", "/admin/skin/error-id", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed in body parser", func(t *testing.T) {
		_, _, app := setup()
		req := httptest.NewRequest("PUT", "/admin/skin/1", bytes.NewBuffer([]byte("invalid body")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("can upload if not provide image", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateSkin",
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

		req := httptest.NewRequest("PUT", "/admin/skin/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to update skin", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateSkin",
			mock.Anything,
			mock.Anything,
			mock.AnythingOfType("*multipart.FileHeader"),
			mock.Anything,
		).Return(entities.Skin{}, errors.New("service error"))

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/skin/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test-token")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestDeleteAdminHandler(t *testing.T) {
	setup := func() (*MockSkinService, *adapters.HttpSkinHandler, *fiber.App) {
		mockService := new(MockSkinService)
		handler := adapters.NewHttpSkinHandler(mockService)

		app := fiber.New()
		app.Delete("/admin/skin/:id", handler.DeleteSkin)

		return mockService, handler, app
	}

	expectData := entities.Skin{
		Model: gorm.Model{
			ID: 1,
		},
		Name:     "skintype1",
		Image:    "skin/type1/path",
		CreateBY: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("DeleteSkin", int(expectData.ID)).Return(nil)

		req := httptest.NewRequest("DELETE", "/admin/skin/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to convert id to int", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("DELETE", "/admin/skin/error-id", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to delete skin", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("DeleteSkin", int(expectData.ID)).Return(errors.New("service error"))

		req := httptest.NewRequest("DELETE", "/admin/skin/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
