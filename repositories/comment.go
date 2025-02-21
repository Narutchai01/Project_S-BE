package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type CommentRepository interface {
	CreateComment(comment entities.Comment) (entities.Comment, error)
}
