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

func (repo *GormThreadRepository) CreateThread(thread entities.Thread) (entities.Thread, error) {
	err := repo.db.Create(&thread).Error
	return thread, err
}

func (repo *GormThreadRepository) CreateThreadImage(thread entities.ThreadImage) (entities.ThreadImage, error) {
	err := repo.db.Create(&thread).Error
	return thread, err
}

func (repo *GormThreadRepository) GetThread(thread_id uint) (entities.Thread, error) {
	thread := entities.Thread{}
	err := repo.db.Preload("User").Where("id = ?", thread_id).First(&thread).Error
	return thread, err
}

func (repo *GormThreadRepository) GetThreadImages(thread_id uint) ([]entities.ThreadImage, error) {
	threadImages := []entities.ThreadImage{}
	err := repo.db.Where("thread_id = ?", thread_id).Find(&threadImages).Error
	return threadImages, err
}
