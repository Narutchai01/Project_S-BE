package adapters_test

import (
	"bytes"
	"errors"
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

type MockCommunityUsecase struct {
	mock.Mock
}

// GetCommunitiesByUserID implements usecases.CommunityUseCase.
func (m *MockCommunityUsecase) GetCommunitiesByUserID(user_id uint, type_community string, token string) ([]entities.Community, error) {
	panic("unimplemented")
}

func (m *MockCommunityUsecase) CreateCommunityThread(community entities.Community, token string, files []*multipart.FileHeader, c *fiber.Ctx, communityType string) (entities.Community, error) {
	args := m.Called(community, token, files, c, communityType)
	return args.Get(0).(entities.Community), args.Error(1)
}

func (m *MockCommunityUsecase) GetCommunity(id uint, communityType string, token string) (entities.Community, error) {
	args := m.Called(id, communityType, token)
	return args.Get(0).(entities.Community), args.Error(1)
}

func (m *MockCommunityUsecase) GetCommunities(communityType string, token string) ([]entities.Community, error) {
	args := m.Called(communityType, token)
	return args.Get(0).([]entities.Community), args.Error(1)
}

func TestCreateThreadHandler(t *testing.T) {
	setup := func() (*MockCommunityUsecase, *adapters.HttpThreadRepository, *fiber.App) {
		mockService := new(MockCommunityUsecase)
		handler := adapters.NewHttpThreadRepository(mockService)

		app := fiber.New()
		app.Post("/thread", handler.CreateThread)

		return mockService, handler, app
	}

	expectedCommunity := entities.Community{Model: gorm.Model{ID: 1}, Title: "Test Title", Caption: "Test Caption"}

	t.Run("Create thread successfully", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("CreateCommunityThread", mock.Anything, "test_token", mock.Anything, mock.Anything, "Thread").Return(expectedCommunity, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("title", "Test Title")
		writer.WriteField("caption", "Test Caption")
		part, _ := writer.CreateFormFile("files", "test.jpg")
		part.Write([]byte("test data"))
		writer.Close()

		req := httptest.NewRequest("POST", "/thread", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test_token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("Missing token", func(t *testing.T) {
		_, _, app := setup()

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("title", "Test Title")
		writer.WriteField("caption", "Test Caption")
		part, _ := writer.CreateFormFile("files", "test.jpg")
		part.Write([]byte("test data"))
		writer.Close()

		req := httptest.NewRequest("POST", "/thread", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("Missing files", func(t *testing.T) {
		_, _, app := setup()

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("title", "Test Title")
		writer.WriteField("caption", "Test Caption")
		writer.Close()

		req := httptest.NewRequest("POST", "/thread", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test_token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Invalid form", func(t *testing.T) {
		_, _, app := setup()

		// Create a malformed multipart form
		body := new(bytes.Buffer)
		body.WriteString("--invalidboundary\r\n")
		body.WriteString("Content-Disposition: form-data; name=\"title\"\r\n\r\n")
		body.WriteString("Test Title\r\n")
		body.WriteString("--invalidboundary--\r\n")
		// Missing proper closing boundary

		req := httptest.NewRequest("POST", "/thread", body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary=invalidboundary")
		req.Header.Set("token", "test_token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Service error", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("CreateCommunityThread",
			mock.Anything, "test_token", mock.Anything, mock.Anything, "Thread",
		).Return(entities.Community{}, fiber.ErrInternalServerError)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("title", "Test Title")
		writer.WriteField("caption", "Test Caption")
		part, _ := writer.CreateFormFile("files", "test.jpg")
		part.Write([]byte("test data"))
		writer.Close()

		req := httptest.NewRequest("POST", "/thread", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test_token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("Multiple files", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("CreateCommunityThread",
			mock.Anything, "test_token", mock.Anything, mock.Anything, "Thread",
		).Return(expectedCommunity, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("title", "Test Title")
		writer.WriteField("caption", "Test Caption")

		part1, _ := writer.CreateFormFile("files", "test1.jpg")
		part1.Write([]byte("test data 1"))

		part2, _ := writer.CreateFormFile("files", "test2.jpg")
		part2.Write([]byte("test data 2"))

		writer.Close()

		req := httptest.NewRequest("POST", "/thread", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test_token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("Empty JSON payload", func(t *testing.T) {
		mockService, _, app := setup()
		// Even with empty JSON, the handler should pass an empty Community object
		mockService.On("CreateCommunityThread",
			entities.Community{}, "test_token", mock.Anything, mock.Anything, "Thread",
		).Return(expectedCommunity, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("title", "")
		writer.WriteField("caption", "")
		part, _ := writer.CreateFormFile("files", "test.jpg")
		part.Write([]byte("test data"))
		writer.Close()

		req := httptest.NewRequest("POST", "/thread", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "test_token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
func TestGetThreadHandler(t *testing.T) {
	setup := func() (*MockCommunityUsecase, *adapters.HttpThreadRepository, *fiber.App) {
		mockService := new(MockCommunityUsecase)
		handler := adapters.NewHttpThreadRepository(mockService)

		app := fiber.New()
		app.Get("/thread/:id", handler.GetThread)

		return mockService, handler, app
	}

	expectedCommunity := entities.Community{
		Model:   gorm.Model{ID: 1},
		Title:   "Test Title",
		Caption: "Test Caption",
	}

	t.Run("Get thread successfully", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetCommunity", uint(1), "Thread", "test_token").Return(expectedCommunity, nil)

		req := httptest.NewRequest("GET", "/thread/1", nil)
		req.Header.Set("token", "test_token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("Invalid ID parameter", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest("GET", "/thread/invalid", nil)
		req.Header.Set("token", "test_token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Thread not found", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetCommunity", uint(999), "Thread", "test_token").
			Return(entities.Community{}, errors.New("community not found"))

		req := httptest.NewRequest("GET", "/thread/999", nil)
		req.Header.Set("token", "test_token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("User not found", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetCommunity", uint(1), "Thread", "invalid_token").
			Return(entities.Community{}, errors.New("user not found"))

		req := httptest.NewRequest("GET", "/thread/1", nil)
		req.Header.Set("token", "invalid_token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("Other error from service", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetCommunity", uint(1), "Thread", "test_token").
			Return(entities.Community{}, errors.New("unknown error"))

		req := httptest.NewRequest("GET", "/thread/1", nil)
		req.Header.Set("token", "test_token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		// The implementation returns StatusOK even for unknown errors
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
func TestGetThreadsHandler(t *testing.T) {
	setup := func() (*MockCommunityUsecase, *adapters.HttpThreadRepository, *fiber.App) {
		mockService := new(MockCommunityUsecase)
		handler := adapters.NewHttpThreadRepository(mockService)

		app := fiber.New()
		app.Get("/thread", handler.GetThreads)

		return mockService, handler, app
	}

	expectedCommunities := []entities.Community{
		{
			Model:   gorm.Model{ID: 1},
			Title:   "Test Title 1",
			Caption: "Test Caption 1",
		},
		{
			Model:   gorm.Model{ID: 2},
			Title:   "Test Title 2",
			Caption: "Test Caption 2",
		},
	}

	t.Run("Get threads successfully", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetCommunities", "Thread", "test_token").Return(expectedCommunities, nil)

		req := httptest.NewRequest("GET", "/thread", nil)
		req.Header.Set("token", "test_token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("Missing token", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest("GET", "/thread", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("Service error", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetCommunities", "Thread", "test_token").
			Return([]entities.Community{}, errors.New("internal server error"))

		req := httptest.NewRequest("GET", "/thread", nil)
		req.Header.Set("token", "test_token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("Empty result set", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetCommunities", "Thread", "test_token").Return([]entities.Community{}, nil)

		req := httptest.NewRequest("GET", "/thread", nil)
		req.Header.Set("token", "test_token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
