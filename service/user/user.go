package user

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/repository"

	"github.com/hackathon-21winter-05/HiQidas/model"
)

type UserService interface {
	GetUsersID(c context.Context) (model.UserIDs, error)
	GetUserByID(c context.Context, myUserID uuid.UUID) (*model.User, error)
	GetHeyaByUserMe(c context.Context, myUserID uuid.UUID) ([]*model.Heya, error)
	CreateUser(c context.Context, id, iconFileID uuid.UUID, name string) (*model.User, error)
}

type UserServiceImpl struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) UserService {
	return &UserServiceImpl{
		repo: repo,
	}
}
