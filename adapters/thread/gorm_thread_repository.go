package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
)

type GormThreadRepository struct {
	db *gorm.DB
}

func NewGormThreadRepository(db *gorm.DB) repositories.ThreadRepository {
	return &GormThreadRepository{db: db}
}

func (repo *GormThreadRepository) CreateThread(threadDetails []entities.ThreadDetail, user_id uint) (entities.Thread, error) {
	var user entities.User
	if err := repo.db.First(&user, user_id).Error; err != nil {
		return entities.Thread{}, err
	}

	thread := entities.Thread{
		UserID:  user.ID,
		User:    user,
		Threads: make([]entities.ThreadDetail, 0),
	}

	if err := repo.db.Create(&thread).Error; err != nil {
		return entities.Thread{}, err
	}

	for _, detail := range threadDetails {
		var skincare entities.Skincare
		if err := repo.db.First(&skincare, detail.SkincareID).Error; err != nil {
			return entities.Thread{}, err
		}
		threadDetail := entities.ThreadDetail{
			ThreadID:   thread.ID,
			SkincareID: detail.SkincareID,
			Skincare:   skincare,
			Caption:    detail.Caption,
		}

		if err := repo.db.Create(&threadDetail).Error; err != nil {
			return entities.Thread{}, err
		}

		thread.Threads = append(thread.Threads, threadDetail)
	}

	return thread, nil
}
func (repo *GormThreadRepository) GetThreads() ([]entities.Thread, error) {
	var threads []entities.Thread
	if err := repo.db.Preload("User").Preload("Threads").Find(&threads).Error; err != nil {
		return []entities.Thread{}, err
	}
	return threads, nil
}

func (repo *GormThreadRepository) GetThread(id uint) (entities.Thread, error) {
	var thread entities.Thread

	if err := repo.db.Preload("User").Preload("Threads").First(&thread, id).Error; err != nil {
		return entities.Thread{}, err
	}

	return thread, nil
}

func (repo *GormThreadRepository) GetThreadDetails(thread_id uint) ([]entities.ThreadDetail, error) {
	var threadDetail []entities.ThreadDetail

	if err := repo.db.Preload("Skincare").Where("thread_id = ?", thread_id).Find(&threadDetail).Error; err != nil {
		return []entities.ThreadDetail{}, err
	}

	return threadDetail, nil
}
