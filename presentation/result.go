package presentation

import (
	"github.com/Narutchai01/Project_S-BE/entities"
)

func PublicResultResponse(data entities.Result) Result {

	skincares := []Skincare{}
	for _, skincare := range data.Skincare {
		skincares = append(skincares, PubliceSkincare(skincare.Skincare))
	}

	result := Result{
		ID:         data.ID,
		UserID:     data.UserID,
		Image:      data.Image,
		SkincareID: data.SkincareID,
		AcneTpye:   data.AcneType,
		FacialType: data.FacialType,
		SkinID:     data.SkinID,
		Skin:       PublicSkinResponse(data.Skin),
		Skincare:   skincares,
		CreateAt:   &data.CreatedAt,
	}
	return result
}

func ToResultResponse(data entities.Result) *Responses {
	result := PublicResultResponse(data)
	return &Responses{
		Status: true,
		Data:   result,
		Error:  nil,
	}
}

func ToResultsResponse(data []entities.Result) *Responses {
	results := []Result{}
	for _, result := range data {
		results = append(results, PublicResultResponse(result))
	}

	return &Responses{
		Status: true,
		Data:   results,
		Error:  nil,
	}
}
