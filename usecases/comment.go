package usecases

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
)

type CommentUsecase interface {
	CreateComment(comment entities.Comment, token string) (entities.Comment, error)
}

type commentService struct {
	repo repositories.CommentRepository
}

func NewCommentUseCase(repo repositories.CommentRepository) CommentUsecase {
	return &commentService{repo}
}

func (service *commentService) CreateComment(comment entities.Comment, token string) (entities.Comment, error) {

	user_id, err := utils.ExtractToken(token)

	if err != nil {
		return entities.Comment{}, err
	}

	comment.UserID = user_id

	return service.repo.CreateComment(comment)
}
