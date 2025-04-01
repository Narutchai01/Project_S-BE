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

func (repo *GormCommentRepository) GetComment(id uint) (entities.Comment, error) {
	var comment entities.Comment
	if err := repo.db.Preload("User").Where("id = ?", id).First(&comment).Error; err != nil {
		return entities.Comment{}, err
	}
	return comment, nil
}

func (repo *GormCommentRepository) GetComments(community_id uint) ([]entities.Comment, error) {
	var comments []entities.Comment

	if err := repo.db.Preload("User").Where("community_id = ?", community_id).Find(&comments).Error; err != nil {
		return []entities.Comment{}, err
	}

	return comments, nil
}
