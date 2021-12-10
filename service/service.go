package service

import (
	"github.com/hackathon-21winter-05/HiQidas/repository"
	"github.com/hackathon-21winter-05/HiQidas/service/client"
	"github.com/hackathon-21winter-05/HiQidas/service/hiqidashi"
)

type Service struct {
	UserService
	client.HeyaClientService
	hiqidashi.HiqidashiService
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		UserService:       newUserService(repo),
		HeyaClientService: client.NewClientService(),
		HiqidashiService:  hiqidashi.NewHiqidashiService(repo),
	}
}
