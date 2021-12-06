package repository

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	DeleteUserByID(ctx context.Context, id uuid.UUID) error
	UpdateUserByID(ctx context.Context, user *model.User) error
}
