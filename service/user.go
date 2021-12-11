package service

import (
	"context"
	"github.com/hackathon-21winter-05/HiQidas/service/utils"

	"github.com/hackathon-21winter-05/HiQidas/model"
)

type UserService interface {
	GetUsersID(ctx context.Context) (model.UserIDs, error)
}

func (s *Service) GetUsersID(ctx context.Context) (model.UserIDs, error) {
	ctx ,cancel := utils.CreateTxContext(ctx)
	defer cancel()
	userIDs, err := s.repo.GetUsersID(ctx)
	if err != nil {
		return nil, err
	}

	return userIDs, nil
}
