package usecases

import (
	"fmt"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
	"github.com/gofiber/fiber/v2"
)

type RecoveryUsecases interface {
	CreateRecovery(recovery entities.Recovery, email string, c *fiber.Ctx) (entities.Recovery, error)
	DeleteRecoveryById(id int) (entities.Recovery, error)
	GetRecoveries() ([]entities.Recovery, error)
	GetRecoveryById(id int) (entities.Recovery, error)
	GetRecoveryByUserId(user_id int) (entities.Recovery, error)
	OtpValidation(id int, otp string) (bool, error)
	UpdateRecoveryOtpById(recovery entities.Recovery, email string) (entities.Recovery, error)
}
type recoveryService struct {
	repo repositories.RecoveryRepository
}

func NewRecoveryUseCase(repo repositories.RecoveryRepository) RecoveryUsecases {
	return &recoveryService{repo}
}

func (service *recoveryService) CreateRecovery(recovery entities.Recovery, email string, c *fiber.Ctx) (entities.Recovery, error) {
	generateOTP, err := utils.GenerateOTP()
	if err != nil {
		return recovery, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := utils.SendEmailVerification(email, generateOTP); err != nil {
		return recovery, fmt.Errorf("failed to send email: %w", err)
	}

	recovery.OTP = generateOTP

	return service.repo.CreateRecovery(recovery)
}

func (service *recoveryService) DeleteRecoveryById(id int) (entities.Recovery, error) {
	return service.repo.DeleteRecoveryById(id)
}

func (service *recoveryService) GetRecoveries() ([]entities.Recovery, error) {
	return service.repo.GetRecoveries()
}

func (service *recoveryService) GetRecoveryById(id int) (entities.Recovery, error) {
	return service.repo.GetRecoveryById(id)
}

func (service *recoveryService) GetRecoveryByUserId(user_id int) (entities.Recovery, error) {
	return service.repo.GetRecoveryByUserId(user_id)
}

func (service *recoveryService) OtpValidation(id int, otp string) (bool, error) {
	recovery, err := service.repo.GetRecoveryById(id)
	if err != nil {
		return false, fmt.Errorf("recovery not found: %w", err)
	}

	if recovery.OTP != otp {
		return false, fmt.Errorf("invalid OTP")
	}

	return true, nil
}

func (service *recoveryService) UpdateRecoveryOtpById(recovery entities.Recovery, email string) (entities.Recovery, error) {
	generateOTP, err := utils.GenerateOTP()
	if err != nil {
		return entities.Recovery{}, err
	}

	if err := utils.SendEmailVerification(email, generateOTP); err != nil {
		return recovery, fmt.Errorf("failed to send email: %w", err)
	}

	return service.repo.UpdateRecoveryOtpById(int(recovery.ID), generateOTP)
}

