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

func MapSkin(data entities.Skin) Skin {
	skin := Skin{
		ID:       data.ID,
		Name:     data.Name,
		Image:    data.Image,
		CreateBY: data.CreateBY,
	}
	return skin
}

func ToResultResponse(data entities.Result) *Responses {
	result := Result{
		ID:         data.ID,
		UserID:     data.UserID,
		Image:      data.Image,
		SkincareID: data.SkincareID,
		AcneTpye:   data.AcneType,
		FacialType: data.FacialType,
		SkinID:     data.SkinID,
		Skin:       MapSkin(data.Skin),
		Skincare:   MapSkinCare(data.Skincare),
		CreateAt:   &data.CreatedAt,
	}
	return &Responses{
		Status: true,
		Data:   result,
		Error:  nil,
	}
}

func ToResultsResponse(data []entities.Result) *Responses {
	results := []Result{}
	for _, result := range data {
		results = append(results, Result{
			ID:         result.ID,
			UserID:     result.ID,
			Image:      result.Image,
			SkincareID: result.SkincareID,
			AcneTpye:   result.AcneType,
			FacialType: result.FacialType,
			SkinID:     result.SkinID,
			Skin:       MapSkin(result.Skin),
			Skincare:   MapSkinCare(result.Skincare),
			CreateAt:   &result.CreatedAt,
		})
	}

	return &Responses{
		Status: true,
		Data:   results,
		Error:  nil,
	}
}
