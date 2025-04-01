package usecases

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
	"golang.org/x/crypto/bcrypt"
)

type RecoveryUsecases interface {
	CreateRecovery(email string) (entities.Recovery, error)
	ValidateRecovery(otp string, user_id uint) (entities.Recovery, error)
	ResetPassword(newPassword string, user_id uint) (entities.User, error)
}
type recoveryService struct {
	repo     repositories.RecoveryRepository
	userRepo repositories.UserRepository
}

func NewRecoveryUseCase(repo repositories.RecoveryRepository, userRepo repositories.UserRepository) RecoveryUsecases {
	return &recoveryService{repo, userRepo}
}

func GenerateOTP(length int) string {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	otp := ""
	for i := 0; i < length; i++ {
		otp += fmt.Sprintf("%d", rand.Intn(10)) // Generate a random digit (0-9)
	}
	return otp
}

func (service *recoveryService) CreateRecovery(email string) (entities.Recovery, error) {

	user, err := service.userRepo.GetUserByEmail(email)
	if err != nil {
		return entities.Recovery{}, errors.New("user not found")
	}

	otp := GenerateOTP(6)

	recovery := entities.Recovery{
		OTP:    otp,
		UserID: uint(user.ID),
	}

	if err := utils.SendEmailVerification(email, otp); err != nil {
		return entities.Recovery{}, errors.New("failed to send email")
	}

	return service.repo.CreateRecovery(recovery)

}

func (service *recoveryService) ValidateRecovery(otp string, user_id uint) (entities.Recovery, error) {

	recovery, err := service.repo.FindRecoveryByOTP(otp, user_id)
	if err != nil {
		return entities.Recovery{}, errors.New("invalid OTP")
	}

	err = service.repo.DeleteRecoveryById(recovery.ID)
	if err != nil {
		return entities.Recovery{}, errors.New("failed to delete recovery record")
	}

	return recovery, nil
}

func (service *recoveryService) ResetPassword(newPassword string, user_id uint) (entities.User, error) {
	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.User{}, errors.New("user not found")
	}

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return entities.User{}, errors.New("failed to hash password")
	}

	user.Password = string(hashedNewPassword)

	return service.userRepo.UpdateUser(user)

}
