package presentation

import (
	"github.com/Narutchai01/Project_S-BE/entities"
)

func PubliceSkincare(data entities.Skincare) Skincare {
	skincare := Skincare{
		ID:          data.ID,
		Name:        data.Name,
		Image:       data.Image,
		Description: data.Description,
		CreateBY:    data.CreateBY,
	}

	return skincare
}

func MapPubliceSkincare(data []entities.Skincare) []Skincare {
	var skincares []Skincare

	for _, skincare := range data {
		skincares = append(skincares, PubliceSkincare(skincare))
	}

	return skincares
}

func SkincareResponse(data entities.Skincare) *Responses {
	skincare := Skincare{
		ID:          data.ID,
		Name:        data.Name,
		Image:       data.Image,
		Description: data.Description,
		CreateBY:    data.CreateBY,
	}

	return &Responses{
		Status: true,
		Data:   skincare,
		Error:  nil,
	}
}

func SkincaresResponse(data []entities.Skincare) *Responses {
	skincares := []Skincare{}

	for _, skincare := range data {
		skincares = append(skincares, Skincare{
			ID:       skincare.ID,
			Name:     skincare.Name,
			Image:    skincare.Image,
			CreateBY: skincare.CreateBY,
		})
	}

	return &Responses{
		Status: true,
		Data:   skincares,
		Error:  nil,
	}
}
