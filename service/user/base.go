package user

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

func (us *UserServiceImpl) GetUsersID() (model.UserIDs, error) {
	userIDs, err := us.repo.GetUsersID(context.Background())
	if err != nil {
		return nil, err
	}

	return userIDs, nil
}

func (us *UserServiceImpl) CreateUser(name string) (*model.User, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:         id,
		Name:       name,
		IconFileID: uuid.Nil,
	}

	err = us.repo.Do(context.Background(), nil, func(ctx context.Context) error {
		return us.repo.CreateUser(ctx, user)
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
