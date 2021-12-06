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

// GetUserByID IDからUserを取得する
func (u *User) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	db, err := u.db.GetDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	var user *model.User

	err = db.
		Where("id = ?", id).
		First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id : %w", err)
	}

	return user, nil
}

// CreateUser Userを作成
func (u *User) CreateUser(ctx context.Context, user *model.User) error {
	db, err := u.db.GetDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	err = db.Create(&user).Error
	if err != nil {
		return fmt.Errorf("failed to create user :%w", err)
	}

	return nil
}

func (u *User) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	db, err := u.db.GetDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	result := db.
		Where("id = ?", id).
		Delete(&model.User{})
	err = result.Error
	if err != nil {
		return fmt.Errorf("failed to delete user :%w", err)
	}
	if result.RowsAffected == 0 {
		return ErrNoRecordDeleted
	}

	return nil
}

// UpdateUserByID ユーザーの情報を更新
func (u *User) UpdateUserByID(ctx context.Context, user *model.User) error {
	db, err := u.db.GetDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	result := db.
		Model(model.User{}).
		Where("id = ?", user.ID).
		// どのフィールドもゼロ値になることがないため構造体で渡す
		Updates(&user)
	err = result.Error
	if err != nil {
		return fmt.Errorf("failed to update user : %w", err)
	}
	if result.RowsAffected == 0 {
		return ErrNoRecordUpdated
	}

	return nil
}
