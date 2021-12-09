package repository

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

// GetHistoriesByUserID ユーザーの履歴の取得
func (repo *GormRepository) GetHistoriesByUserID(ctx context.Context, userID uuid.UUID) ([]*model.History, error) {
	db, err := repo.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	histories := make([]*model.History, len(userID))

	err = db.
		Where("id = ?", userID).
		Find(&histories).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get histories : %w", err)
	}

	return histories, nil
}

// CreateHistory 履歴の作成
func (repo *GormRepository) CreateHistory(ctx context.Context, history *model.History) error {
	if history.UserID == uuid.Nil || history.HeyaID == uuid.Nil {
		return model.ErrNillUUID
	}

	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	err = db.Create(&history).Error
	if err != nil {
		return fmt.Errorf("failed to create history :%w", err)
	}

	return nil
}

// DeleteHistoryByHeyaID DeleteHistory ヘヤが削除されたときに削除する履歴
func (repo *GormRepository) DeleteHistoryByHeyaID(ctx context.Context, heyaID uuid.UUID) error {
	if heyaID == uuid.Nil {
		return model.ErrNillUUID
	}

	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	// ヘヤが削除されたときにそのヘヤを含む履歴を一括削除
	result := db.
		Where("heya_id = ?", heyaID).
		Delete(&model.History{})
	err = result.Error

	if err != nil {
		return fmt.Errorf("failed to delete history :%w", err)
	}
	if result.RowsAffected == 0 {
		return model.ErrNoRecordDeleted
	}

	return nil
}
