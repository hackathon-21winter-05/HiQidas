package service

import "github.com/hackathon-21winter-05/HiQidas/repository"

type Service struct {
	UserService
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		UserService: newUserService(repo),
	}
}
