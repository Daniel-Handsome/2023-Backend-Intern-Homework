package user

import (
	"context"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/model"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/repository"
)

type UserService interface {
	GetUserArticlesHeadKey(ctx context.Context, uuid string) (string, error)
}

type userService struct {
	Repo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{Repo: userRepo}
}

func (srv userService) GetUserArticlesHeadKey(ctx context.Context, uuid string) (string, error) {
	pageLinkedList, err := srv.Repo.GetPageLinkedList(ctx, uuid, model.RECOMMEND)
	if err != nil {
		return "", err
	}

	return pageLinkedList.Head, nil
}
