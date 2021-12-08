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
	ur repository.UserRepository
}

func (us *UserServiceImpl) GetUsersID() (model.UserIDs, error) {
	userIDs, err := us.ur.GetUsersID(context.Background())
	if err != nil {
		return nil, err
	}

	return userIDs, nil
}
