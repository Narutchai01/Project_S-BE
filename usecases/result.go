package usecases

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ResultsUsecase interface {
	CreateResult(file multipart.FileHeader, token string, c *fiber.Ctx) (entities.Result, error)
	GetResults(token string) ([]entities.Result, error)
	// GetResult(id uint) (entities.Result, error)
	GetResultLatest(token string) (entities.Result, error)
	// UpdateResult(result entities.Result, id uint) (entities.Result, error)
	// DeleteResult(id uint) error
}

type resultService struct {
	repo repositories.ResultsRepository
}

func NewResultsUsecase(repo repositories.ResultsRepository) ResultsUsecase {
	return &resultService{repo}
}

func CallAPI(url string, image string, id uint) (entities.Result, error) {

	client := http.Client{}

	body := map[string]interface{}{
		"image": image,
		"id":    id,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return entities.Result{}, err
	}

	resq, err := client.Post(url, "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		if resq.StatusCode != http.StatusOK {
			return entities.Result{}, fmt.Errorf("API request failed with status: %s", resq.Status)
		}

		defer resq.Body.Close()
	}

	defer resq.Body.Close()

	var data entities.Result

	if err := json.NewDecoder(resq.Body).Decode(&data); err != nil {
		return entities.Result{}, err
	}

	return data, nil
}

func (service *resultService) CreateResult(file multipart.FileHeader, token string, c *fiber.Ctx) (entities.Result, error) {
	fileName := uuid.New().String() + ".jpg"

	if err := utils.CheckDirectoryExist(); err != nil {
		return entities.Result{}, err
	}

	if err := c.SaveFile(&file, "./uploads/"+fileName); err != nil {
		return entities.Result{}, err
	}

	imageUrl, err := utils.UploadImage(fileName, "/results")
	if err != nil {
		return entities.Result{}, err
	}

	if err := os.Remove("./uploads/" + fileName); err != nil {
		return entities.Result{}, err
	}

	createBy, err := utils.ExtractToken(token)
	if err != nil {
		return entities.Result{}, err
	}

	api_model := os.Getenv("API_MODEL")

	data, err := CallAPI(api_model+"/predict", imageUrl, createBy)
	if err != nil {
		return entities.Result{}, err
	}

	return service.repo.CreateResult(data)
}

func (service *resultService) GetResults(token string) ([]entities.Result, error) {
	id, err := utils.ExtractToken(token)
	if err != nil {
		return nil, err
	}
	return service.repo.GetResults(id)
}

func (service *resultService) GetResultLatest(token string) (entities.Result, error) {
	id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.Result{}, err
	}

	return service.repo.GetResultLatest(id)
}
