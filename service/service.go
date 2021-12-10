package service

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/repository"
)

type Service struct {
	repo repository.Repository
}



func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}
