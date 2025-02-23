package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type ThreadRepository interface {
	CreateThread(user_id uint, title string, image string) (entities.Thread, error)
	CreateThreadDetail(threadDetail entities.ThreadDetail) (entities.ThreadDetail, error)
	GetThreads() ([]entities.Thread, error)
	GetThread(id uint) (entities.Thread, error)
	GetThreadDetails(thread_id uint) ([]entities.ThreadDetail, error)
	DeleteThread(thread_id uint) error
	UpdateThread(thread entities.Thread) (entities.Thread, error)
	UpdateThreadDetail(threadDetails entities.ThreadDetail) (entities.ThreadDetail, error)
	GetThreadDetail(id uint) (entities.ThreadDetail, error)
}
