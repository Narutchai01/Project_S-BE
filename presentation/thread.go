package presentation

import "github.com/Narutchai01/Project_S-BE/entities"

func PublicThreadImage(threadImages []entities.ThreadImage) []ThreadImage {
	var images []ThreadImage
	for _, image := range threadImages {
		images = append(images, ThreadImage{
			ID:       image.ID,
			ThreadID: image.ThreadID,
			Image:    image.Image,
		})
	}

	return images
}

func PublicThread(thread entities.Thread) Thread {
	return Thread{
		ID:            thread.ID,
		Title:         thread.Title,
		Favorite:      thread.Favorite,
		FavoriteCount: thread.FavoriteCount,
		Bookmark:      thread.Bookmark,
		User:          *PublicUser(thread.User),
		Caption:       thread.Caption,
		Images:        PublicThreadImage(thread.Images),
		CreateAt:      thread.CreatedAt,
	}
}

func ToThreadResponse(data entities.Thread) *Responses {
	return &Responses{
		Status: true,
		Data:   PublicThread(data),
		Error:  nil,
	}
}

func ToThreadsResponse(data []entities.Thread) *Responses {
	var threads []Thread
	for _, thread := range data {
		threads = append(threads, PublicThread(thread))
	}

	return &Responses{
		Status: true,
		Data:   threads,
		Error:  nil,
	}
}
