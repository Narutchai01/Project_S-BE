package presentation

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
)


type Skincare struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	CreateBY    uint `json:"create_by"`
}

func SkincareResponse(data entities.Skincare) *fiber.Map {
	skincare := Skincare{
		ID:       data.ID,
		Name: data.Name,
		Image:    data.Image,
		Description: data.Description,
		CreateBY: data.CreateBY,
	}

	return &fiber.Map{
		"status": true,
		"skincare":  skincare,
		"error":  nil,
	}
}

func SkincaresResponse(data []entities.Skincare) *fiber.Map {
	skincares := []Skincare{}

	for _, skincare := range data {
		skincares = append(skincares, Skincare{
			ID:       skincare.ID,
			Name: skincare.Name,
			Image:    skincare.Image,
			CreateBY: skincare.CreateBY,
		})
	}
	return &fiber.Map{
		"status": true,
		"data":   skincares,
		"error":  nil,
	}
}

func SkincareErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"skincare":  nil,
		"error":  err.Error(),
	}
}

func DeleteSkincareResponse(id int) *fiber.Map {
	return &fiber.Map{
		"status":    true,
		"delete_id": id,
		"error":     nil,
	}
}
