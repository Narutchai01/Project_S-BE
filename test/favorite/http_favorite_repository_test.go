package adapter_test

import (
	"net/http/httptest"
	"testing"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/favorite"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFavoritesUsecase struct {
	mock.Mock
}

// FavoriteCommnetReviewSkincare implements usecases.FavoriteUseCase.
func (m *MockFavoritesUsecase) FavoriteCommnetReviewSkincare(comment_id uint, token string) (entities.FavoriteCommentReviewSkincare, error) {
	args := m.Called(comment_id, token)
	return args.Get(0).(entities.FavoriteCommentReviewSkincare), args.Error(1)
}

func (m *MockFavoritesUsecase) FavoriteCommentThread(thread_id uint, token string) (entities.FavoriteCommentThread, error) {
	args := m.Called(thread_id, token)
	return args.Get(0).(entities.FavoriteCommentThread), args.Error(1)
}

func (m *MockFavoritesUsecase) FavoriteThread(thread_id uint, token string) (entities.FavoriteThread, error) {
	args := m.Called(thread_id, token)
	return args.Get(0).(entities.FavoriteThread), args.Error(1)
}

func (m *MockFavoritesUsecase) FavoriteReviewSkincare(review_id uint, token string) (entities.FavoriteReviewSkincare, error) {
	args := m.Called(review_id, token)
	return args.Get(0).(entities.FavoriteReviewSkincare), args.Error(1)
}

func TestFavoriteCommentThreadHandler(t *testing.T) {

	setup := func() (*MockFavoritesUsecase, *adapters.HttpFavoriteHandler, *fiber.App) {
		mockFavoriteUseCase := new(MockFavoritesUsecase)
		httpFavoriteHandler := adapters.NewHttpFavoriteHandler(mockFavoriteUseCase)
		app := fiber.New()

		app.Post("/favorite/comment/thread/:id", httpFavoriteHandler.HandleFavoriteCommentThread)

		return mockFavoriteUseCase, httpFavoriteHandler, app
	}

	t.Run("FavoriteComment", func(t *testing.T) {
		mockFavortieUseCase, _, app := setup()

		mockFavortieUseCase.On("FavoriteCommentThread", uint(1), "token").Return(entities.FavoriteCommentThread{}, nil)

		req := httptest.NewRequest(fiber.MethodPost, "/favorite/comment/thread/1", nil)
		req.Header.Set("token", "token")
		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)

		mockFavortieUseCase.AssertExpectations(t)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockFavoriteUseCase, _, app := setup()

		mockFavoriteUseCase.On("FavoriteComment", uint(1), "token").Return(entities.FavoriteCommentThread{}, nil)

		req := httptest.NewRequest(fiber.MethodPost, "/favorite/comment/thread/1", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, 401, resp.StatusCode)

	})

}

func TestFavoriteThreads(t *testing.T) {

	setup := func() (*MockFavoritesUsecase, *adapters.HttpFavoriteHandler, *fiber.App) {
		mockFavoriteUseCase := new(MockFavoritesUsecase)
		httpFavoriteHandler := adapters.NewHttpFavoriteHandler(mockFavoriteUseCase)
		app := fiber.New()

		app.Post("/favorite/thread/:id", httpFavoriteHandler.HandleFavoriteThread)

		return mockFavoriteUseCase, httpFavoriteHandler, app
	}

	t.Run("FavoriteThread", func(t *testing.T) {
		mockFavortieUseCase, _, app := setup()

		mockFavortieUseCase.On("FavoriteThread", uint(1), "token").Return(entities.FavoriteThread{}, nil)

		req := httptest.NewRequest(fiber.MethodPost, "/favorite/thread/1", nil)
		req.Header.Set("token", "token")
		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)

		mockFavortieUseCase.AssertExpectations(t)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockFavoriteUseCase, _, app := setup()

		mockFavoriteUseCase.On("FavoriteThread", uint(1), "token").Return(entities.FavoriteThread{}, nil)

		req := httptest.NewRequest(fiber.MethodPost, "/favorite/thread/1", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, 401, resp.StatusCode)

	})
}
func TestFavoriteReviewSkincareHandler(t *testing.T) {

	setup := func() (*MockFavoritesUsecase, *adapters.HttpFavoriteHandler, *fiber.App) {
		mockFavoriteUseCase := new(MockFavoritesUsecase)
		httpFavoriteHandler := adapters.NewHttpFavoriteHandler(mockFavoriteUseCase)
		app := fiber.New()

		app.Post("/favorite/review/skincare/:id", httpFavoriteHandler.HandleFavoriteReviewSkincare)

		return mockFavoriteUseCase, httpFavoriteHandler, app
	}

	t.Run("FavoriteReviewSkincare", func(t *testing.T) {
		mockFavoriteUseCase, _, app := setup()

		mockFavoriteUseCase.On("FavoriteReviewSkincare", uint(1), "token").Return(entities.FavoriteReviewSkincare{}, nil)

		req := httptest.NewRequest(fiber.MethodPost, "/favorite/review/skincare/1", nil)
		req.Header.Set("token", "token")
		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)

		mockFavoriteUseCase.AssertExpectations(t)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockFavoriteUseCase, _, app := setup()

		mockFavoriteUseCase.On("FavoriteReviewSkincare", uint(1), "token").Return(entities.FavoriteReviewSkincare{}, nil)

		req := httptest.NewRequest(fiber.MethodPost, "/favorite/review/skincare/1", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, 401, resp.StatusCode)

	})

	t.Run("BadRequest", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest(fiber.MethodPost, "/favorite/review/skincare/invalid", nil)
		req.Header.Set("token", "token")
		resp, _ := app.Test(req)

		assert.Equal(t, 400, resp.StatusCode)
	})
}
func TestFavoriteCommentReviewSkincareHandler(t *testing.T) {

	setup := func() (*MockFavoritesUsecase, *adapters.HttpFavoriteHandler, *fiber.App) {
		mockFavoriteUseCase := new(MockFavoritesUsecase)
		httpFavoriteHandler := adapters.NewHttpFavoriteHandler(mockFavoriteUseCase)
		app := fiber.New()

		app.Post("/favorite/comment/review/skincare/:id", httpFavoriteHandler.HandleFavoriteCommentReviewSkincare)

		return mockFavoriteUseCase, httpFavoriteHandler, app
	}

	t.Run("FavoriteCommentReviewSkincare", func(t *testing.T) {
		mockFavoriteUseCase, _, app := setup()

		mockFavoriteUseCase.On("FavoriteCommnetReviewSkincare", uint(1), "token").Return(entities.FavoriteCommentReviewSkincare{}, nil)

		req := httptest.NewRequest(fiber.MethodPost, "/favorite/comment/review/skincare/1", nil)
		req.Header.Set("token", "token")
		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)

		mockFavoriteUseCase.AssertExpectations(t)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockFavoriteUseCase, _, app := setup()

		mockFavoriteUseCase.On("FavoriteCommnetReviewSkincare", uint(1), "token").Return(entities.FavoriteCommentReviewSkincare{}, nil)

		req := httptest.NewRequest(fiber.MethodPost, "/favorite/comment/review/skincare/1", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, 401, resp.StatusCode)

	})

	t.Run("BadRequest", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest(fiber.MethodPost, "/favorite/comment/review/skincare/invalid", nil)
		req.Header.Set("token", "token")
		resp, _ := app.Test(req)

		assert.Equal(t, 400, resp.StatusCode)
	})
}
