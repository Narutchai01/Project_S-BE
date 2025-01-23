package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type RecoveryRepository interface {
	CreateRecovery(recovery entities.Recovery) (entities.Recovery, error)
	DeleteRecoveryById(id int) (entities.Recovery, error)
	GetRecoveries() ([]entities.Recovery, error)
	GetRecoveryById(id int) (entities.Recovery, error)
	GetRecoveryByUserId(user_id int) (entities.Recovery, error)
	UpdateRecoveryOtpById(id int, otp string) (entities.Recovery, error)
}