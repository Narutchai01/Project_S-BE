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

// Favorite implements usecases.FavoriteUseCase.
func (m *MockFavoritesUsecase) Favorite(favorite entities.Favorite, type_community string, token string) (entities.Favorite, error) {
	args := m.Called(favorite, type_community, token)
	return args.Get(0).(entities.Favorite), args.Error(1)
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

		mockFavortieUseCase.On("Favorite", mock.Anything, mock.Anything, mock.Anything).Return(entities.Favorite{}, nil)

		req := httptest.NewRequest(fiber.MethodPost, "/favorite/comment/thread/1", nil)
		req.Header.Set("token", "token")
		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)

		mockFavortieUseCase.AssertExpectations(t)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockFavoriteUseCase, _, app := setup()

		mockFavoriteUseCase.On("Favorite", mock.Anything, mock.Anything, mock.Anything).Return(entities.Favorite{}, nil)

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

		mockFavortieUseCase.On("Favorite", mock.Anything, mock.Anything, mock.Anything).Return(entities.Favorite{}, nil)

		req := httptest.NewRequest(fiber.MethodPost, "/favorite/thread/1", nil)
		req.Header.Set("token", "token")
		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)

		mockFavortieUseCase.AssertExpectations(t)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockFavoriteUseCase, _, app := setup()

		mockFavoriteUseCase.On("Favorite", mock.Anything, mock.Anything, mock.Anything).Return(entities.Favorite{}, nil)

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

		mockFavoriteUseCase.On("Favorite", mock.Anything, mock.Anything, mock.Anything).Return(entities.Favorite{}, nil)

		req := httptest.NewRequest(fiber.MethodPost, "/favorite/review/skincare/1", nil)
		req.Header.Set("token", "token")
		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)

		mockFavoriteUseCase.AssertExpectations(t)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockFavoriteUseCase, _, app := setup()

		mockFavoriteUseCase.On("Favorite", mock.Anything, mock.Anything, mock.Anything).Return(entities.Favorite{}, nil)

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

		mockFavoriteUseCase.On("Favorite", mock.Anything, mock.Anything, mock.Anything).Return(entities.Favorite{}, nil)

		req := httptest.NewRequest(fiber.MethodPost, "/favorite/comment/review/skincare/1", nil)
		req.Header.Set("token", "token")
		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)

		mockFavoriteUseCase.AssertExpectations(t)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockFavoriteUseCase, _, app := setup()

		mockFavoriteUseCase.On("Favorite", mock.Anything, mock.Anything, mock.Anything).Return(entities.Favorite{}, nil)

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
