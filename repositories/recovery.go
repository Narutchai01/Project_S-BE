package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type RecoveryRepository interface {
	CreateRecovery(recovery entities.Recovery) (entities.Recovery, error)
	FindRecoveryByOTP(otp string, user_id uint) (entities.Recovery, error)
	DeleteRecoveryById(id uint) error
}
