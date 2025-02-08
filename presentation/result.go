package presentation

import (
	"github.com/Narutchai01/Project_S-BE/entities"
)

func MapSkinCare(data []entities.Skincare) []Skincare {
	skincare := []Skincare{}
	for _, skin := range data {
		skincare = append(skincare, Skincare{
			ID:          skin.ID,
			Name:        skin.Name,
			Description: skin.Description,
			Image:       skin.Image,
		})
	}
	return skincare
}

func ToResultResponse(data entities.Result) *Responses {
	result := Result{
		ID:         data.ID,
		UserID:     data.UserID,
		Image:      data.Image,
		AcneTpye:   data.AcneType,
		FacialType: data.FacialType,
		Skincare:   MapSkinCare(data.Skincare),
	}
	return &Responses{
		Status: true,
		Data:   []Result{result},
		Error:  nil,
	}
}

func ToResultsResponse(data []entities.Result) *Responses {
	results := make([]Result, len(data))
	for i, result := range data {
		results[i] = Result{
			ID:         result.ID,
			UserID:     result.UserID,
			Image:      result.Image,
			AcneTpye:   result.AcneType,
			FacialType: result.FacialType,
			Skincare:   MapSkinCare(result.Skincare),
		}
	}
	return &Responses{
		Status: true,
		Data:   results,
		Error:  nil,
	}
}
