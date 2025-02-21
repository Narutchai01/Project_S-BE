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
func (m *MockCommentUsecase) GetComments(thread_id uint, token string) ([]entities.Comment, error) {
	args := m.Called(thread_id, token)
	return args.Get(0).([]entities.Comment), args.Error(1)
}

func (m *MockCommentUsecase) CreateComment(comment entities.Comment, token string) (entities.Comment, error) {
	args := m.Called(comment, token)
	return args.Get(0).(entities.Comment), args.Error(1)
}

func TestCreateCommentsuite(t *testing.T) {
	app := fiber.New()

	mockUsecase := new(MockCommentUsecase)
	handler := adapters.NewHttpCommentHandler(mockUsecase)

	app.Post("/comments", handler.CreateComment)

	t.Run("missing token", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/comments", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("invalid body", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/comments", strings.NewReader("invalid body"))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "valid-token")
		resp, _ := app.Test(req)

		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("usecase error", func(t *testing.T) {
		comment := entities.Comment{Text: "test comment"}
		mockUsecase.On("CreateComment", comment, "valid-token").Return(entities.Comment{}, errors.New("usecase error"))

		body, _ := json.Marshal(comment)
		req := httptest.NewRequest("POST", "/comments", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "valid-token")
		resp, _ := app.Test(req)

		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("successful creation", func(t *testing.T) {
		comment := entities.Comment{Text: "test comment", ThreadID: 1}
		mockUsecase.On("CreateComment", comment, "valid-token").Return(comment, nil)

		body, _ := json.Marshal(comment)
		req := httptest.NewRequest("POST", "/comments", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "valid-token")
		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)
	})
}
