package adapters_test

import (
	"bytes"
	"mime/multipart"
	"net/http/httptest"
	"testing"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/thread"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockThreadUseCase struct {
	mock.Mock
}

// GetThread implements usecases.ThreadUseCase.
func (m *MockThreadUseCase) GetThread(thread_id uint, token string) (entities.Thread, error) {
	args := m.Called(thread_id, token)
	return args.Get(0).(entities.Thread), args.Error(1)
}

func (m *MockThreadUseCase) CreateThread(thread entities.Thread, token string, files []*multipart.FileHeader, c *fiber.Ctx) (entities.Thread, error) {
	args := m.Called(thread, token, files, c)
	return args.Get(0).(entities.Thread), args.Error(1)
}

func TestCreateThread(t *testing.T) {
	app := fiber.New()
	mockUseCase := new(MockThreadUseCase)
	repo := adapters.NewHttpThreadRepository(mockUseCase)

	app.Post("/threads", repo.CreateThread)

	t.Run("successful creation", func(t *testing.T) {
		thread := entities.Thread{Title: "Test Title", Caption: "Test Caption"}
		mockUseCase.On("CreateThread", thread, "test-token", mock.Anything, mock.Anything).Return(thread, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("title", "Test Title")
		writer.WriteField("caption", "Test Caption")
		writer.Close()

		req := httptest.NewRequest("POST", "/threads", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("invalid request body", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/threads", bytes.NewBufferString("invalid body"))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("missing title or caption", func(t *testing.T) {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("title", "")
		writer.WriteField("caption", "")
		writer.Close()

		req := httptest.NewRequest("POST", "/threads", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("multipart form error", func(t *testing.T) {
		thread := entities.Thread{Title: "Test Title", Caption: "Test Caption"}
		mockUseCase.On("CreateThread", thread, "test-token", mock.Anything, mock.Anything).Return(thread, nil)

		req := httptest.NewRequest("POST", "/threads", bytes.NewBufferString("invalid body"))
		req.Header.Set("Content-Type", "multipart/form-data")
		req.Header.Set("token", "test-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}
func TestGetThread(t *testing.T) {
	app := fiber.New()
	mockUseCase := new(MockThreadUseCase)
	repo := adapters.NewHttpThreadRepository(mockUseCase)

	app.Get("/threads/:id", repo.GetThread)

	t.Run("successful retrieval", func(t *testing.T) {
		thread := entities.Thread{Model: gorm.Model{ID: 1}, Title: "Test Title", Caption: "Test Caption"}
		mockUseCase.On("GetThread", uint(1), "test-token").Return(thread, nil)

		req := httptest.NewRequest("GET", "/threads/1", nil)
		req.Header.Set("token", "test-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("invalid thread id", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/threads/invalid", nil)
		req.Header.Set("token", "test-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

}
