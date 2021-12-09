package user

import (
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/repository"
)

type UserService interface {
	GetUsersID() (model.UserIDs, error)
	CreateUser(name string) (*model.User, error)
}

type UserServiceImpl struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) UserService {
	return &UserServiceImpl{repo: repo}
}
