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

func (repo *GormThreadRepository) CreateThread(user_id uint, title string, image string) (entities.Thread, error) {

	thread := entities.Thread{
		UserID: user_id,
		Title:  title,
		Image:  image,
	}

	if err := repo.db.Create(&thread).Error; err != nil {
		return entities.Thread{}, err
	}

	return thread, nil

}

func (repo *GormThreadRepository) CreateThreadDetail(threadDetail entities.ThreadDetail) (entities.ThreadDetail, error) {
	if err := repo.db.Create(&threadDetail).Error; err != nil {
		return entities.ThreadDetail{}, err
	}

	return threadDetail, nil
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

func (repo *GormThreadRepository) DeleteThread(thread_id uint) error {
	if err := repo.db.Where("id = ?", thread_id).Delete(&entities.Thread{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *GormThreadRepository) UpdateThread(thread entities.Thread) (entities.Thread, error) {
	if err := repo.db.Save(&thread).Error; err != nil {
		return entities.Thread{}, err
	}
	return thread, nil
}

func (repo *GormThreadRepository) UpdateThreadDetail(threadDetail entities.ThreadDetail) (entities.ThreadDetail, error) {
	if err := repo.db.Save(&threadDetail).Error; err != nil {
		return entities.ThreadDetail{}, err
	}
	return threadDetail, nil
}

func (repo *GormThreadRepository) GetThreadDetail(id uint) (entities.ThreadDetail, error) {
	var threadDetail entities.ThreadDetail

	if err := repo.db.First(&threadDetail, id).Error; err != nil {
		return entities.ThreadDetail{}, err
	}

	return threadDetail, nil
}
