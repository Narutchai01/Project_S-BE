package presentation

import (
	"github.com/Narutchai01/Project_S-BE/entities"
)

func ToResultResponse(data entities.Result) *Responses {
	result := Result{
		ID:         data.ID,
		Image:      data.Image,
		UserId:     data.UserId,
		AcneType:   data.AcneType,
		FacialType: data.FacialType,
		SkinType:   data.SkinType,
		Skincare:   data.Skincare,
	}

	return &Responses{
		Status: true,
		Data:   result,
		Error:  nil,
	}
}

func ResultsResponse(data []entities.Result) *Responses {
	results := []Result{}

	for _, result := range data {
		results = append(results, Result{
			ID:         result.ID,
			Image:      result.Image,
			UserId:     result.UserId,
			AcneType:   result.AcneType,
			FacialType: result.FacialType,
			SkinType:   result.SkinType,
			Skincare:   result.Skincare,
		})
	}

	return &Responses{
		Status: true,
		Data:   results,
		Error:  nil,
	}
}
