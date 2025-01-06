package usecases

import (
      "github.com/Narutchai01/Project_S-BE/entities"
      "github.com/Narutchai01/Project_S-BE/repositories"
      "github.com/Narutchai01/Project_S-BE/utils"
      "github.com/gofiber/fiber/v2"
)

type RecoveryUsecases interface {
      CreateRecovery(recovery entities.Recovery, c *fiber.Ctx) (entities.Recovery, error)
}

type recoveryService struct {
      repo repositories.RecoveryRepository
}

func NewRecoveryUseCase(repo repositories.RecoveryRepository) RecoveryUsecases {
      return &recoveryService{repo}
}

func (service *recoveryService) CreateRecovery(recovery entities.Recovery, c *fiber.Ctx) (entities.Recovery, error) {


      generateOTP, err := utils.GenerateOTP()

      if err != nil {
            return recovery, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                  "message": err.Error(),
            })
      }

      recovery.OTP = string(generateOTP)

      return service.repo.CreateRecovery(recovery)
}