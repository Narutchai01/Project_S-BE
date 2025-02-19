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
