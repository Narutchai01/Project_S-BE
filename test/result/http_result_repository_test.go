package adapters_test

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/result"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockResultsUsecase struct {
	mock.Mock
}

func (m *MockResultsUsecase) CreateResult(file multipart.FileHeader, token string, c *fiber.Ctx) (entities.Result, error) {
	args := m.Called(file, token, c)
	return args.Get(0).(entities.Result), args.Error(1)
}

func (m *MockResultsUsecase) GetResults(token string) ([]entities.Result, error) {
	args := m.Called(token)
	return args.Get(0).([]entities.Result), args.Error(1)
}

func TestCreateResult(t *testing.T) {
	setup := func() (*MockResultsUsecase, *adapters.HttpResultHandler, *fiber.App) {
		m := new(MockResultsUsecase)
		handler := adapters.NewHttpResultHandler(m)
		app := fiber.New()
		app.Post("/result", handler.CreateResult)
		return m, handler, app
	}

	t.Run("successful creation", func(t *testing.T) {
		m, _, app := setup()

		// Mock the usecase response
		expectedResult := entities.Result{}

		m.On("CreateResult", mock.Anything, "token123", mock.Anything).Return(expectedResult, nil)

		// Create a new file upload request
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("file", "test.txt")
		part.Write([]byte("test content"))
		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/result", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "token123")

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, fiber.StatusCreated, resp.StatusCode)
	})

	t.Run("bad request on missing file", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest(http.MethodPost, "/result", nil)
		req.Header.Set("token", "token123")

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("internal server error on usecase failure", func(t *testing.T) {
		m, _, app := setup()

		// Mock the usecase response
		m.On("CreateResult", mock.Anything, "token123", mock.Anything).Return(entities.Result{}, fiber.ErrInternalServerError)

		// Create a new file upload request
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("file", "test.txt")
		part.Write([]byte("test content"))
		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/result", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "token123")

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})
}
func (m *MockResultsUsecase) GetResultLatest(token string) (entities.Result, error) {
	args := m.Called(token)
	return args.Get(0).(entities.Result), args.Error(1)
}

func TestGetResultLatest(t *testing.T) {
	setup := func() (*MockResultsUsecase, *adapters.HttpResultHandler, *fiber.App) {
		m := new(MockResultsUsecase)
		handler := adapters.NewHttpResultHandler(m)
		app := fiber.New()
		app.Get("/result/latest", handler.GetResultLatest)
		return m, handler, app
	}

	t.Run("successful retrieval", func(t *testing.T) {
		m, _, app := setup()

		// Mock the usecase response
		expectedResult := entities.Result{}

		m.On("GetResultLatest", "token123").Return(expectedResult, nil)

		req := httptest.NewRequest(http.MethodGet, "/result/latest", nil)
		req.Header.Set("token", "token123")

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("bad request on missing token", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest(http.MethodGet, "/result/latest", nil)

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("internal server error on usecase failure", func(t *testing.T) {
		m, _, app := setup()

		// Mock the usecase response
		m.On("GetResultLatest", "token123").Return(entities.Result{}, fiber.ErrInternalServerError)

		req := httptest.NewRequest(http.MethodGet, "/result/latest", nil)
		req.Header.Set("token", "token123")

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})
}
func (m *MockResultsUsecase) GetResult(id uint) (entities.Result, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Result), args.Error(1)
}

func TestGetResult(t *testing.T) {
	setup := func() (*MockResultsUsecase, *adapters.HttpResultHandler, *fiber.App) {
		m := new(MockResultsUsecase)
		handler := adapters.NewHttpResultHandler(m)
		app := fiber.New()
		app.Get("/result/:id", handler.GetResult)
		return m, handler, app
	}

	t.Run("successful retrieval", func(t *testing.T) {
		m, _, app := setup()

		// Mock the usecase response
		expectedResult := entities.Result{}
		m.On("GetResult", uint(1)).Return(expectedResult, nil)

		req := httptest.NewRequest(http.MethodGet, "/result/1", nil)

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("bad request on invalid id", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest(http.MethodGet, "/result/invalid", nil)

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Result not found", func(t *testing.T) {
		m, _, app := setup()

		// Mock the usecase response
		m.On("GetResult", uint(1)).Return(entities.Result{}, fiber.ErrInternalServerError)

		req := httptest.NewRequest(http.MethodGet, "/result/1", nil)

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})
}
func (m *MockResultsUsecase) GetResultByIDs(ids []uint) ([]entities.Result, error) {
	args := m.Called(ids)
	return args.Get(0).([]entities.Result), args.Error(1)
}

func TestGetResultByIDs(t *testing.T) {
	setup := func() (*MockResultsUsecase, *adapters.HttpResultHandler, *fiber.App) {
		m := new(MockResultsUsecase)
		handler := adapters.NewHttpResultHandler(m)
		app := fiber.New()
		app.Post("/results/compare", handler.GetResultByIDs)
		return m, handler, app
	}

	t.Run("successful retrieval", func(t *testing.T) {
		m, _, app := setup()

		// Mock the usecase response
		expectedResults := []entities.Result{}
		m.On("GetResultByIDs", []uint{1, 2, 3}).Return(expectedResults, nil)

		body := bytes.NewBufferString(`{"IDs":[1,2,3]}`)
		req := httptest.NewRequest(http.MethodPost, "/results/compare", body)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "token123")

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("bad request on invalid body", func(t *testing.T) {
		_, _, app := setup()

		body := bytes.NewBufferString(`invalid body`)
		req := httptest.NewRequest(http.MethodPost, "/results/compare", body)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "token123")

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	// t.Run("internal server error on usecase failure", func(t *testing.T) {
	// 	m, _, app := setup()

	// 	// Mock the usecase response
	// 	m.On("GetResultByIDs", []uint{1, 2, 3}).Return(nil, fiber.ErrInternalServerError)

	// 	body := bytes.NewBufferString(`{"IDs":[1,2,3]}`)
	// 	req := httptest.NewRequest(http.MethodPost, "/results/compare", body)
	// 	req.Header.Set("Content-Type", "application/json")
	// 	req.Header.Set("token", "token123")

	// 	resp, err := app.Test(req)
	// 	require.NoError(t, err)
	// 	require.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	// })
}
