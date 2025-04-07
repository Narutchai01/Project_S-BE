package usecases

import (
	"bytes"
	"encoding/json"
	"errors"
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
	GetResult(id uint) (entities.Result, error)
	GetResults(token string) ([]entities.Result, error)
	GetResultLatest(token string) (entities.Result, error)
	GetResultByIDs(ids []uint, token string) ([]entities.Result, error)
}

type resultService struct {
	repo     repositories.ResultsRepository
	userRepo repositories.UserRepository
}

func NewResultsUsecase(repo repositories.ResultsRepository, userRepo repositories.UserRepository) ResultsUsecase {
	return &resultService{repo, userRepo}
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
		return entities.Result{}, err
	}

	if resq.StatusCode != http.StatusOK {
		defer resq.Body.Close()
		return entities.Result{}, fmt.Errorf("API request failed with status: %s", resq.Status)
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

	_ = c.SaveFile(&file, "./uploads/"+fileName)

	imageUrl, err := utils.UploadImage(fileName, "/results")
	if err != nil {
		return entities.Result{}, err
	}

	_ = os.Remove("./uploads/" + fileName)

	createBy, err := utils.ExtractToken(token)
	if err != nil {
		return entities.Result{}, err
	}

	api_model := os.Getenv("API_MODEL")

	data, err := CallAPI(api_model+"/predict", imageUrl, createBy)
	if err != nil {
		return entities.Result{}, err
	}

	data.Image = imageUrl

	result, err := service.repo.CreateResult(data)
	if err != nil {
		return entities.Result{}, err
	}

	if len(data.SkincareID) > 0 {
		for _, skincare := range data.SkincareID {
			skincareResult := entities.SkincareResult{
				SkincareID: skincare,
				ResultID:   result.ID,
			}

			_, err := service.repo.CreateSkincareResult(skincareResult)
			if err != nil {
				return entities.Result{}, err
			}
		}
	}

	result, err = service.repo.GetResult(result.ID)
	if err != nil {
		return entities.Result{}, err
	}

	return result, nil

}

func (service *resultService) GetResult(id uint) (entities.Result, error) {
	result, err := service.repo.GetResult(id)
	if err != nil {
		return entities.Result{}, err
	}

	return result, nil
}

func (service *resultService) GetResults(token string) ([]entities.Result, error) {
	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return nil, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return nil, err
	}

	results, err := service.repo.GetResults(user.ID)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (service *resultService) GetResultLatest(token string) (entities.Result, error) {
	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.Result{}, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.Result{}, err
	}

	result, err := service.repo.GetResults(user.ID)
	if err != nil {
		return entities.Result{}, err
	}

	if len(result) < 1 {
		return entities.Result{}, errors.New("result not found")
	}

	return result[len(result)-1], nil
}

func (service *resultService) GetResultByIDs(ids []uint, token string) ([]entities.Result, error) {
	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return []entities.Result{}, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return []entities.Result{}, err
	}

	results, err := service.repo.GetResults(user.ID)
	if err != nil {
		return []entities.Result{}, err
	}

	filteredResults := make([]entities.Result, 0)
	for _, result := range results {
		for _, id := range ids {
			if result.ID == id {
				filteredResults = append(filteredResults, result)
				break
			}
		}
	}

	if len(filteredResults) < len(ids) {
		return []entities.Result{}, errors.New("results not found")
	}

	return filteredResults, nil
}
