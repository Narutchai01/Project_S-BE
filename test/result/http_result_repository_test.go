package result_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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