package adapters_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/comment"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCommentUsecase struct {
	mock.Mock
}

// GetComments implements usecases.CommentUsecase.
func (m *MockCommentUsecase) GetCommentsThread(thread_id uint, token string) ([]entities.CommentThread, error) {
	args := m.Called(thread_id, token)
	return args.Get(0).([]entities.CommentThread), args.Error(1)
}

func (m *MockCommentUsecase) CreateCommentThread(comment entities.CommentThread, token string) (entities.CommentThread, error) {
	args := m.Called(comment, token)
	return args.Get(0).(entities.CommentThread), args.Error(1)
}

func (m *MockCommentUsecase) CreateCommentReviewSkicnare(comment entities.CommentReviewSkicare, token string) (entities.CommentReviewSkicare, error) {
	args := m.Called(comment, token)
	return args.Get(0).(entities.CommentReviewSkicare), args.Error(1)
}

func (m *MockCommentUsecase) GetCommentsReviewSkincare(review_id uint, token string) ([]entities.CommentReviewSkicare, error) {
	args := m.Called(review_id, token)
	return args.Get(0).([]entities.CommentReviewSkicare), args.Error(1)
}

func TestCreateCommentThreadHandler(t *testing.T) {
	app := fiber.New()

	mockUsecase := new(MockCommentUsecase)
	handler := adapters.NewHttpCommentHandler(mockUsecase)

	app.Post("/comment/thread", handler.CreateCommentThread)

	t.Run("missing token", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/comment/thread", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("invalid body", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/comment/thread", strings.NewReader("invalid body"))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "valid-token")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("usecase error", func(t *testing.T) {
		comment := entities.CommentThread{Text: "test comment"}
		mockUsecase.On("CreateCommentThread", comment, "valid-token").Return(entities.CommentThread{}, errors.New("usecase error"))

		body, _ := json.Marshal(comment)
		req := httptest.NewRequest("POST", "/comment/thread", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "valid-token")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("successful creation", func(t *testing.T) {
		comment := entities.CommentThread{Text: "test comment", ThreadID: 1}
		mockUsecase.On("CreateCommentThread", comment, "valid-token").Return(comment, nil)

		body, _ := json.Marshal(comment)
		req := httptest.NewRequest("POST", "/comment/thread", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "valid-token")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	})
}
func TestHandleGetCommentReviewSkincare(t *testing.T) {
	app := fiber.New()

	mockUsecase := new(MockCommentUsecase)
	handler := adapters.NewHttpCommentHandler(mockUsecase)

	app.Get("/comment/reviews/skincare/:review_id", handler.HandleGetCommentReviewSkincare)

	t.Run("Success", func(t *testing.T) {
		mockUsecase.On("GetCommentsReviewSkincare", uint(1), mock.Anything).Return([]entities.CommentReviewSkicare{}, nil)

		req := httptest.NewRequest("GET", "/comment/reviews/skincare/1", nil)
		req.Header.Set("token", mock.Anything)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("Invalid review_id", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/comment/reviews/skincare/invalid", nil)
		req.Header.Set("token", mock.Anything)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Miss token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/comment/reviews/skincare/1", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})
}
func TestCreateCommentReviewSkicnareHandler(t *testing.T) {
	app := fiber.New()

	mockUsecase := new(MockCommentUsecase)
	handler := adapters.NewHttpCommentHandler(mockUsecase)

	app.Post("/comment/reviews/skincare", handler.CreateCommentReviewSkicnare)

	t.Run("missing token", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/comment/reviews/skincare", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("invalid body", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/comment/reviews/skincare", strings.NewReader("invalid body"))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "valid-token")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("successful creation", func(t *testing.T) {
		comment := entities.CommentReviewSkicare{
			ReviewSkincareID: 1,
			Content:          "test ",
		}
		mockUsecase.On("CreateCommentReviewSkicnare", comment, "valid-token").Return(comment, nil)

		body, _ := json.Marshal(comment)
		req := httptest.NewRequest("POST", "/comment/reviews/skincare", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "valid-token")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	})
}
func TestGetCommentsThreadHandler(t *testing.T) {
	app := fiber.New()

	mockUsecase := new(MockCommentUsecase)
	handler := adapters.NewHttpCommentHandler(mockUsecase)

	app.Get("/comment/thread/:thread_id", handler.GetCommentsThread)

	t.Run("Success", func(t *testing.T) {
		mockUsecase.On("GetCommentsThread", uint(1), mock.Anything).Return([]entities.CommentThread{}, nil)

		req := httptest.NewRequest("GET", "/comment/thread/1", nil)
		req.Header.Set("token", mock.Anything)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("Invalid thread_id", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/comment/thread/invalid", nil)
		req.Header.Set("token", mock.Anything)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

}
