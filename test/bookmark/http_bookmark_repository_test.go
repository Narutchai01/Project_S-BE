package adapters_test

import (
	"net/http/httptest"
	"testing"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/bookmark"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBookmarkUsecase struct {
	mock.Mock
}

func (m *MockBookmarkUsecase) BookmarkThread(user_id uint, token string) (entities.BookmarkThread, error) {
	args := m.Called(user_id, token)
	return args.Get(0).(entities.BookmarkThread), args.Error(1)
}

func TestBookmarkThread(t *testing.T) {
	app := fiber.New()

	mockBookmarkUsecase := new(MockBookmarkUsecase)
	handler := adapters.NewHttpBookmarkHandler(mockBookmarkUsecase)

	app.Post("/bookmark/:id", handler.BookMarkThread)

	t.Run("Bookmark Thread Success", func(t *testing.T) {
		mockBookmarkUsecase.On("BookmarkThread", uint(1), "token").Return(entities.BookmarkThread{}, nil)

		req := httptest.NewRequest("POST", "/bookmark/1", nil)
		req.Header.Add("token", "token")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockBookmarkUsecase.On("BookmarkThread", uint(1), "").Return(nil, assert.AnError)

		req := httptest.NewRequest("POST", "/bookmark/1", nil)
		req.Header.Add("token", "")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})
}
