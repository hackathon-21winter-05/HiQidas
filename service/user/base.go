package user

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/service/utils"
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

	ctx, cancel := utils.CreateTxContext()
	defer cancel()

	err = us.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
