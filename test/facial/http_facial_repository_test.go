package adapters_test

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http/httptest"
	"testing"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/facial"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockFaceProblemUseCase struct {
	mock.Mock
}

func (m *MockFaceProblemUseCase) CreateProblem(problem entities.FaceProblem, file multipart.FileHeader, c *fiber.Ctx, token string, type_problem string) (entities.FaceProblem, error) {
	args := m.Called(problem, file, c, token, type_problem)
	return args.Get(0).(entities.FaceProblem), args.Error(1)
}

func (m *MockFaceProblemUseCase) DeleteFaceProblem(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockFaceProblemUseCase) GetProblem(id uint64) (entities.FaceProblem, error) {
	args := m.Called(id)
	return args.Get(0).(entities.FaceProblem), args.Error(1)
}

func (m *MockFaceProblemUseCase) GetProblems(type_problem string) ([]entities.FaceProblem, error) {
	args := m.Called(type_problem)
	return args.Get(0).([]entities.FaceProblem), args.Error(1)
}

func (m *MockFaceProblemUseCase) UpdateFaceProblems(id int, problem entities.FaceProblem, file *multipart.FileHeader, c *fiber.Ctx) (entities.FaceProblem, error) {
	args := m.Called(id, problem, file, c)
	return args.Get(0).(entities.FaceProblem), args.Error(1)
}

func TestCreateFacialHandler(t *testing.T) {
	setup := func() (*MockFaceProblemUseCase, *adapters.HttpFacialHandler, *fiber.App) {
		mockService := new(MockFaceProblemUseCase)
		handler := adapters.NewHttpFacialHandler(mockService)

		app := fiber.New()
		app.Post("/admin/facial", handler.CreateFacial)

		return mockService, handler, app
	}

	expectData := entities.FaceProblem{
		Name:      "facial_type1",
		Image:     "facial/type1/path",
		CreatedBy: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("CreateProblem",
			mock.Anything,
			mock.AnythingOfType("multipart.FileHeader"),
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(expectData, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("POST", "/admin/facial", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("bad_request_missing_file", func(t *testing.T) {
		_, _, app := setup()

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		writer.Close()

		req := httptest.NewRequest("POST", "/admin/facial", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("bad_request_empty_name", func(t *testing.T) {
		_, _, app := setup()

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", "")
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("POST", "/admin/facial", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("bad_request_invalid_body", func(t *testing.T) {
		_, _, app := setup()

		// Invalid JSON body
		req := httptest.NewRequest("POST", "/admin/facial", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}
func TestGetFacialsHandler(t *testing.T) {
	setup := func() (*MockFaceProblemUseCase, *adapters.HttpFacialHandler, *fiber.App) {
		mockService := new(MockFaceProblemUseCase)
		handler := adapters.NewHttpFacialHandler(mockService)

		app := fiber.New()
		app.Get("/facial", handler.GetFacials)

		return mockService, handler, app
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()

		facials := []entities.FaceProblem{
			{
				Model:     gorm.Model{ID: 1},
				Name:      "facial_type1",
				Image:     "facial/type1/path",
				CreatedBy: 1,
			},
			{
				Model:     gorm.Model{ID: 2},
				Name:      "facial_type2",
				Image:     "facial/type2/path",
				CreatedBy: 1,
			},
		}

		mockService.On("GetProblems", "facial").Return(facials, nil)

		req := httptest.NewRequest("GET", "/facial", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("internal_server_error", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("GetProblems", "facial").Return([]entities.FaceProblem{}, errors.New("database error"))

		req := httptest.NewRequest("GET", "/facial", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
func TestGetFacialHandler(t *testing.T) {
	setup := func() (*MockFaceProblemUseCase, *adapters.HttpFacialHandler, *fiber.App) {
		mockService := new(MockFaceProblemUseCase)
		handler := adapters.NewHttpFacialHandler(mockService)

		app := fiber.New()
		app.Get("/facial/:id", handler.GetFacial)

		return mockService, handler, app
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()

		facial := entities.FaceProblem{
			Model:     gorm.Model{ID: 1},
			Name:      "facial_type1",
			Image:     "facial/type1/path",
			CreatedBy: 1,
		}

		mockService.On("GetProblem", uint64(1)).Return(facial, nil)

		req := httptest.NewRequest("GET", "/facial/1", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("invalid_id_param", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest("GET", "/facial/invalid", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})

	t.Run("facial_not_found", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("GetProblem", uint64(999)).Return(entities.FaceProblem{}, errors.New("facial not found"))

		req := httptest.NewRequest("GET", "/facial/999", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
func TestUpdateFacialHandler(t *testing.T) {
	setup := func() (*MockFaceProblemUseCase, *adapters.HttpFacialHandler, *fiber.App) {
		mockService := new(MockFaceProblemUseCase)
		handler := adapters.NewHttpFacialHandler(mockService)

		app := fiber.New()
		app.Put("/admin/facial/:id", handler.UpdateFacial)

		return mockService, handler, app
	}

	expectData := entities.FaceProblem{
		Model:     gorm.Model{ID: 1},
		Name:      "updated_facial_type",
		Image:     "facial/updated/path",
		CreatedBy: 1,
	}

	t.Run("success_with_file", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("UpdateFaceProblems",
			1,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(expectData, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/facial/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("success_without_file", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("UpdateFaceProblems",
			1,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(expectData, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/facial/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("invalid_id_param", func(t *testing.T) {
		_, _, app := setup()

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/facial/invalid", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("bad_request_invalid_body", func(t *testing.T) {
		_, _, app := setup()

		// Invalid body
		req := httptest.NewRequest("PUT", "/admin/facial/1", bytes.NewBuffer([]byte("invalid body")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("facial_not_found", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("UpdateFaceProblems",
			1,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(entities.FaceProblem{}, errors.New("facial not found"))

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/facial/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("internal_server_error", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("UpdateFaceProblems",
			1,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(entities.FaceProblem{}, errors.New("database error"))

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/facial/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
func TestDeleteFacialHandler(t *testing.T) {
	setup := func() (*MockFaceProblemUseCase, *adapters.HttpFacialHandler, *fiber.App) {
		mockService := new(MockFaceProblemUseCase)
		handler := adapters.NewHttpFacialHandler(mockService)

		app := fiber.New()
		app.Delete("/admin/facial/:id", handler.DeleteFacial)

		return mockService, handler, app
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("DeleteFaceProblem", 1).Return(nil)

		req := httptest.NewRequest("DELETE", "/admin/facial/1", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("invalid_id_param", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest("DELETE", "/admin/facial/invalid", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("facial_not_found", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("DeleteFaceProblem", 999).Return(errors.New("facial not found"))

		req := httptest.NewRequest("DELETE", "/admin/facial/999", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("internal_server_error", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("DeleteFaceProblem", 1).Return(errors.New("database error"))

		req := httptest.NewRequest("DELETE", "/admin/facial/1", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
