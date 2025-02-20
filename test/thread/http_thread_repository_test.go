package adapter_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/thread"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockThreadUseCase is a mock implementation of the ThreadUseCase interface
type MockThreadUseCase struct {
	mock.Mock
}

// GetThread implements usecases.ThreadUseCase.
func (m *MockThreadUseCase) GetThread(id uint) (entities.Thread, error) {
	panic("unimplemented")
}

// GetThreads implements usecases.ThreadUseCase.
func (m *MockThreadUseCase) GetThreads() ([]entities.Thread, error) {
	panic("unimplemented")
}

func (m *MockThreadUseCase) CreateThread(thread entities.ThreadRequest, token string) (entities.Thread, error) {
	args := m.Called(thread, token)
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

		mockThreadUseCase.On("GetThreads").Return([]entities.Thread{}, nil)

		req := httptest.NewRequest(fiber.MethodGet, "/thread", nil)
		resp, _ := app.Test(req)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockThreadUseCase.AssertExpectations(t)
	})
}
