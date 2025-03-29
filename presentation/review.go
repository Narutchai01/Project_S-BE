package presentation

import (
	"github.com/Narutchai01/Project_S-BE/entities"
)

func PublicReviewSkincare(data entities.Community) ReviewSkincare {

	var skincares []entities.Skincare
	for _, skincare := range data.Skincares {
		skincares = append(skincares, skincare.Skincare)
	}

	return ReviewSkincare{
		ID:            data.ID,
		Title:         data.Title,
		Content:       data.Caption,
		Favortie:      data.Favorite,
		FavoriteCount: int64(data.Likes),
		Bookmark:      data.Bookmark,
		Owner:         data.Owner,
		Image:         data.Images[0].Image,
		User:          *PublicUser(data.User),
		Skincare:      MapPubliceSkincare(skincares),
		CreateAt:      data.CreatedAt,
	}
}

func ToReviewResponse(data entities.Community) *Responses {
	return &Responses{
		Status: true,
		Data:   PublicReviewSkincare(data),
		Error:  nil,
	}
}

func ToReviewsResponse(data []entities.Community) *Responses {
	var responses []ReviewSkincare

	for _, review := range data {
		responses = append(responses, PublicReviewSkincare(review))
	}
	return &Responses{
		Status: true,
		Data:   responses,
		Error:  nil,
	}

}
