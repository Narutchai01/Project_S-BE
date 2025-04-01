package presentation

import "github.com/Narutchai01/Project_S-BE/entities"

func PublicSkinResponse(data entities.FaceProblem) Skin {
	skin := Skin{
		ID:       data.ID,
		Name:     data.Name,
		Image:    data.Image,
		CreateBY: uint(data.CreatedBy),
	}
	return skin
}

func ToSkinResponse(data entities.FaceProblem) *Responses {
	skin := PublicSkinResponse(data)
	return &Responses{
		Status: true,
		Data:   skin,
		Error:  nil,
	}
}

func ToSkinsResponse(data []entities.FaceProblem) *Responses {
	skins := []Skin{}

	for _, skin := range data {
		skins = append(skins, PublicSkinResponse(skin))
	}
	return &Responses{
		Status: true,
		Data:   skins,
		Error:  nil,
	}
}
