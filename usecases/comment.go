package usecases

import (
	"strings"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
)

type CommentUsecase interface {
	CreateComment(comment entities.Comment, token string, type_community string) (entities.Comment, error)
	GetComments(community_id uint, type_community string, token string) ([]entities.Comment, error)
}

type commentService struct {
	repo          repositories.CommentRepository
	favoriteRepo  repositories.FavoriteRepository
	userRepo      repositories.UserRepository
	communityRepo repositories.CommunityRepository
}

func NewCommentUseCase(repo repositories.CommentRepository, favoriteRepo repositories.FavoriteRepository, userRepo repositories.UserRepository, communityRepo repositories.CommunityRepository) CommentUsecase {
	return &commentService{repo, favoriteRepo, userRepo, communityRepo}
}

func (service *commentService) CreateComment(comment entities.Comment, token string, type_community string) (entities.Comment, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.Comment{}, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.Comment{}, err
	}

	community_type, err := service.communityRepo.GetCommunityType(strings.ToLower(type_community))
	if err != nil {
		return entities.Comment{}, err
	}

	community, err := service.communityRepo.GetCommunity(comment.CommunityID, uint64(community_type.ID))
	if err != nil {
		return entities.Comment{}, err
	}

	comment.CommunityID = community.ID
	comment.UserID = user.ID

	comment, err = service.repo.CreateComment(comment)
	if err != nil {
		return entities.Comment{}, err
	}

	comment, err = service.repo.GetComment(comment.ID)
	if err != nil {
		return entities.Comment{}, err
	}

	return comment, nil
}

func (service *commentService) GetComments(community_id uint, type_community string, token string) ([]entities.Comment, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return []entities.Comment{}, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return []entities.Comment{}, err
	}

	community_type, err := service.communityRepo.GetCommunityType(strings.ToLower(type_community))
	if err != nil {
		return []entities.Comment{}, err
	}

	community, err := service.communityRepo.GetCommunity(community_id, uint64(community_type.ID))
	if err != nil {
		return []entities.Comment{}, err
	}

	comments, err := service.repo.GetComments(community.ID)
	if err != nil {
		return []entities.Comment{}, err
	}

	for i, comment := range comments {
		isFavorite, _, err := service.favoriteRepo.FindFavorite(comment.ID, "comment_id", user.ID)
		if err != nil {
			return []entities.Comment{}, err
		}
		comments[i].Favorite = isFavorite
		comments[i].FavoriteCount = int(service.favoriteRepo.CountFavorite(comment.ID, "comment_id"))
	}

	return comments, nil
}
