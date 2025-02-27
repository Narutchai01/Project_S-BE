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

func (repo *GormReviewRepository) GetReviewSkincare(id uint) (entities.ReviewSkincare, error) {
	var review entities.ReviewSkincare
	err := repo.db.Preload("User").Where("id = ? ", id).First(&review).Error

	return review, err
}

func (repo *GormReviewRepository) GetReviewSkincares() ([]entities.ReviewSkincare, error) {
	var reviews []entities.ReviewSkincare
	err := repo.db.Preload("User").Find(&reviews).Error

	return reviews, err
}
