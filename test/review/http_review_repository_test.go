package adapters_test

import (
	"bytes"
	"mime/multipart"
	"net/http/httptest"
	"testing"

	adapters "github.com/Narutchai01/Project_S-BE/adapters/review"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockReviewUseCase struct {
	mock.Mock
}

func (m *MockReviewUseCase) CreateReviewSkincare(review entities.ReviewSkincare, token string, file multipart.FileHeader, c *fiber.Ctx) (entities.ReviewSkincare, error) {
	args := m.Called(review, token, file, c)
	return args.Get(0).(entities.ReviewSkincare), args.Error(1)
}

func (m *MockReviewUseCase) GetReviewSkincare(review_id uint, token string) (entities.ReviewSkincare, error) {
	args := m.Called(review_id, token)
	return args.Get(0).(entities.ReviewSkincare), args.Error(1)
}

func (m *MockReviewUseCase) GetReviewSkincares(token string) ([]entities.ReviewSkincare, error) {
	args := m.Called(token)
	return args.Get(0).([]entities.ReviewSkincare), args.Error(1)
}

func TestCreateReviewSkincare(t *testing.T) {
	app := fiber.New()
	mockUseCase := new(MockReviewUseCase)
	repo := adapters.NewHttpReviewRepository(mockUseCase)
	app.Post("/review", repo.CreateReviewSkincare)

	expectData := entities.ReviewSkincare{
		Model:      gorm.Model{ID: 1},
		UserID:     1,
		SkincareID: []int{1, 2},
		Content:    "test content",
		Image:      "test.jpg",
		Title:      "test title",
	}

	t.Run("CreateReviewSkincare_Success", func(t *testing.T) {
		mockUseCase.On("CreateReviewSkincare", mock.Anything, "token", mock.AnythingOfType("multipart.FileHeader"), mock.Anything).Return(expectData, nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("title", "test title")
		writer.WriteField("content", "test content")
		writer.WriteField("skincare_id", "[1,2]")
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("file content"))
		writer.Close()

		req := httptest.NewRequest("POST", "/review", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "token")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("CreateReviewSkincare_Unauthorized", func(t *testing.T) {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("title", "test title")
		writer.WriteField("content", "test content")
		writer.WriteField("skincare_id", "[1,2]")
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("file content"))
		writer.Close()

		req := httptest.NewRequest("POST", "/review", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("CreateReviewSkincare_BadRequest", func(t *testing.T) {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("title", "test title")
		writer.WriteField("content", "test content")
		writer.WriteField("skincare_id", "invalid_json")
		writer.Close()

		req := httptest.NewRequest("POST", "/review", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("token", "token")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}
func TestGetReviewSkincare(t *testing.T) {
	app := fiber.New()
	mockUseCase := new(MockReviewUseCase)
	repo := adapters.NewHttpReviewRepository(mockUseCase)
	app.Get("/review/:id", repo.GetReviewSkincare)

	expectData := entities.ReviewSkincare{
		Model:      gorm.Model{ID: 1},
		UserID:     1,
		SkincareID: []int{1, 2},
		Content:    "test content",
		Image:      "test.jpg",
		Title:      "test title",
	}

	t.Run("GetReviewSkincare_Success", func(t *testing.T) {
		mockUseCase.On("GetReviewSkincare", uint(1), "token").Return(expectData, nil)

		req := httptest.NewRequest("GET", "/review/1", nil)
		req.Header.Set("token", "token")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("GetReviewSkincare_Unauthorized", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/review/1", nil)

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("GetReviewSkincare_BadRequest", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/review/invalid_id", nil)
		req.Header.Set("token", "token")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

}
