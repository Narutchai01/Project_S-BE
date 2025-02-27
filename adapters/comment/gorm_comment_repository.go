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

func (repo *GormCommentRepository) CreateCommentThread(comment entities.CommentThread) (entities.CommentThread, error) {
	if err := repo.db.Create(&comment).Error; err != nil {
		return entities.CommentThread{}, err
	}
	return comment, nil
}

func (repo *GormCommentRepository) GetCommentsThread(thread_id uint) ([]entities.CommentThread, error) {
	var comments []entities.CommentThread

	if err := repo.db.Preload("User").Where("thread_id = ?", thread_id).Find(&comments).Error; err != nil {
		return []entities.CommentThread{}, err
	}

	return comments, nil
}

func (repo *GormCommentRepository) CreateCommentReviewSkicnare(comment entities.CommentReviewSkicare) (entities.CommentReviewSkicare, error) {
	if err := repo.db.Create(&comment).Error; err != nil {
		return entities.CommentReviewSkicare{}, err
	}
	return comment, nil
}

func (repo *GormCommentRepository) GetCommentsReviewSkincare(review_id uint) ([]entities.CommentReviewSkicare, error) {
	var comments []entities.CommentReviewSkicare

	if err := repo.db.Preload("User").Where("review_skincare_id = ?", review_id).Find(&comments).Error; err != nil {
		return []entities.CommentReviewSkicare{}, err
	}

	return comments, nil
}
