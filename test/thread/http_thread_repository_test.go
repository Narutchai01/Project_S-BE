package adapter_test

import (
	"errors"
	"net/http/httptest"
	"strings"
	"testing"

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

func (m *MockThreadUseCase) GetThread(id uint) (entities.Thread, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Thread), args.Error(1)
}

// GetThreads implements usecases.ThreadUseCase.
func (m *MockThreadUseCase) GetThreads() ([]entities.Thread, error) {
	args := m.Called()
	return args.Get(0).([]entities.Thread), args.Error(1)
}

func (m *MockThreadUseCase) CreateThread(thread entities.ThreadRequest, token string) (entities.Thread, error) {
	args := m.Called(thread, token)
	return args.Get(0).(entities.Thread), args.Error(1)
}

func (m *MockThreadUseCase) DeleteThread(thread_id uint) error {
	args := m.Called(thread_id)
	return args.Error(0)
}

func (m *MockThreadUseCase) AddBookmark(thread_id uint, token string) (entities.Bookmark, error) {
	args := m.Called(thread_id, token)
	return args.Get(0).(entities.Bookmark), args.Error(1)
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

		mockThreadUseCase.On("CreateThread", mock.Anything, mock.Anything).Return(entities.Thread{}, nil)

		req := httptest.NewRequest(fiber.MethodPost, "/thread", strings.NewReader(`{
			"thread_detail": [
				{
					"skincare_id": 4,
					"caption": "caption"
				},
				{
					"skincare_id": 3,
					"caption": "caption"
				}
			]
		}`))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "test-token")

		resp, _ := app.Test(req)
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
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Missing ThreadDetail", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest(fiber.MethodPost, "/thread", strings.NewReader(`{}`))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "test-token")

		resp, _ := app.Test(req)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
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

		mockThreads := []entities.Thread{
			{
				Model:   gorm.Model{ID: 1},
				UserID:  1,
				User:    user,
				Threads: thread_detail,
			},
		}

		mockThreadUseCase.On("GetThreads").Return(mockThreads, nil)

		req := httptest.NewRequest(fiber.MethodGet, "/thread", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockThreadUseCase.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockThreadUseCase, _, app := setup()

		mockThreadUseCase.On("GetThreads").Return([]entities.Thread{}, errors.New("something went wrong"))

		req := httptest.NewRequest(fiber.MethodGet, "/thread", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockThreadUseCase.AssertExpectations(t)
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
			Model:   gorm.Model{ID: 1},
			UserID:  1,
			User:    user,
			Threads: thread_detail,
		}

		mockThreadUseCase.On("GetThread", uint(1)).Return(mockThread, nil)

		req := httptest.NewRequest(fiber.MethodGet, "/thread/1", nil)
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

		mockThreadUseCase.On("GetThread", uint(1)).Return(entities.Thread{}, errors.New("thread not found"))

		req := httptest.NewRequest(fiber.MethodGet, "/thread/1", nil)
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
func TestBookmarkThread(t *testing.T) {
	setup := func() (*MockThreadUseCase, *adapters.HttpThreadHandler, *fiber.App) {
		mockThreadUseCase := new(MockThreadUseCase)
		httpThreadHandler := adapters.NewHttpThreadHandler(mockThreadUseCase)
		app := fiber.New()

		app.Post("/thread/:id/bookmark", httpThreadHandler.BookMark)

		return mockThreadUseCase, httpThreadHandler, app
	}

	t.Run("Success", func(t *testing.T) {
		mockThreadUseCase, _, app := setup()

		mockBookmark := entities.Bookmark{
			ThreadID: 1,
			UserID:   1,
		}

		mockThreadUseCase.On("AddBookmark", uint(1), "test-token").Return(mockBookmark, nil)

		req := httptest.NewRequest(fiber.MethodPost, "/thread/1/bookmark", nil)
		req.Header.Set("token", "test-token")

		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockThreadUseCase.AssertExpectations(t)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest(fiber.MethodPost, "/thread/invalid/bookmark", nil)
		req.Header.Set("token", "test-token")

		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Missing Token", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest(fiber.MethodPost, "/thread/1/bookmark", nil)

		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Error Adding Bookmark", func(t *testing.T) {
		mockThreadUseCase, _, app := setup()

		mockThreadUseCase.On("AddBookmark", uint(1), "test-token").Return(entities.Bookmark{}, errors.New("invalid thread ID"))

		req := httptest.NewRequest(fiber.MethodPost, "/thread/1/bookmark", nil)
		req.Header.Set("token", "test-token")

		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockThreadUseCase.AssertExpectations(t)
	})
}
