package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type ReviewRepository interface {
	CreateReviewSkincare(reviewThread entities.ReviewSkincare) (entities.ReviewSkincare, error)
	GetReviewSkincare(id uint) (entities.ReviewSkincare, error)
	GetReviewSkincares() ([]entities.ReviewSkincare, error)
}
