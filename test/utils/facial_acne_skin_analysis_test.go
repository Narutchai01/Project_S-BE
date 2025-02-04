package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FacialAcneSkinAnalysisMock struct {
	mock.Mock
}

type FacialAcneSkinAnalysisService interface {
	Analyze(apiKey string, modelID string, imageURL string) (string, error)
}

func (m *FacialAcneSkinAnalysisMock) Analyze(apiKey string, modelID string, imageURL string) (string, error) {
	args := m.Called(apiKey, modelID, imageURL)
	return args.String(0), args.Error(1)
}

func TestFacialAcneSkinAnalysisMock(t *testing.T) {
	mockService := new(FacialAcneSkinAnalysisMock)

	// กำหนดค่า Mock Response
	mockService.On("Analyze", "fake_api_key", "fake_model", "http://example.com/image.jpg").
		Return("{\"result\":\"Clear Skin\"}", nil)

	result, err := mockService.Analyze("fake_api_key", "fake_model", "http://example.com/image.jpg")

	assert.Nil(t, err)
	assert.JSONEq(t, "{\"result\":\"Clear Skin\"}", result)
	mockService.AssertCalled(t, "Analyze", "fake_api_key", "fake_model", "http://example.com/image.jpg")
}
