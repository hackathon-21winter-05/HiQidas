package database

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

type User struct {
	db *DB
}

func NewUser(db *DB) *User {
	return &User{db: db}
}

func (u *User) GetUsers(ctx context.Context) ([]*model.User, error) {
	panic("implement me")
}

func (u *User) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	panic("implement me")
}

func (u *User) CreateUser(ctx context.Context, user *model.User) error {
	panic("implement me")
}

func (u *User) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}

func (u *User) UpdateUserByID(ctx context.Context, user *model.User) error {
	panic("implement me")
}



