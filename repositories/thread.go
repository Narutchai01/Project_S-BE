package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type ThreadRepository interface {
	CreateThread(user_id uint) (entities.Thread, error)
	CreateThreadDetail(threadDetail entities.ThreadDetail) (entities.ThreadDetail, error)
	GetThreads() ([]entities.Thread, error)
	GetThread(id uint) (entities.Thread, error)
	GetThreadDetails(thread_id uint) ([]entities.ThreadDetail, error)
}
