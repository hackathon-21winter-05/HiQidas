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

func (repo *GormRepository) CreateHistory(ctx context.Context, history *model.History) error {
	panic("implement me")
}

func (repo *GormRepository) DeleteHistory(ctx context.Context, history *model.History) error {
	panic("implement me")
}
