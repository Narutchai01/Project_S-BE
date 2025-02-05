package presentation

import (
	"github.com/Narutchai01/Project_S-BE/entities"
)

func ToResultResponse(data entities.Result) *Responses {
	skincares := make([]Skincare, len(data.Skincare))
	for i, s := range data.Skincare {
		skincares[i] = Skincare{
			ID: s.ID,
			Image:       s.Image,
			Name:        s.Name,
			Description: s.Description,
		}
	}

	result := Result{
		ID:         data.ID,
		Image:      data.Image,
		UserId:     data.UserId,
		AcneType:   data.AcneType,
		FacialType: data.FacialType,
		SkinType:   data.SkinType,
		Skincare:   skincares,
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
		skincares := make([]Skincare, len(result.Skincare))
		for i, s := range result.Skincare {
			skincares[i] = Skincare{
				ID: s.ID,
				Image:       s.Image,
				Name:        s.Name,
				Description: s.Description,
			}
		}

		results = append(results, Result{
			ID:         result.ID,
			Image:      result.Image,
			UserId:     result.UserId,
			AcneType:   result.AcneType,
			FacialType: result.FacialType,
			SkinType:   result.SkinType,
			Skincare:   skincares,
		})
	}

	return &Responses{
		Status: true,
		Data:   results,
		Error:  nil,
	}
}