package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type RecoveryRepository interface {
	CreateRecovery(recovery entities.Recovery) (entities.Recovery, error)
	DeleteRecovery(id int) (entities.Recovery, error)
}