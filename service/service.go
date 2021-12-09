package service

import (
	"github.com/hackathon-21winter-05/HiQidas/service/heya"
)

type Service struct {
	UserService
	heya.HeyaService
}

func NewService(userService UserService, heyaService heya.HeyaService) *Service {
	return &Service{UserService: userService, HeyaService: heyaService}
}