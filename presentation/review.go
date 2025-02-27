package presentation

import (
	"github.com/Narutchai01/Project_S-BE/entities"
)

func PublicReviewSkincare(data entities.ReviewSkincare) ReviewSkincare {
	return ReviewSkincare{
		ID:            data.ID,
		Title:         data.Title,
		Content:       data.Content,
		Favortie:      data.Favorite,
		FavoriteCount: data.FavoriteCount,
		Bookmark:      data.Bookmark,
		Owner:         data.Owner,
		Image:         data.Image,
		User:          *PublicUser(data.User),
		Skincare:      MapPubliceSkincare(data.Skincare),
		CreateAt:      data.CreatedAt,
	}
}

func ToReviewResponse(data entities.ReviewSkincare) *Responses {
	return &Responses{
		Status: true,
		Data:   PublicReviewSkincare(data),
		Error:  nil,
	}
}

func ToReviewsResponse(data []entities.ReviewSkincare) *Responses {
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
