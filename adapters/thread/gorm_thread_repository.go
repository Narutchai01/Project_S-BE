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

func (repo *GormThreadRepository) CreateThread(user_id uint) (entities.Thread, error) {

	thread := entities.Thread{
		UserID: user_id,
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

func (repo *GormThreadRepository) CreateBookmark(thread_id uint, user_id uint) (entities.Bookmark, error) {
	bookmark := entities.Bookmark{
		ThreadID: thread_id,
		UserID:   user_id,
	}
	if err := repo.db.Create(&bookmark).Error; err != nil {
		return entities.Bookmark{}, err
	}
	return bookmark, nil
}

func (repo *GormThreadRepository) FindBookMark(thread_id uint, user_id uint) (entities.Bookmark, error) {
	var bookmark entities.Bookmark
	if err := repo.db.Where("thread_id = ? AND user_id = ?", thread_id, user_id).First(&bookmark).Error; err != nil {
		return entities.Bookmark{}, err
	}
	return bookmark, nil
}

func (repo *GormThreadRepository) UpdateBookMark(thread_id uint, user_id uint, status bool) (entities.Bookmark, error) {
	var bookmark entities.Bookmark
	if err := repo.db.Where("thread_id = ? AND user_id = ?", thread_id, user_id).First(&bookmark).Error; err != nil {
		return entities.Bookmark{}, err
	}

	bookmark.Status = &status
	if err := repo.db.Save(&bookmark).Error; err != nil {
		return entities.Bookmark{}, err
	}

	return bookmark, nil
}
