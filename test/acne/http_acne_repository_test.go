package adapters_test

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http/httptest"
	"testing"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/acne"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockFaceProblemUseCase struct {
	mock.Mock
}

// CreateProblem implements usecases.FaceProblemUseCase.
func (m *MockFaceProblemUseCase) CreateProblem(problem entities.FaceProblem, file multipart.FileHeader, c *fiber.Ctx, token string, type_problem string) (entities.FaceProblem, error) {
	args := m.Called(problem, file, c, token, type_problem)
	return args.Get(0).(entities.FaceProblem), args.Error(1)
}

// DeleteFaceProblem implements usecases.FaceProblemUseCase.
func (m *MockFaceProblemUseCase) DeleteFaceProblem(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// GetProblem implements usecases.FaceProblemUseCase.
func (m *MockFaceProblemUseCase) GetProblem(id uint64) (entities.FaceProblem, error) {
	args := m.Called(id)
	return args.Get(0).(entities.FaceProblem), args.Error(1)
}

// GetProblems implements usecases.FaceProblemUseCase.
func (m *MockFaceProblemUseCase) GetProblems(type_problem string) ([]entities.FaceProblem, error) {
	args := m.Called(type_problem)
	return args.Get(0).([]entities.FaceProblem), args.Error(1)
}

// UpdateFaceProblems implements usecases.FaceProblemUseCase.
func (m *MockFaceProblemUseCase) UpdateFaceProblems(id int, problem entities.FaceProblem, file *multipart.FileHeader, c *fiber.Ctx) (entities.FaceProblem, error) {
	args := m.Called(id, problem, file, c)
	return args.Get(0).(entities.FaceProblem), args.Error(1)
}

func TestCreateAcneHandler(t *testing.T) {
	setup := func() (*MockFaceProblemUseCase, *adapters.HttpAcneHandler, *fiber.App) {
		mockService := new(MockFaceProblemUseCase)
		handler := adapters.NewHttpAcneHandler(mockService)

		app := fiber.New()
		app.Post("/admin/acne", handler.CreateAcne)

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

		req := httptest.NewRequest("POST", "/admin/acne", body)
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

		req := httptest.NewRequest("POST", "/admin/acne", body)
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

		req := httptest.NewRequest("POST", "/admin/acne", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("bad_request_invalid_body", func(t *testing.T) {
		_, _, app := setup()

		// Invalid JSON body
		req := httptest.NewRequest("POST", "/admin/acne", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}
func TestGetAcnesHandler(t *testing.T) {
	setup := func() (*MockFaceProblemUseCase, *adapters.HttpAcneHandler, *fiber.App) {
		mockService := new(MockFaceProblemUseCase)
		handler := adapters.NewHttpAcneHandler(mockService)

		app := fiber.New()
		app.Get("/acne", handler.GetAcnes)

		return mockService, handler, app
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()

		expectedData := []entities.FaceProblem{
			{
				Model:     gorm.Model{ID: 1},
				Name:      "acne_type1",
				Image:     "acne/type1/path",
				CreatedBy: 1,
			},
			{
				Model:     gorm.Model{ID: 2},
				Name:      "acne_type2",
				Image:     "acne/type2/path",
				CreatedBy: 1,
			},
		}
		mockService.On("GetProblems", "acne").Return(expectedData, nil)
		req := httptest.NewRequest("GET", "/acne", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("error from service", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("GetProblems", "acne").Return([]entities.FaceProblem{}, fiber.ErrInternalServerError)
		req := httptest.NewRequest("GET", "/acne", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("empty result", func(t *testing.T) {
		mockService, _, app := setup()

		// Return empty slice but no error
		mockService.On("GetProblems", "acne").Return([]entities.FaceProblem{}, nil)
		req := httptest.NewRequest("GET", "/acne", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("with query parameter", func(t *testing.T) {
		mockService, _, app := setup()

		expectedData := []entities.FaceProblem{
			{
				Model:     gorm.Model{ID: 1},
				Name:      "filtered_acne",
				Image:     "acne/filtered/path",
				CreatedBy: 1,
			},
		}

		// You might need to adjust this depending on how your handler processes query parameters
		mockService.On("GetProblems", "acne").Return(expectedData, nil)
		req := httptest.NewRequest("GET", "/acne?name=filtered", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
func TestGetAcneHandler(t *testing.T) {
	setup := func() (*MockFaceProblemUseCase, *adapters.HttpAcneHandler, *fiber.App) {
		mockService := new(MockFaceProblemUseCase)
		handler := adapters.NewHttpAcneHandler(mockService)

		app := fiber.New()
		app.Get("/acne/:id", handler.GetAcne)

		return mockService, handler, app
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()

		expectedData := entities.FaceProblem{
			Model:     gorm.Model{ID: 1},
			Name:      "acne_type1",
			Image:     "acne/type1/path",
			CreatedBy: 1,
		}
		mockService.On("GetProblem", uint64(1)).Return(expectedData, nil)

		req := httptest.NewRequest("GET", "/acne/1", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("invalid_id", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest("GET", "/acne/invalid", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("not_found", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("GetProblem", uint64(999)).Return(entities.FaceProblem{}, fiber.NewError(fiber.StatusNotFound, "acne not found"))

		req := httptest.NewRequest("GET", "/acne/999", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("zero_id", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("GetProblem", uint64(0)).Return(entities.FaceProblem{}, fiber.NewError(fiber.StatusNotFound, "acne not found"))

		req := httptest.NewRequest("GET", "/acne/0", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
func TestUpdateAcneHandler(t *testing.T) {
	setup := func() (*MockFaceProblemUseCase, *adapters.HttpAcneHandler, *fiber.App) {
		mockService := new(MockFaceProblemUseCase)
		handler := adapters.NewHttpAcneHandler(mockService)

		app := fiber.New()
		app.Put("/admin/acne/:id", handler.UpdateAcne)

		return mockService, handler, app
	}

	expectData := entities.FaceProblem{
		Model:     gorm.Model{ID: 1},
		Name:      "updated_acne",
		Image:     "acne/updated/path",
		CreatedBy: 1,
	}

	t.Run("success_with_file", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateFaceProblems",
			1,
			mock.AnythingOfType("entities.FaceProblem"),
			mock.AnythingOfType("*multipart.FileHeader"),
			mock.Anything,
		).Return(expectData, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/acne/1", body)
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
			mock.MatchedBy(func(problem entities.FaceProblem) bool {
				return problem.Name == expectData.Name
			}),
			mock.Anything, // Changed from nil to mock.Anything because the handler might pass a nil pointer
			mock.Anything,
		).Return(expectData, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/acne/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("invalid_id", func(t *testing.T) {
		_, _, app := setup()

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/acne/invalid", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("not_found", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("UpdateFaceProblems",
			999,
			mock.AnythingOfType("entities.FaceProblem"),
			mock.AnythingOfType("*multipart.FileHeader"),
			mock.Anything,
		).Return(entities.FaceProblem{}, fiber.NewError(fiber.StatusNotFound, "acne not found"))

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/acne/999", body)
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
			mock.AnythingOfType("entities.FaceProblem"),
			mock.AnythingOfType("*multipart.FileHeader"),
			mock.Anything,
		).Return(entities.FaceProblem{}, fiber.NewError(fiber.StatusInternalServerError, "internal server error"))

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", expectData.Name)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/acne/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("invalid_body", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest("PUT", "/admin/acne/1", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}
func TestDeleteAcneHandler(t *testing.T) {
	setup := func() (*MockFaceProblemUseCase, *adapters.HttpAcneHandler, *fiber.App) {
		mockService := new(MockFaceProblemUseCase)
		handler := adapters.NewHttpAcneHandler(mockService)

		app := fiber.New()
		app.Delete("/admin/acne/:id", handler.DeleteAcne)

		return mockService, handler, app
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("DeleteFaceProblem", 1).Return(nil)

		req := httptest.NewRequest("DELETE", "/admin/acne/1", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("invalid_id", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest("DELETE", "/admin/acne/invalid", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("not_found", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("DeleteFaceProblem", 999).Return(errors.New("acne not found"))

		req := httptest.NewRequest("DELETE", "/admin/acne/999", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("internal_server_error", func(t *testing.T) {
		mockService, _, app := setup()
		mockService.On("DeleteFaceProblem", 1).Return(errors.New("database error"))

		req := httptest.NewRequest("DELETE", "/admin/acne/1", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
