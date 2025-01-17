package presentation

import "github.com/Narutchai01/Project_S-BE/entities"

type Skin struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	CreateBY uint   `json:"create_by"`
}

func ToSkinResponse(data entities.Skin) *Responses {
	skin := Skin{
		ID:       data.ID,
		Name:     data.Name,
		Image:    data.Image,
		CreateBY: data.CreateBY,
	}
	return &Responses{
		Status: true,
		Data:   skin,
		Error:  nil,
	}
}

func ToSkinsResponse(data []entities.Skin) *Responses {
	skins := []Skin{}

	for _, skin := range data {
		skins = append(skins, Skin{
			ID:       skin.ID,
			Name:     skin.Name,
			Image:    skin.Image,
			CreateBY: skin.CreateBY,
		})
	}
	return &Responses{
		Status: true,
		Data:   skins,
		Error:  nil,
	}
}
