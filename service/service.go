package service

import (
	"github.com/hackathon-21winter-05/HiQidas/repository"
	"github.com/hackathon-21winter-05/HiQidas/service/client"
)

type Service struct {
	UserService
	client.HeyaClientService
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		UserService:   newUserService(repo),
		HeyaClientService: client.NewClientService(),
	}
}
