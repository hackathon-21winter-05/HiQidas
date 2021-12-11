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
	CreateTraPUser(c context.Context, id, iconFileID uuid.UUID, name string) (*model.User, error)
	GetUserMeFavorites(c context.Context, userID uuid.UUID) ([]uuid.UUID, error)
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
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	user, err := s.repo.GetUserByID(ctx, myUserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetHeyaByUserMe(c context.Context, myUserID uuid.UUID) ([]*model.Heya, error) {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	heyas, err := s.repo.GetHeyasByCreatorID(ctx, myUserID)
	if err != nil {
		return nil, err
	}

	return heyas, nil
}

func (s *Service) GetUserMeFavorites(c context.Context, userID uuid.UUID) ([]*model.Favorite, error) {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	favorites,err := s.repo.GetFavoritesByUserID(ctx,userID)
	if err != nil {
		return nil, err
	}

	return favorites,nil
}

func (s *Service) CreateUser(c context.Context, name string) (*model.User, error) {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	id := utils.GetUUID()
	user := model.User{
		ID:         id,
		Name:       name,
		IconFileID: utils.NullUUID(),
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

func (s *Service) CreateTraPUser(c context.Context, id, iconFileID uuid.UUID, name string) (*model.User, error) {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	user := model.User{
		ID:         id,
		Name:       name,
		IconFileID: utils.NullUUIDFrom(iconFileID),
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
