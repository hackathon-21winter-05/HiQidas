package service

import (
	"github.com/hackathon-21winter-05/HiQidas/repository"
	"github.com/hackathon-21winter-05/HiQidas/service/user"
)

type Service struct {
	user.UserService
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		UserService: user.NewUserService(repo),
	}
}
