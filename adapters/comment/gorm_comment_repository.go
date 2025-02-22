package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
)

type GormCommentRepository struct {
	db *gorm.DB
}

func NewGormCommentRepository(db *gorm.DB) repositories.CommentRepository {
	return &GormCommentRepository{db: db}
}

func (repo *GormCommentRepository) CreateComment(comment entities.Comment) (entities.Comment, error) {
	if err := repo.db.Create(&comment).Error; err != nil {
		return entities.Comment{}, err
	}
	return comment, nil
}

func (repo *GormCommentRepository) GetComments(thread_id uint) ([]entities.Comment, error) {
	var comments []entities.Comment

	if err := repo.db.Preload("User").Where("thread_id = ?", thread_id).Find(&comments).Error; err != nil {
		return []entities.Comment{}, err
	}

	return comments, nil
}
