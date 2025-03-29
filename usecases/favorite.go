package usecases

import (
	"errors"
	"strings"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
)

type FavoriteUseCase interface {
	Favorite(favorite entities.Favorite, type_community string, token string) (entities.Favorite, error)
}

type favoriteService struct {
	repo     repositories.FavoriteRepository
	userRepo repositories.UserRepository

	commentRepo   repositories.CommentRepository
	communityRepo repositories.CommunityRepository
}

func NewFavoriteUseCase(repo repositories.FavoriteRepository, userRepo repositories.UserRepository, commemntRepo repositories.CommentRepository, communityrepo repositories.CommunityRepository) FavoriteUseCase {
	return &favoriteService{repo, userRepo, commemntRepo, communityrepo}
}

func (service *favoriteService) Favorite(favorite entities.Favorite, type_community string, token string) (entities.Favorite, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.Favorite{}, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.Favorite{}, errors.New("user not found")
	}

	favorite.UserID = user.ID

	if favorite.CommunityID != 0 && favorite.CommentID == 0 {
		community_type, err := service.communityRepo.GetCommunityType(strings.ToLower(type_community))
		if err != nil {
			return entities.Favorite{}, errors.New("community not found")
		}

		community, err := service.communityRepo.GetCommunity(favorite.CommunityID, uint64(community_type.ID))
		if err != nil {
			return entities.Favorite{}, errors.New("community not found")
		}

		result, fav_id, err := service.repo.FindFavorite(community.ID, "community_id", uint(user.ID))
		if err != nil {
			return entities.Favorite{}, err
		}

		if !result {
			return service.repo.Favorite(favorite)
		} else {
			return service.repo.DeleteFavorite(fav_id)
		}
	} else if favorite.CommunityID == 0 && favorite.CommentID != 0 {
		comment, err := service.commentRepo.GetComment(favorite.CommentID)
		if err != nil {
			return entities.Favorite{}, errors.New("comment not found")
		}

		result, fav_id, err := service.repo.FindFavorite(comment.ID, "comment_id", user.ID)
		if err != nil {
			return entities.Favorite{}, err
		}

		if !result {
			return service.repo.Favorite(favorite)
		} else {
			return service.repo.DeleteFavorite(fav_id)
		}
	}

	return favorite, nil
}
