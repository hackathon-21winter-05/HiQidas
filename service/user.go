package service

import (
	"github.com/hackathon-21winter-05/HiQidas/repository"
)

type UserService interface {
}

type UserServiceImpl struct {
	ur repository.UserRepository
}
