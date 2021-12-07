package repository

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

// GetUsersID GetUsers 全てのUserIDを取得
func (repo *GormRepository) GetUsersID(ctx context.Context) ([]uuid.UUID, error) {
	db, err := repo.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	var usersID []uuid.UUID
	err = db.
		Model(&model.User{}).
		Pluck("id", &usersID).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get users : %w", err)
	}

	return usersID, nil
}

// GetUserByID IDからUserを取得する
func (repo *GormRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	db, err := repo.getDB(ctx)
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
func (repo *GormRepository) CreateUser(ctx context.Context, user *model.User) error {
	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	err = db.Create(&user).Error
	if err != nil {
		return fmt.Errorf("failed to create user :%w", err)
	}

	return nil
}

// DeleteUserByID Userを削除する
func (repo *GormRepository) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	db, err := repo.getDB(ctx)
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
func (repo *GormRepository) UpdateUserByID(ctx context.Context, user *model.User) error {
	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	result := db.
		Model(model.User{}).
		Where("id = ?", user.ID).
		//更新するものがIconのみ
		Update("icon_file_id", &user.IconFileID)
	err = result.Error
	if err != nil {
		return fmt.Errorf("failed to update user : %w", err)
	}
	if result.RowsAffected == 0 {
		return ErrNoRecordUpdated
	}

	return nil
}
