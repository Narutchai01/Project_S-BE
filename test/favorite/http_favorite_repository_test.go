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

func (m *MockFavoritesUsecase) FavoriteComment(thread_id uint, token string) (entities.FavoriteComment, error) {
	args := m.Called(thread_id, token)
	return args.Get(0).(entities.FavoriteComment), args.Error(1)
}

func (m *MockFavoritesUsecase) FavoriteThread(thread_id uint, token string) (entities.FavoriteThread, error) {
	args := m.Called(thread_id, token)
	return args.Get(0).(entities.FavoriteThread), args.Error(1)
}

func (m *MockFavoritesUsecase) FavoriteReviewSkincare(review_id uint, token string) (entities.FavoriteReviewSkincare, error) {
	args := m.Called(review_id, token)
	return args.Get(0).(entities.FavoriteReviewSkincare), args.Error(1)
}

func TestFavoriteComment(t *testing.T) {

	setup := func() (*MockFavoritesUsecase, *adapters.HttpFavoriteHandler, *fiber.App) {
		mockFavoriteUseCase := new(MockFavoritesUsecase)
		httpFavoriteHandler := adapters.NewHttpFavoriteHandler(mockFavoriteUseCase)
		app := fiber.New()

		app.Post("/favorite/comment/:id", httpFavoriteHandler.HandleFavoriteComment)

		return mockFavoriteUseCase, httpFavoriteHandler, app
	}

	t.Run("FavoriteComment", func(t *testing.T) {
		mockFavortieUseCase, _, app := setup()

		mockFavortieUseCase.On("FavoriteComment", uint(1), "token").Return(entities.FavoriteComment{}, nil)

		req := httptest.NewRequest(fiber.MethodPost, "/favorite/comment/1", nil)
		req.Header.Set("token", "token")
		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)

		mockFavortieUseCase.AssertExpectations(t)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockFavoriteUseCase, _, app := setup()

		mockFavoriteUseCase.On("FavoriteComment", uint(1), "token").Return(entities.FavoriteComment{}, nil)

		req := httptest.NewRequest(fiber.MethodPost, "/favorite/comment/1", nil)
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
