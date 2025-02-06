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
	CreateResult(token string, file multipart.FileHeader, c *fiber.Ctx) (entities.ResultWithSkincare, error)
	GetResults() ([]entities.ResultWithSkincare, error)
	GetResultById(id int) (entities.ResultWithSkincare, error)
	UpdateResultById(id int, result entities.Result) (entities.ResultWithSkincare, error)
	DeleteResultById(id int) error
	GetResultsByUserIdFromToken(token string) ([]entities.ResultWithSkincare, error)
	GetResultsByUserId(id int) ([]entities.ResultWithSkincare, error)
	GetLatestResultByUserIdFromToken(token string) (entities.ResultWithSkincare, error)
}
type resultService struct {
	repo repositories.ResultRepository
}

func NewResultUsecase(repo repositories.ResultRepository) ResultUsecases {
	return &resultService{repo}
}

func (service *resultService) CreateResult(token string, file multipart.FileHeader, c *fiber.Ctx) (entities.ResultWithSkincare, error) {
	var result entities.Result

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.ResultWithSkincare{}, err
	}

	fileName := uuid.New().String() + ".jpg"

	if err := utils.CheckDirectoryExist(); err != nil {
		return entities.ResultWithSkincare{}, err
	}

	if err := c.SaveFile(&file, "./uploads/"+fileName); err != nil {
		return entities.ResultWithSkincare{}, err
	}

	imageUrl, err := utils.UploadImage(fileName, fmt.Sprintf("/result/%s", strconv.Itoa(int(user_id))))
	if err != nil {
		return entities.ResultWithSkincare{}, c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	err = os.Remove("./uploads/" + fileName)
	if err != nil {
		return entities.ResultWithSkincare{}, c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	// fullURL := fmt.Sprintf("https://detect.roboflow.com/%s", "ucare/1")

	// client := client.New()
	// client.SetParams(map[string]string{
	// 	"api_key": "g3jhJHAwYUYatCbvFFLK",
	// 	"image": imageUrl,
	// })
	// resp, err := client.Post(fullURL)
	// if err != nil {
	// 	return entities.Result{}, err
	// }

	// err = json.Unmarshal(resp.Body(), &result)
	// if err != nil {
	// 	return entities.Result{}, err
	// }

	//mock ไว้ก่อน
	result.AcneType = []entities.Acne_Facial_Result{
		{ID: 1, Count: 10},
		{ID: 2, Count: 5},
	}
	result.FacialType = []entities.Acne_Facial_Result{
		{ID: 1, Count: 10},
		{ID: 2, Count: 5},
	}
	result.SkinType = 1
	result.Skincare = []uint{1,2}

	result.Image = imageUrl
	result.UserId = uint(user_id)

	return service.repo.CreateResult(result)
}

func (service *resultService) GetResults() ([]entities.ResultWithSkincare, error) {
	return service.repo.GetResults()
}

func (service *resultService) GetResultById(id int) (entities.ResultWithSkincare, error) {
	return service.repo.GetResultById(id)
}

func (service *resultService) UpdateResultById(id int, result entities.Result) (entities.ResultWithSkincare, error) {

	old_result, err := service.repo.GetResultByIdWithoutSkincare(id)
	if err != nil {
		return entities.ResultWithSkincare{}, err
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
	}

	return service.repo.UpdateResultById(id, result)
}

func (service *resultService) DeleteResultById(id int) error {
	return service.repo.DeleteResultById(id)
}

func (service *resultService) GetResultsByUserIdFromToken(token string) ([]entities.ResultWithSkincare, error) {
	user_id, err := utils.ExtractToken(token)

	if err != nil {
		return []entities.ResultWithSkincare{}, err
	}
	return service.repo.GetResultsByUserId(int(user_id))
}

func (service *resultService) GetResultsByUserId(user_id int) ([]entities.ResultWithSkincare, error) {
	return service.repo.GetResultsByUserId(user_id)
}

func (service *resultService) GetLatestResultByUserIdFromToken(token string) (entities.ResultWithSkincare, error) {
	user_id, err := utils.ExtractToken(token)

	if err != nil {
		return entities.ResultWithSkincare{}, err
	}
	return service.repo.GetLatestResultByUserIdFromToken(int(user_id))
}
