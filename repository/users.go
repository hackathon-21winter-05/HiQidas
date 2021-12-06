package repository

import (
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

type UserRepository interface {
	GetUsers() ([]*model.User, error)
	GetUserByID(id uuid.UUID) (*model.User, error)
	CreateUser(user *model.User) error
	DeleteUserByID(id uuid.UUID) error
	UpdateUserByID(user *model.User) error
}
