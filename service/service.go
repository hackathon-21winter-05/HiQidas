package service

import (
	"github.com/hackathon-21winter-05/HiQidas/repository"
	"github.com/hackathon-21winter-05/HiQidas/service/client"
	"github.com/hackathon-21winter-05/HiQidas/service/hiqidashi"
)

type Service struct {
	client.ClientService
	hiqidashi.HiqidashiService
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		ClientService:    client.NewClientService(),
		HiqidashiService: hiqidashi.NewHiqidashiService(repo),
		repo:             repo,
	}
}
