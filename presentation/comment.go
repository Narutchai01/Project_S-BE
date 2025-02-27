package presentation

import "github.com/Narutchai01/Project_S-BE/entities"

func PublicCommentThread(data entities.CommentThread) CommentThread {
	var commment = CommentThread{
		ID:            data.ID,
		ThreadID:      data.ThreadID,
		User:          *PublicUser(data.User),
		Favorite:      data.Favorite,
		FavoriteCount: data.FavoriteCount,
		Text:          data.Text,
	}
	return commment
}

func PublicCommentsThread(datas []entities.CommentThread) []CommentThread {
	var comments []CommentThread

	for _, comment := range datas {
		comments = append(comments, PublicCommentThread(comment))
	}

	return comments
}

func ToCommentThread(data entities.CommentThread) *Responses {

	comment := PublicCommentThread(data)
	return &Responses{
		Status: true,
		Data:   comment,
		Error:  nil,
	}
}

func ToCommentsThread(datas []entities.CommentThread) *Responses {
	comments := PublicCommentsThread(datas)

	return &Responses{
		Status: true,
		Data:   comments,
		Error:  nil,
	}
}
