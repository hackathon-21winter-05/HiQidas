package service

import (
	"context"

	"github.com/hackathon-21winter-05/HiQidas/model"
)

type UserService interface {
	GetUsersID() (model.UserIDs, error)
}

func (s *Service) GetUsersID() (model.UserIDs, error) {
	userIDs, err := s.repo.GetUsersID(context.Background())
	if err != nil {
		return nil, err
	}

	return userIDs, nil
}
