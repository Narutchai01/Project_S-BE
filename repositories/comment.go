package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type CommentRepository interface {
	CreateComment(comment entities.Comment) (entities.Comment, error)
	GetComments(thread_id uint) ([]entities.Comment, error)
}
