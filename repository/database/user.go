package database

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

type User struct {
	db *DB
}

func NewUser(db *DB) *User {
	return &User{db: db}
}

// GetUsers 全てのUserを取得
func (u *User) GetUsers(ctx context.Context) ([]*model.User, error) {
	db, err := u.db.GetDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}
	users := make([]*model.User, 0)

	err = db.
		Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get users : %w", err)
	}

	return users, nil
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
