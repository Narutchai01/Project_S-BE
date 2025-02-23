package presentation

import "github.com/Narutchai01/Project_S-BE/entities"

func SkincareMap(data entities.Skincare) Skincare {
	skincare := Skincare{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Image:       data.Image,
		CreateBY:    data.CreateBY,
	}

	return skincare
}

func MapThreadDetails(data []entities.ThreadDetail) []ThreadDetail {
	threadDetail := []ThreadDetail{}

	for _, detail := range data {
		threadDetail = append(threadDetail, ThreadDetail{
			ID:       detail.ID,
			Skincare: SkincareMap(detail.Skincare),
			Caption:  detail.Caption,
		})
	}

	return threadDetail
}

func PublicThread(data entities.Thread) *Thread {
	thread := Thread{
		ID:           data.ID,
		UserID:       data.UserID,
		Title:        data.Title,
		Bookmark:     data.Bookmark,
		Image:        data.Image,
		Favorite:     data.Favorite,
		Owner:        data.Owner,
		User:         *PublicUser(data.User),
		ThreadDetail: MapThreadDetails(data.Threads),
	}

	return &thread
}

func ToThreadResponse(data entities.Thread) *Responses {

	thread := PublicThread(data)

	return &Responses{
		Status: true,
		Data:   thread,
		Error:  nil,
	}
}

func ToThreadListResponse(data []entities.Thread) *Responses {
	threads := []Thread{}
	for _, thread := range data {
		threads = append(threads, *PublicThread(thread))
	}
	return &Responses{
		Status: true,
		Data:   threads,
		Error:  nil,
	}
}
