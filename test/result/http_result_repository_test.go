package result_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"

	"testing"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/result"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockResultService struct {
	mock.Mock
}

func (m *MockResultService) CreateResult(result entities.Result) (entities.Result, error) {
	args := m.Called(result)
	return args.Get(0).(entities.Result), args.Error(1)
}

func (m *MockResultService) GetResults() ([]entities.Result, error) {
	args := m.Called()
	return args.Get(0).([]entities.Result), args.Error(1)
}

func (m *MockResultService) GetResultById(id int) (entities.Result, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Result), args.Error(1)
}

func (m *MockResultService) UpdateResultById(id int, result entities.Result) (entities.Result, error) {
	args := m.Called(id, result)
	return args.Get(0).(entities.Result), args.Error(1)
}

func (m *MockResultService) DeleteResultById(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockResultService) GetResultsByUserIdFromToken(token string) ([]entities.Result, error) {
	args := m.Called(token)
	return args.Get(0).([]entities.Result), args.Error(1)
}

func (m *MockResultService) GetResultsByUserId(user_id int) ([]entities.Result, error) {
	args := m.Called(user_id)
	return args.Get(0).([]entities.Result), args.Error(1)
}

func (m *MockResultService) GetLatestResultByUserIdFromToken(token string) (entities.Result, error) {
	args := m.Called(token)
	return args.Get(0).(entities.Result), args.Error(1)
}
func TestCreateResultHandler(t *testing.T) {
	setup := func() (*MockResultService, *adapters.HttpResultHandler, *fiber.App) {
		mockService := new(MockResultService)
		handler := adapters.NewHttpResultHandler(mockService)

		app := fiber.New()
		app.Post("/result", handler.CreateResult)

		return mockService, handler, app
	}

	expectData := entities.Result{
		Model: gorm.Model{
			ID: 1,
		},
		Image: "image_url_test",
		UserId: 1,
		AcneType: []entities.Acne_Facial_Result{
			{ID: 1, Count: 10},
			{ID: 2, Count: 5},
		},
		FacialType: []entities.Acne_Facial_Result{
			{ID: 1, Count: 10},
			{ID: 2, Count: 5},
		},
		SkinType: 1,
		Skincare: []uint{1, 2, 3},
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("CreateResult",
			mock.Anything,
		).Return(expectData, nil)

		body, _ := json.Marshal(expectData)

		req := httptest.NewRequest("POST", "/result", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed in body parser", func(t *testing.T) {
		_, _, app := setup()
		req := httptest.NewRequest("POST", "/result", bytes.NewBuffer([]byte("invalid body")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("failed to create result", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("CreateResult",
			mock.Anything,
		).Return(entities.Result{}, errors.New("service error"))

		body, _ := json.Marshal(expectData)

		req := httptest.NewRequest("POST", "/result", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetResultsHandler(t *testing.T) {
	setup := func() (*MockResultService, *adapters.HttpResultHandler, *fiber.App) {
		mockService := new(MockResultService)
		handler := adapters.NewHttpResultHandler(mockService)

		app := fiber.New()
		app.Get("/result", handler.GetResults)

		return mockService, handler, app
	}

	expectData := []entities.Result{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Image: "image_url_test",
			UserId: 1,
			AcneType: []entities.Acne_Facial_Result{
				{ID: 1, Count: 10},
				{ID: 2, Count: 5},
			},
			FacialType: []entities.Acne_Facial_Result{
				{ID: 1, Count: 10},
				{ID: 2, Count: 5},
			},
			SkinType: 1,
			Skincare: []uint{1, 2, 3},
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Image: "image_url_test",
			UserId: 1,
			AcneType: []entities.Acne_Facial_Result{
				{ID: 1, Count: 10},
				{ID: 2, Count: 5},
			},
			FacialType: []entities.Acne_Facial_Result{
				{ID: 1, Count: 10},
				{ID: 2, Count: 5},
			},
			SkinType: 1,
			Skincare: []uint{1, 2, 3},
		},
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetResults").Return(expectData, nil)

		req := httptest.NewRequest("GET", "/result", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get results", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GetResults").Return([]entities.Result{}, errors.New("service error"))

		req := httptest.NewRequest("GET", "/result", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetResultByIdHandler(t *testing.T) {
	setup := func() (*MockResultService, *adapters.HttpResultHandler, *fiber.App) {
		mockService := new(MockResultService)
		handler := adapters.NewHttpResultHandler(mockService)

		app := fiber.New()
		app.Get("/result/:id", handler.GetResultById)

		return mockService, handler, app
	}

	expectData := entities.Result{
		Model: gorm.Model{
			ID: 1,
		},
		Image: "image_url_test",
		UserId: 1,
		AcneType: []entities.Acne_Facial_Result{
			{ID: 1, Count: 10},
			{ID: 2, Count: 5},
		},
		FacialType: []entities.Acne_Facial_Result{
			{ID: 1, Count: 10},
			{ID: 2, Count: 5},
		},
		SkinType: 1,
		Skincare: []uint{1, 2, 3},
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetResultById", int(expectData.ID)).Return(expectData, nil)

		req := httptest.NewRequest("GET", "/result/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to convert id to int", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("GET", "/result/error-id", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get skin", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GetResultById", int(expectData.ID)).Return(entities.Result{}, errors.New("service error"))

		req := httptest.NewRequest("GET", "/result/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateResultByIdHandler(t *testing.T) {
	setup := func() (*MockResultService, *adapters.HttpResultHandler, *fiber.App) {
		mockService := new(MockResultService)
		handler := adapters.NewHttpResultHandler(mockService)

		app := fiber.New()
		app.Put("/result/:id", handler.UpdateResultById)

		return mockService, handler, app
	}

	expectData := entities.Result{
		Model: gorm.Model{
			ID: 1,
		},
		Image: "image_url_test",
		UserId: 1,
		AcneType: []entities.Acne_Facial_Result{
			{ID: 1, Count: 10},
			{ID: 2, Count: 5},
		},
		FacialType: []entities.Acne_Facial_Result{
			{ID: 1, Count: 10},
			{ID: 2, Count: 5},
		},
		SkinType: 1,
		Skincare: []uint{1, 2, 3},
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateResultById",
			mock.Anything,
			mock.Anything,
		).Return(expectData, nil)

		body, _ := json.Marshal(expectData)

		req := httptest.NewRequest("PUT", "/result/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to convert id to int", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("PUT", "/result/error-id", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed in body parser", func(t *testing.T) {
		_, _, app := setup()
		req := httptest.NewRequest("PUT", "/result/1", bytes.NewBuffer([]byte("invalid body")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("failed to update result", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateResultById",
			mock.Anything,
			mock.Anything,
		).Return(entities.Result{}, errors.New("service error"))

		body, _ := json.Marshal(expectData)

		req := httptest.NewRequest("PUT", "/result/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestDeleteResultByIdHandler(t *testing.T) {
	setup := func() (*MockResultService, *adapters.HttpResultHandler, *fiber.App) {
		mockService := new(MockResultService)
		handler := adapters.NewHttpResultHandler(mockService)

		app := fiber.New()
		app.Delete("/result/:id", handler.DeleteResultById)

		return mockService, handler, app
	}


	expectData := entities.Acne{
		Model: gorm.Model{
			ID: 1,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("DeleteResultById", int(expectData.ID)).Return(nil)

		req := httptest.NewRequest("DELETE", "/result/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to convert id to int", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("DELETE", "/result/error-id", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to delete result", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("DeleteResultById", int(expectData.ID)).Return(errors.New("service error"))

		req := httptest.NewRequest("DELETE", "/result/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetResultsByUserIdFromTokenHandler(t *testing.T) {
	setup := func() (*MockResultService, *adapters.HttpResultHandler, *fiber.App) {
		mockService := new(MockResultService)
		handler := adapters.NewHttpResultHandler(mockService)

		app := fiber.New()
		app.Get("/user/result", handler.GetResultsByUserIdFromToken)

		return mockService, handler, app
	}

	expectData := []entities.Result{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Image: "image_url_test",
			UserId: 1,
			AcneType: []entities.Acne_Facial_Result{
				{ID: 1, Count: 10},
				{ID: 2, Count: 5},
			},
			FacialType: []entities.Acne_Facial_Result{
				{ID: 1, Count: 10},
				{ID: 2, Count: 5},
			},
			SkinType: 1,
			Skincare: []uint{1, 2, 3},
		},
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetResultsByUserIdFromToken", "Bearer example-token").Return(expectData, nil)

		req := httptest.NewRequest("GET", "/user/result", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "Bearer example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get results", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GetResultsByUserIdFromToken", "Bearer example-token").Return([]entities.Result{}, errors.New("service error"))

		req := httptest.NewRequest("GET", "/user/result", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "Bearer example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetResultsByUserIdHandler(t *testing.T) {
	setup := func() (*MockResultService, *adapters.HttpResultHandler, *fiber.App) {
		mockService := new(MockResultService)
		handler := adapters.NewHttpResultHandler(mockService)

		app := fiber.New()
		app.Get("/result/user/:userId", handler.GetResultsByUserId)

		return mockService, handler, app
	}

	expectData := []entities.Result{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Image: "image_url_test",
			UserId: 1,
			AcneType: []entities.Acne_Facial_Result{
				{ID: 1, Count: 10},
				{ID: 2, Count: 5},
			},
			FacialType: []entities.Acne_Facial_Result{
				{ID: 1, Count: 10},
				{ID: 2, Count: 5},
			},
			SkinType: 1,
			Skincare: []uint{1, 2, 3},
		},
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetResultsByUserId", int(expectData[0].UserId)).Return(expectData, nil)

		req := httptest.NewRequest("GET", "/result/user/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to convert userId to int", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("GET", "/result/user/error-id", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get results", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GetResultsByUserId", int(expectData[0].UserId)).Return([]entities.Result{}, errors.New("service error"))

		req := httptest.NewRequest("GET", "/result/user/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetLatestResultByUserIdFromTokenHandler(t *testing.T) {
	setup := func() (*MockResultService, *adapters.HttpResultHandler, *fiber.App) {
		mockService := new(MockResultService)
		handler := adapters.NewHttpResultHandler(mockService)

		app := fiber.New()
		app.Get("/user/result/latest", handler.GetLatestResultByUserIdFromToken)

		return mockService, handler, app
	}

	expectData := entities.Result{
			Model: gorm.Model{
				ID: 1,
			},
			Image: "image_url_test",
			UserId: 1,
			AcneType: []entities.Acne_Facial_Result{
				{ID: 1, Count: 10},
				{ID: 2, Count: 5},
			},
			FacialType: []entities.Acne_Facial_Result{
				{ID: 1, Count: 10},
				{ID: 2, Count: 5},
			},
			SkinType: 1,
			Skincare: []uint{1, 2, 3},
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetLatestResultByUserIdFromToken", "Bearer example-token").Return(expectData, nil)

		req := httptest.NewRequest("GET", "/user/result/latest", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "Bearer example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get results", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GetLatestResultByUserIdFromToken", "Bearer example-token").Return(entities.Result{}, errors.New("service error"))

		req := httptest.NewRequest("GET", "/user/result/latest", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "Bearer example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}