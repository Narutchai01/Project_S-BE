package usecases

import (
	"fmt"
	"mime/multipart"
	"os"
	"strconv"
	

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ResultUsecases interface {
	CreateResult(token string, file multipart.FileHeader, c *fiber.Ctx) (entities.Result, error)
	GetResults() ([]entities.Result, error)
	GetResultById(id int) (entities.Result, error)
	UpdateResultById(id int, result entities.Result) (entities.Result, error)
	DeleteResultById(id int) error
	GetResultsByUserIdFromToken(token string) ([]entities.Result, error)
	GetResultsByUserId(id int) ([]entities.Result, error)
	GetLatestResultByUserIdFromToken(token string) (entities.Result, error)
}
type resultService struct {
	repo repositories.ResultRepository
}

func NewResultUsecase(repo repositories.ResultRepository) ResultUsecases {
	return &resultService{repo}
}

func (service *resultService) CreateResult(token string, file multipart.FileHeader, c *fiber.Ctx) (entities.Result, error) {
	var result entities.Result

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.Result{}, err
	}

	fileName := uuid.New().String() + ".jpg"

	if err := utils.CheckDirectoryExist(); err != nil {
		return entities.Result{}, err
	}

	if err := c.SaveFile(&file, "./uploads/"+fileName); err != nil {
		return entities.Result{}, err
	}

	imageUrl, err := utils.UploadImage(fileName, fmt.Sprintf("/result/%s", strconv.Itoa(int(user_id))))
	if err != nil {
		return result, c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	err = os.Remove("./uploads/" + fileName)
	if err != nil {
		return result, c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	analyzeResult, err := utils.FacialAcneSkinAnalysis("g3jhJHAwYUYatCbvFFLK", "ucare/1", imageUrl)
	if err != nil {
		return result, c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	result.Image = imageUrl
	result.AcneType = analyzeResult.AcneType
	result.FacialType = analyzeResult.FacialType
	result.SkinType = analyzeResult.SkinType
	result.Skincare = analyzeResult.Skincare
	result.UserId = uint(user_id)

	return service.repo.CreateResult(result)
}

func (service *resultService) GetResults() ([]entities.Result, error) {
	return service.repo.GetResults()
}

func (service *resultService) GetResultById(id int) (entities.Result, error) {
	return service.repo.GetResultById(id)
}

func (service *resultService) UpdateResultById(id int, result entities.Result) (entities.Result, error) {

	old_result, err := service.repo.GetResultById(id)
	if err != nil {
		return entities.Result{}, err
	}

	old_result.Image = utils.CheckEmptyValueBeforeUpdate(result.Image, old_result.Image)
	if result.UserId == 0 {
		result.UserId = old_result.UserId
	}
	if len(result.AcneType) <= 0 {
		result.AcneType = old_result.AcneType
	}
	if len(result.FacialType) <= 0 {
		result.FacialType = old_result.FacialType
	}
	if result.SkinType == 0 {
		result.SkinType = old_result.SkinType
	}
	if len(result.Skincare) <= 0 {
		result.Skincare = old_result.Skincare
		// !reflect.DeepEqual(result.Skincare, old_result.Skincare)
	}

	return service.repo.UpdateResultById(id, result)
}

func (service *resultService) DeleteResultById(id int) error {
	return service.repo.DeleteResultById(id)
}

func (service *resultService) GetResultsByUserIdFromToken(token string) ([]entities.Result, error) {
	user_id, err := utils.ExtractToken(token)

	if err != nil {
		return []entities.Result{}, err
	}
	return service.repo.GetResultsByUserId(int(user_id))
}

func (service *resultService) GetResultsByUserId(user_id int) ([]entities.Result, error) {
	return service.repo.GetResultsByUserId(user_id)
}

func (service *resultService) GetLatestResultByUserIdFromToken(token string) (entities.Result, error) {
	user_id, err := utils.ExtractToken(token)

	if err != nil {
		return entities.Result{}, err
	}
	return service.repo.GetLatestResultByUserIdFromToken(int(user_id))
}
