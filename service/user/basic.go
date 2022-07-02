package user

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/service/utils"
)

func (s *UserServiceImpl) GetUsersID(c context.Context) (model.UserIDs, error) {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()
	userIDs, err := s.repo.GetUsersID(ctx)
	if err != nil {
		return nil, err
	}

	return userIDs, nil
}

func (s *UserServiceImpl) GetUserByID(c context.Context, myUserID uuid.UUID) (*model.User, error) {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	user, err := s.repo.GetUserByID(ctx, myUserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServiceImpl) GetHeyaByUserMe(c context.Context, myUserID uuid.UUID) ([]*model.Heya, error) {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	heyas, err := s.repo.GetHeyasByCreatorID(ctx, myUserID)
	if err != nil {
		return nil, err
	}

	return heyas, nil
}

func (s *UserServiceImpl) GetUserMeFavorites(c context.Context, userID uuid.UUID) ([]*model.Favorite, error) {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	favorites, err := s.repo.GetFavoritesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return favorites, nil
}

func (s *UserServiceImpl) CreateUser(c context.Context, id, iconFileID uuid.UUID, name string) (*model.User, error) {
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
