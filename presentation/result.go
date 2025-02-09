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
		Data:   result,
		Error:  nil,
	}
}
