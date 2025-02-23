package adapter_test

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/thread"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockThreadUseCase is a mock implementation of the ThreadUseCase interface
type MockThreadUseCase struct {
	mock.Mock
}

func (m *MockThreadUseCase) GetThread(id uint, token string) (entities.Thread, error) {
	args := m.Called(id, token)
	return args.Get(0).(entities.Thread), args.Error(1)
}

// GetThreads implements usecases.ThreadUseCase.
func (m *MockThreadUseCase) GetThreads(token string) ([]entities.Thread, error) {
	args := m.Called(token)
	return args.Get(0).([]entities.Thread), args.Error(1)
}

func (m *MockThreadUseCase) CreateThread(threadDetails []entities.ThreadDetail, title string, token string, file multipart.FileHeader, c *fiber.Ctx) (entities.Thread, error) {
	args := m.Called(threadDetails, title, token, file, c)
	return args.Get(0).(entities.Thread), args.Error(1)
}

func (m *MockThreadUseCase) DeleteThread(thread_id uint) error {
	args := m.Called(thread_id)
	return args.Error(0)
}

func (m *MockThreadUseCase) UpdateThread(thread_id uint, token string, title string, thread_details []entities.ThreadDetail, file *multipart.FileHeader, c *fiber.Ctx) (entities.Thread, error) {
	args := m.Called(thread_id, token, title, thread_details, file, c)
	return args.Get(0).(entities.Thread), args.Error(1)
}

func TestCraeteThread(t *testing.T) {
	setup := func() (*MockThreadUseCase, *adapters.HttpThreadHandler, *fiber.App) {
		mockThreadUseCase := new(MockThreadUseCase)
		httpThreadHandler := adapters.NewHttpThreadHandler(mockThreadUseCase)
		app := fiber.New()

		app.Post("/thread", httpThreadHandler.CreateThread)

		return mockThreadUseCase, httpThreadHandler, app
	}
	t.Run("Success", func(t *testing.T) {
		mockThreadUseCase, _, app := setup()

		mockThreadUseCase.On("CreateThread", mock.Anything, "Test title", "test-token", mock.AnythingOfType("multipart.FileHeader"), mock.Anything).Return(entities.Thread{}, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("title", "Test title")
		writer.WriteField("thread_details", `[{"SkincareID": 1, "Caption": "test caption"}]`)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test"))
		writer.Close()

		req := httptest.NewRequest(fiber.MethodPost, "/thread", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockThreadUseCase.AssertExpectations(t)
	})

	t.Run("Missing Token", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest(fiber.MethodPost, "/thread", strings.NewReader(`{
			"ThreadDetail": "Test thread detail"
		}`))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("Missing ThreadDetail", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest(fiber.MethodPost, "/thread", strings.NewReader(`{}`))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "test-token")

		resp, _ := app.Test(req)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})
}
func TestGetThreads(t *testing.T) {
	setup := func() (*MockThreadUseCase, *adapters.HttpThreadHandler, *fiber.App) {
		mockThreadUseCase := new(MockThreadUseCase)
		httpThreadHandler := adapters.NewHttpThreadHandler(mockThreadUseCase)
		app := fiber.New()

		app.Get("/thread", httpThreadHandler.GetThreads)

		return mockThreadUseCase, httpThreadHandler, app
	}

	t.Run("Success", func(t *testing.T) {
		mockThreadUseCase, _, app := setup()

		user := entities.User{
			Model:         gorm.Model{ID: 1},
			FullName:      "09 Narutchai Mauensaen",
			Email:         "mauensaennarutchai@gmail.com",
			Birthday:      (*time.Time)(nil),
			SensitiveSkin: (*bool)(nil),
			Image:         "",
		}

		skincare := entities.Skincare{
			Model:       gorm.Model{ID: 1},
			Name:        "test skincares",
			Description: "test",
			Image:       "imageurl",
			CreateBY:    1,
		}

		thread_detail := []entities.ThreadDetail{
			{
				Model:      gorm.Model{ID: 1},
				SkincareID: 1,
				Skincare:   skincare,
				Caption:    "test 1",
			},
		}

		mockThreads := []entities.Thread{
			{
				Model:   gorm.Model{ID: 1},
				UserID:  1,
				User:    user,
				Threads: thread_detail,
			},
		}

		mockThreadUseCase.On("GetThreads", "test-token").Return(mockThreads, nil)

		req := httptest.NewRequest(fiber.MethodGet, "/thread", nil)
		req.Header.Set("token", "test-token")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockThreadUseCase.AssertExpectations(t)
	})

	t.Run("Error Without Token", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest(fiber.MethodGet, "/thread", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})
}

func TestGetThreadByID(t *testing.T) {
	setup := func() (*MockThreadUseCase, *adapters.HttpThreadHandler, *fiber.App) {
		mockThreadUseCase := new(MockThreadUseCase)
		httpThreadHandler := adapters.NewHttpThreadHandler(mockThreadUseCase)
		app := fiber.New()

		app.Get("/thread/:id", httpThreadHandler.GetThread)

		return mockThreadUseCase, httpThreadHandler, app
	}

	t.Run("Success", func(t *testing.T) {
		mockThreadUseCase, _, app := setup()

		user := entities.User{
			Model:         gorm.Model{ID: 1},
			FullName:      "09 Narutchai Mauensaen",
			Email:         "mauensaennarutchai@gmail.com",
			Birthday:      nil,
			SensitiveSkin: nil,
			Image:         "",
		}

		skincare := entities.Skincare{
			Model:       gorm.Model{ID: 1},
			Name:        "test skincares",
			Description: "test",
			Image:       "imageurl",
			CreateBY:    1,
		}

		thread_detail := []entities.ThreadDetail{
			{
				Model:      gorm.Model{ID: 1},
				SkincareID: 1,
				Skincare:   skincare,
				Caption:    "test 1",
			},
		}

		mockThread := entities.Thread{
			Model:    gorm.Model{ID: 1},
			UserID:   1,
			Bookmark: true,
			User:     user,
			Threads:  thread_detail,
		}

		mockThreadUseCase.On("GetThread", uint(1), "test-token").Return(mockThread, nil)
		req := httptest.NewRequest(fiber.MethodGet, "/thread/1", nil)
		req.Header.Set("token", "test-token")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockThreadUseCase.AssertExpectations(t)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest(fiber.MethodGet, "/thread/invalid", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Thread Not Found", func(t *testing.T) {
		mockThreadUseCase, _, app := setup()

		mockThreadUseCase.On("GetThread", uint(1), "test-token").Return(entities.Thread{}, errors.New("thread not found"))

		req := httptest.NewRequest(fiber.MethodGet, "/thread/1", nil)
		req.Header.Set("token", "test-token")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockThreadUseCase.AssertExpectations(t)
	})
}

func TestDeleteThread(t *testing.T) {
	setup := func() (*MockThreadUseCase, *adapters.HttpThreadHandler, *fiber.App) {
		mockThreadUseCase := new(MockThreadUseCase)
		httpThreadHandler := adapters.NewHttpThreadHandler(mockThreadUseCase)
		app := fiber.New()

		app.Delete("/thread/:id", httpThreadHandler.DeleteThread)

		return mockThreadUseCase, httpThreadHandler, app
	}

	t.Run("Success", func(t *testing.T) {
		mockThreadUseCase, _, app := setup()

		mockThreadUseCase.On("DeleteThread", uint(1)).Return(nil)

		req := httptest.NewRequest(fiber.MethodDelete, "/thread/1", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockThreadUseCase.AssertExpectations(t)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest(fiber.MethodDelete, "/thread/invalid", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Thread Not Found", func(t *testing.T) {
		mockThreadUseCase, _, app := setup()

		mockThreadUseCase.On("DeleteThread", uint(1)).Return(errors.New("thread not found"))

		req := httptest.NewRequest(fiber.MethodDelete, "/thread/1", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockThreadUseCase.AssertExpectations(t)
	})
}

func TestUpdateThread(t *testing.T) {
	setup := func() (*MockThreadUseCase, *adapters.HttpThreadHandler, *fiber.App) {
		mockThreadUseCase := new(MockThreadUseCase)
		httpThreadHandler := adapters.NewHttpThreadHandler(mockThreadUseCase)
		app := fiber.New()

		app.Put("/thread/:id", httpThreadHandler.UpdateThread)

		return mockThreadUseCase, httpThreadHandler, app
	}

	t.Run("Success", func(t *testing.T) {
		mockThreadUseCase, _, app := setup()

		mockThreadUseCase.On("UpdateThread", uint(1), "test-token", "Test title", mock.Anything, mock.Anything, mock.Anything).Return(entities.Thread{}, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("title", "Test title")
		writer.WriteField("thread_details", `[{"ID": 1, "SkincareID": 1, "Caption": "test caption"}]`)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test"))
		writer.Close()

		req := httptest.NewRequest(fiber.MethodPut, "/thread/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockThreadUseCase.AssertExpectations(t)
	})

	t.Run("Missing Token", func(t *testing.T) {

		mockThreadUseCase, _, app := setup()

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("title", "Test title")
		writer.WriteField("thread_details", `[{"ID": 1, "SkincareID": 1, "Caption": "test caption"}]`)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test"))
		writer.Close()

		req := httptest.NewRequest(fiber.MethodPut, "/thread/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, _ := app.Test(req)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
		mockThreadUseCase.AssertExpectations(t)
	})

	t.Run("miss thread details", func(t *testing.T) {
		mockThreadUseCase, _, app := setup()

		mockThreadUseCase.On("UpdateThread", uint(1), "test-token", "Test title", mock.Anything, mock.Anything, mock.Anything).Return(entities.Thread{}, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("title", "Test title")
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test"))
		writer.Close()

		req := httptest.NewRequest(fiber.MethodPut, "/thread/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockThreadUseCase.AssertExpectations(t)
	})

	t.Run("miss title", func(t *testing.T) {
		mockThreadUseCase, _, app := setup()

		mockThreadUseCase.On("UpdateThread", uint(1), "test-token", "", mock.Anything, mock.Anything, mock.Anything).Return(entities.Thread{}, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("title", "")
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test"))
		writer.Close()

		req := httptest.NewRequest(fiber.MethodPut, "/thread/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockThreadUseCase.AssertExpectations(t)
	})

	t.Run("miss image", func(t *testing.T) {
		mockThreadUseCase, _, app := setup()

		mockThreadUseCase.On("UpdateThread", uint(1), "test-token", "Test title", mock.Anything, mock.Anything, mock.Anything).Return(entities.Thread{}, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("title", "Test title")
		writer.WriteField("thread_details", `[{"ID": 1, "SkincareID": 1, "Caption": "test caption"}]`)
		writer.Close()

		req := httptest.NewRequest(fiber.MethodPut, "/thread/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockThreadUseCase.AssertExpectations(t)
	})
}
