package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
)

type GormReviewRepository struct {
	db *gorm.DB
}

func NewGormReviewRepository(db *gorm.DB) repositories.ReviewRepository {
	return &GormReviewRepository{db: db}
}

func (repo *GormReviewRepository) CreateReviewSkincare(review entities.ReviewSkincare) (entities.ReviewSkincare, error) {
	err := repo.db.Create(&review).Error
	if err != nil {
		return entities.ReviewSkincare{}, err
	}
	return review, nil
}
