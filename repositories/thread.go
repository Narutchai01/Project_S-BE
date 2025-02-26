package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type ThreadRepository interface {
	CreateThread(thread entities.Thread) (entities.Thread, error)
	CreateThreadImage(thread entities.ThreadImage) (entities.ThreadImage, error)
	GetThread(thread_id uint) (entities.Thread, error)
	GetThreadImages(thread_id uint) ([]entities.ThreadImage, error)
	GetThreads() ([]entities.Thread, error)
}
