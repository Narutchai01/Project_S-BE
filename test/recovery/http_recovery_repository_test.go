package adapters_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"

	"testing"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/recovery"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockRecoveryService struct {
	mock.Mock
}

func (m *MockRecoveryService) CreateRecovery(recovery entities.Recovery, email string, c *fiber.Ctx) (entities.Recovery, error) {
	args := m.Called(recovery, email, c)
	return args.Get(0).(entities.Recovery), args.Error(1)
}

func (m *MockRecoveryService) DeleteRecoveryById(id int) (entities.Recovery, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Recovery), args.Error(1)
}

func (m *MockRecoveryService) GetRecoveries() ([]entities.Recovery, error) {
	args := m.Called()
	return args.Get(0).([]entities.Recovery), args.Error(1)
}

func (m *MockRecoveryService) GetRecoveryById(id int) (entities.Recovery, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Recovery), args.Error(1)
}

func (m *MockRecoveryService) GetRecoveryByUserId(user_id int) (entities.Recovery, error) {
	args := m.Called(user_id)
	return args.Get(0).(entities.Recovery), args.Error(1)
}

func (m *MockRecoveryService) OtpValidation(id int, otp string) (bool, error) {
	args := m.Called(id, otp)
	return args.Get(0).(bool), args.Error(1)
}

func (m *MockRecoveryService) UpdateRecoveryOtpById(recovery entities.Recovery, email string) (entities.Recovery, error) {
	args := m.Called(recovery, email)
	return args.Get(0).(entities.Recovery), args.Error(1)
}

func TestCreateRecovery(t *testing.T) {
	setup := func() (*MockRecoveryService, *adapters.HttpRecoveryHandler, *fiber.App) {
		mockService := new(MockRecoveryService)
		handler := adapters.NewHttpRecoveryHandler(mockService)

		app := fiber.New()
		app.Post("/recovery", handler.CreateRecovery)

		return mockService, handler, app
	}

	type RequestBody struct {
		Email  string `json:"email"`
		UserId int    `json:"user_id"`
		OTP    string `json:"otp"`
	}

	expectDataFromReq := RequestBody{
		Email:  "aut@gmail.com",
		UserId: 1,
		OTP:    "123456",
	}

	expectData := entities.Recovery{
		Model: gorm.Model{
			ID: 1,
		},
		OTP:    "123456",
		UserID: 1,
	}

	t.Run("success create recovery", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("GetRecoveryByUserId", int(expectDataFromReq.UserId)).Return(entities.Recovery{}, errors.New("service error"))
		mockService.On("CreateRecovery",
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(expectData, nil)

		body, _ := json.Marshal(expectDataFromReq)

		req := httptest.NewRequest("POST", "/recovery", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("success update recovery", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("GetRecoveryByUserId", int(expectData.UserID)).Return(expectData, nil)
		mockService.On("UpdateRecoveryOtpById",
			mock.Anything,
			mock.Anything,
		).Return(expectData, nil)

		body, _ := json.Marshal(expectDataFromReq)

		req := httptest.NewRequest("POST", "/recovery", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("invalid request body", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest("POST", "/recovery", bytes.NewBuffer([]byte("")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("failed to update recovery", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("GetRecoveryByUserId", int(expectData.UserID)).Return(expectData, nil)
		mockService.On("UpdateRecoveryOtpById",
			mock.Anything,
			mock.Anything,
		).Return(entities.Recovery{}, errors.New("service error"))

		body, _ := json.Marshal(expectDataFromReq)

		req := httptest.NewRequest("POST", "/recovery", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to create recovery", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("GetRecoveryByUserId", int(expectData.UserID)).Return(entities.Recovery{}, errors.New("service error"))
		mockService.On("CreateRecovery",
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(entities.Recovery{}, errors.New("service error"))

		body, _ := json.Marshal(expectDataFromReq)

		req := httptest.NewRequest("POST", "/recovery", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

}

func TestDeleteRecoveryByIdHandler(t *testing.T) {
	setup := func() (*MockRecoveryService, *adapters.HttpRecoveryHandler, *fiber.App) {
		mockService := new(MockRecoveryService)
		handler := adapters.NewHttpRecoveryHandler(mockService)

		app := fiber.New()
		app.Delete("/recovery/:id", handler.DeleteRecoveryById)

		return mockService, handler, app
	}

	expectData := entities.Recovery{
		Model: gorm.Model{
			ID: 1,
		},
		OTP:    "123456",
		UserID: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("DeleteRecoveryById", int(expectData.ID)).Return(expectData, nil)

		req := httptest.NewRequest("DELETE", "/recovery/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to convert id to int", func(t *testing.T) {
		mockService, _, app := setup()
		req := httptest.NewRequest("DELETE", "/recovery/error-id", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to delete recovery", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("DeleteRecoveryById", int(expectData.ID)).Return(entities.Recovery{}, errors.New("service error"))

		req := httptest.NewRequest("DELETE", "/recovery/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetRecoveriesHandler(t *testing.T) {
	setup := func() (*MockRecoveryService, *adapters.HttpRecoveryHandler, *fiber.App) {
		mockService := new(MockRecoveryService)
		handler := adapters.NewHttpRecoveryHandler(mockService)

		app := fiber.New()
		app.Get("/recovery", handler.GetRecoveries)

		return mockService, handler, app
	}

	expectData := []entities.Recovery{
		{
			Model: gorm.Model{
				ID: 1,
			},
			OTP:    "123456",
			UserID: 1,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("GetRecoveries").Return(expectData, nil)

		req := httptest.NewRequest("GET", "/recovery", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to get recoveries", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("GetRecoveries").Return([]entities.Recovery{}, errors.New("service error"))

		req := httptest.NewRequest("GET", "/recovery", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestOtpValidationHandler(t *testing.T) {
	setup := func() (*MockRecoveryService, *adapters.HttpRecoveryHandler, *fiber.App) {
		mockService := new(MockRecoveryService)
		handler := adapters.NewHttpRecoveryHandler(mockService)

		app := fiber.New()
		app.Post("/validation", handler.OtpValidation)

		return mockService, handler, app
	}

	expectData := entities.Recovery{
		Model: gorm.Model{
			ID: 1,
		},
		OTP:    "123456",
		UserID: 1,
	}

	t.Run("success create recovery", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("OtpValidation", int(expectData.UserID), expectData.OTP).Return(true, nil)
		mockService.On("DeleteRecoveryById",
			mock.Anything,
		).Return(expectData, nil)

		body, _ := json.Marshal(expectData)

		req := httptest.NewRequest("POST", "/validation", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("invalid request body", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest("POST", "/validation", bytes.NewBuffer([]byte("")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("failed in otpvalidation", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("OtpValidation", int(expectData.UserID), expectData.OTP).Return(false, errors.New("service error"))

		body, _ := json.Marshal(expectData)

		req := httptest.NewRequest("POST", "/validation", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("failed to delete recovery", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.ExpectedCalls = nil
		mockService.On("OtpValidation", int(expectData.UserID), expectData.OTP).Return(true, nil)
		mockService.On("DeleteRecoveryById", int(expectData.ID)).Return(entities.Recovery{}, errors.New("service error"))

		body, _ := json.Marshal(expectData)

		req := httptest.NewRequest("POST", "/validation", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
