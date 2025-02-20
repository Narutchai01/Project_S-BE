package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type ThreadRepository interface {
	CreateThread(thread []entities.ThreadDetail, user_id uint) (entities.Thread, error)
	GetThreads() ([]entities.Thread, error)
	GetThread(id uint) (entities.Thread, error)
	GetThreadDetails(thread_id uint) ([]entities.ThreadDetail, error)
}
