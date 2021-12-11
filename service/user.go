package service

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/service/utils"

	"github.com/hackathon-21winter-05/HiQidas/model"
)

type UserService interface {
	GetUsersID(c context.Context) (model.UserIDs, error)
	GetUserByID(c context.Context, myUserID uuid.UUID) (*model.User, error)
	GetHeyaByUserMe(c context.Context, myUserID uuid.UUID) ([]*model.Heya, error)
	CreateUser(c context.Context, name string) (*model.User, error)

	/* 未実装
	GetUserMeFavorites(c context.Context)
	*/
}

func (s *Service) GetUsersID(c context.Context) (model.UserIDs, error) {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()
	userIDs, err := s.repo.GetUsersID(ctx)
	if err != nil {
		return nil, err
	}

	return userIDs, nil
}

func (s *Service) GetUserByID(c context.Context, myUserID uuid.UUID) (*model.User, error) {
	panic("implement me")
}

func (s *Service) GetHeyaByUserMe(c context.Context, myUserID uuid.UUID) ([]*model.Heya, error) {
	panic("implement me")
}

func (s *Service) CreateUser(c context.Context, name string) (*model.User, error) {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	id := utils.GetUUID()
	user := model.User{
		ID:   id,
		Name: name,
	}

	err := s.repo.Do(ctx, nil, func(ctx context.Context) error {

		if err := s.repo.CreateUser(ctx, &user); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}
