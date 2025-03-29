package adapters_test

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http/httptest"
	"testing"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/skin"
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

func TestCreateSkin(t *testing.T) {
	setup := func() (*MockFaceProblemUseCase, *adapters.HttpSkinHandler, *fiber.App) {
		mockService := new(MockFaceProblemUseCase)
		handler := adapters.NewHttpSkinHandler(mockService)

		app := fiber.New()
		app.Post("/admin/skin", handler.CreateSkin)

		return mockService, handler, app
	}

	expectData := entities.FaceProblem{
		Name:      "skin_type1",
		Image:     "skin/type1/path",
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

		req := httptest.NewRequest("POST", "/admin/skin", body)
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

		req := httptest.NewRequest("POST", "/admin/skin", body)
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

		req := httptest.NewRequest("POST", "/admin/skin", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("bad_request_invalid_body", func(t *testing.T) {
		_, _, app := setup()

		// Invalid JSON body
		req := httptest.NewRequest("POST", "/admin/skin", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", "example-token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}
func TestGetSkins(t *testing.T) {
	setup := func() (*MockFaceProblemUseCase, *adapters.HttpSkinHandler, *fiber.App) {
		mockService := new(MockFaceProblemUseCase)
		handler := adapters.NewHttpSkinHandler(mockService)

		app := fiber.New()
		app.Get("/skin", handler.GetSkins)

		return mockService, handler, app
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()

		expectedSkins := []entities.FaceProblem{
			{Model: gorm.Model{ID: 1}, Name: "skin_type1", Image: "skin/type1/path"},
			{Model: gorm.Model{ID: 2}, Name: "skin_type2", Image: "skin/type2/path"},
		}

		mockService.On("GetProblems", "skin").Return(expectedSkins, nil)

		req := httptest.NewRequest("GET", "/skin", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("GetProblems", "skin").Return([]entities.FaceProblem{}, errors.New("database error"))

		req := httptest.NewRequest("GET", "/skin", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
func TestGetSkin(t *testing.T) {
	setup := func() (*MockFaceProblemUseCase, *adapters.HttpSkinHandler, *fiber.App) {
		mockService := new(MockFaceProblemUseCase)
		handler := adapters.NewHttpSkinHandler(mockService)

		app := fiber.New()
		app.Get("/skin/:id", handler.GetSkin)

		return mockService, handler, app
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()

		expectedSkin := entities.FaceProblem{
			Model: gorm.Model{ID: 1},
			Name:  "skin_type1",
			Image: "skin/type1/path",
		}

		mockService.On("GetProblem", uint64(1)).Return(expectedSkin, nil)

		req := httptest.NewRequest("GET", "/skin/1", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("invalid_id", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest("GET", "/skin/invalid", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("not_found", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("GetProblem", uint64(999)).Return(entities.FaceProblem{}, errors.New("skin not found"))

		req := httptest.NewRequest("GET", "/skin/999", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
func TestUpdateSkin(t *testing.T) {
	setup := func() (*MockFaceProblemUseCase, *adapters.HttpSkinHandler, *fiber.App) {
		mockService := new(MockFaceProblemUseCase)
		handler := adapters.NewHttpSkinHandler(mockService)

		app := fiber.New()
		app.Put("/admin/skin/:id", handler.UpdateSkin)

		return mockService, handler, app
	}

	t.Run("success_with_file", func(t *testing.T) {
		mockService, _, app := setup()

		updatedSkin := entities.FaceProblem{
			Model: gorm.Model{ID: 1},
			Name:  "updated_skin_type",
			Image: "skin/updated/path",
		}

		mockService.On("UpdateFaceProblems",
			1,
			mock.Anything,
			mock.AnythingOfType("*multipart.FileHeader"),
			mock.Anything,
		).Return(updatedSkin, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", updatedSkin.Name)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("test image"))
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/skin/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("success_without_file", func(t *testing.T) {
		mockService, _, app := setup()

		updatedSkin := entities.FaceProblem{
			Model: gorm.Model{ID: 1},
			Name:  "updated_skin_type",
			Image: "skin/updated/path",
		}

		mockService.On("UpdateFaceProblems",
			1,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(updatedSkin, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", updatedSkin.Name)
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/skin/1", body)
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
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/skin/invalid", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("invalid_body", func(t *testing.T) {
		_, _, app := setup()

		// Invalid JSON body
		req := httptest.NewRequest("PUT", "/admin/skin/1", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("skin_not_found", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("UpdateFaceProblems",
			999,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(entities.FaceProblem{}, errors.New("skin not found"))

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", "any_name")
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/skin/999", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("server_error", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("UpdateFaceProblems",
			1,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(entities.FaceProblem{}, errors.New("database error"))

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("name", "any_name")
		writer.Close()

		req := httptest.NewRequest("PUT", "/admin/skin/1", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
func TestDeleteSkin(t *testing.T) {
	setup := func() (*MockFaceProblemUseCase, *adapters.HttpSkinHandler, *fiber.App) {
		mockService := new(MockFaceProblemUseCase)
		handler := adapters.NewHttpSkinHandler(mockService)

		app := fiber.New()
		app.Delete("/admin/skin/:id", handler.DeleteSkin)

		return mockService, handler, app
	}

	t.Run("success", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("DeleteFaceProblem", 1).Return(nil)

		req := httptest.NewRequest("DELETE", "/admin/skin/1", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("invalid_id", func(t *testing.T) {
		_, _, app := setup()

		req := httptest.NewRequest("DELETE", "/admin/skin/invalid", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("skin_not_found", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("DeleteFaceProblem", 999).Return(errors.New("skin not found"))

		req := httptest.NewRequest("DELETE", "/admin/skin/999", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("server_error", func(t *testing.T) {
		mockService, _, app := setup()

		mockService.On("DeleteFaceProblem", 1).Return(errors.New("database error"))

		req := httptest.NewRequest("DELETE", "/admin/skin/1", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
