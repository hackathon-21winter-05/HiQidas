package repository

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

func (repo *GormRepository) GetFavoritesByUserID(ctx context.Context, userID uuid.UUID) ([]*model.Favorite, error) {
	db, err := repo.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	favorites := make([]*model.Favorite, len(userID))

	err = db.
		Where("user_id = ?", userID).
		Find(&favorites).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get favorites : %w", err)
	}

	return favorites, nil
}

func (repo *GormRepository) CreateFavorite(ctx context.Context, favorite *model.Favorite) error {
	if favorite.UserID == uuid.Nil || favorite.HeyaID == uuid.Nil {
		return ErrNillUUID
	}

	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	err = db.Create(&favorite).Error
	if err != nil {
		return fmt.Errorf("failed to create favorite  :%w", err)
	}

	return nil
}

func (repo *GormRepository) DeleteFavoriteByHeyaID(ctx context.Context, heyaID uuid.UUID) error {
	if heyaID == uuid.Nil {
		return ErrNillUUID
	}

	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	// ヘヤが削除されたときにそのヘヤを含むお気に入りを一括削除
	result := db.
		Where("heya_id = ?", heyaID).
		Delete(&model.Favorite{})
	err = result.Error

	if err != nil {
		return fmt.Errorf("failed to delete favorite :%w", err)
	}
	if result.RowsAffected == 0 {
		return ErrNoRecordDeleted
	}

	return nil
}

