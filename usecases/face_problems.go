package usecases

import (
	"errors"
	"mime/multipart"
	"os"
	"strings"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FaceProblemUseCase interface {
	CreateProblem(problem entities.FaceProblem, file multipart.FileHeader, c *fiber.Ctx, token string, type_problem string) (entities.FaceProblem, error)
	GetProblem(id uint64) (entities.FaceProblem, error)
	GetProblems(type_problem string) ([]entities.FaceProblem, error)
	UpdateFaceProblems(id int, problem entities.FaceProblem, file *multipart.FileHeader, c *fiber.Ctx) (entities.FaceProblem, error)
	DeleteFaceProblem(id int) error
}

type faceProblemService struct {
	faceProblemRepo repositories.FaceProblemRepository
}

func NewFaceProblemUseCase(faceProblemRepo repositories.FaceProblemRepository) FaceProblemUseCase {
	return &faceProblemService{faceProblemRepo}
}

func (service *faceProblemService) CreateProblem(problem entities.FaceProblem, file multipart.FileHeader, c *fiber.Ctx, token string, type_problem string) (entities.FaceProblem, error) {

	create_by, err := utils.ExtractToken(token)
	if err != nil {
		return entities.FaceProblem{}, errors.New("Unauthorized")
	}

	problem_type, err := service.faceProblemRepo.GetFaceProblemType(strings.ToLower(type_problem))
	if err != nil {
		return entities.FaceProblem{}, errors.New("problem type not found")
	}

	fileName := uuid.New().String() + ".jpg"

	if err := utils.CheckDirectoryExist(); err != nil {
		return entities.FaceProblem{}, err
	}

	if err := c.SaveFile(&file, "./uploads/"+fileName); err != nil {
		return entities.FaceProblem{}, err
	}

	imageUrl, err := utils.UploadImage(fileName, "/acne")
	if err != nil {
		return entities.FaceProblem{}, errors.New("upload image failed")
	}

	err = os.Remove("./uploads/" + fileName)
	if err != nil {
		return entities.FaceProblem{}, errors.New("delete image failed")
	}

	problem.TypeID = uint64(problem_type.ID)
	problem.CreatedBy = uint64(create_by)
	problem.Image = imageUrl

	return service.faceProblemRepo.CreateFaceProblem(problem)
}

func (service *faceProblemService) GetProblem(id uint64) (entities.FaceProblem, error) {
	return service.faceProblemRepo.GetFaceProblem(id)
}

func (service *faceProblemService) GetProblems(type_problem string) ([]entities.FaceProblem, error) {
	problem_type, err := service.faceProblemRepo.GetFaceProblemType(strings.ToLower(type_problem))
	if err != nil {
		return nil, errors.New("problem type not found")
	}

	return service.faceProblemRepo.GetFaceProblems(uint64(problem_type.ID))
}

func (service *faceProblemService) UpdateFaceProblems(id int, problem entities.FaceProblem, file *multipart.FileHeader, c *fiber.Ctx) (entities.FaceProblem, error) {

	old_value, err := service.faceProblemRepo.GetFaceProblem(uint64(id))
	if err != nil {
		return entities.FaceProblem{}, errors.New("faceproblem not found ")
	}

	if file != nil {
		fileName := uuid.New().String() + ".jpg"

		if err := utils.CheckDirectoryExist(); err != nil {
			return entities.FaceProblem{}, err
		}

		if err := c.SaveFile(file, "./uploads/"+fileName); err != nil {
			return entities.FaceProblem{}, err
		}

		imageUrl, err := utils.UploadImage(fileName, "/acne")
		if err != nil {
			return entities.FaceProblem{}, errors.New("upload image failed")
		}

		err = os.Remove("./uploads/" + fileName)
		if err != nil {
			return entities.FaceProblem{}, errors.New("delete image failed")
		}
		problem.Image = imageUrl
	}

	problem.ID = old_value.ID
	problem.Name = utils.CheckEmptyValueBeforeUpdate(problem.Name, old_value.Name)
	problem.Image = utils.CheckEmptyValueBeforeUpdate(problem.Image, old_value.Image)

	return service.faceProblemRepo.UpdateFaceProblem(uint64(problem.ID), problem)
}

func (service *faceProblemService) DeleteFaceProblem(id int) error {
	_, err := service.faceProblemRepo.GetFaceProblem(uint64(id))
	if err != nil {
		return err
	}

	return service.faceProblemRepo.DeleteFaceProblem(id)
}
