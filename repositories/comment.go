package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type CommentRepository interface {
	CreateComment(comment entities.Comment) (entities.Comment, error)
	GetComment(id uint) (entities.Comment, error)
	GetComments(community_id uint) ([]entities.Comment, error)
}
