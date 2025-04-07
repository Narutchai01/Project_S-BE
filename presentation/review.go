package presentation

import (
	"github.com/Narutchai01/Project_S-BE/entities"
)

func PublicReviewSkincare(data entities.Community) ReviewSkincare {

	var skincares []entities.Skincare
	for _, skincare := range data.Skincares {
		skincares = append(skincares, skincare.Skincare)
	}

	// Initialize with default empty string
	var image string
	var imageID uint
	if len(data.Images) > 0 {
		image = data.Images[0].Image
		imageID = data.Images[0].ID
	}

	return ReviewSkincare{
		ID:            data.ID,
		Title:         data.Title,
		Content:       data.Caption,
		Favortie:      data.Favorite,
		FavoriteCount: int64(data.Likes),
		Bookmark:      data.Bookmark,
		Owner:         data.Owner,
		Image:         image,
		ImageID:       imageID,
		User:          *PublicUser(data.User),
		Skincare:      MapPubliceSkincare(skincares),
		CreateAt:      data.CreatedAt,
	}
}

func ToReviewResponse(data entities.Community) *Responses {
	if data.ID == 0 {
		return &Responses{
			Status: false,
			Data:   nil,
			Error:  []string{"Review not found"},
		}
	}

	return &Responses{
		Status: true,
		Data:   PublicReviewSkincare(data),
		Error:  nil,
	}
}

func ToReviewsResponse(data []entities.Community) *Responses {
	var responses []ReviewSkincare

	if len(data) == 0 {
		return &Responses{
			Status: true,
			Data:   []ReviewSkincare{},
			Error:  nil,
		}
	}

	for _, review := range data {
		if review.ID == 0 {
			continue
		}

		if len(review.Images) == 0 || len(review.Skincares) == 0 {
			continue
		}
		responses = append(responses, PublicReviewSkincare(review))
	}
	return &Responses{
		Status: true,
		Data:   responses,
		Error:  nil,
	}
}
