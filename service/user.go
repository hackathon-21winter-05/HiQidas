package service

import (
	"context"

	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/repository"
)

type UserService interface {
	GetUsersID() (model.UserIDs, error)
}

type UserServiceImpl struct {
	repo repository.Repository
}

func newUserService(repo repository.Repository) UserService {
	return &UserServiceImpl{repo: repo}
}

func (us *UserServiceImpl) GetUsersID() (model.UserIDs, error) {
	userIDs, err := us.repo.GetUsersID(context.Background())
	if err != nil {
		return nil, err
	}

	return userIDs, nil
}
