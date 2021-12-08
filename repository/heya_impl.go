package repository

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

// GetHeyasID すべてのヘヤのIDを取得
func (repo *GormRepository) GetHeyasID(ctx context.Context) ([]uuid.UUID, error) {
	db, err := repo.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	var heyaIDs []uuid.UUID

	err = db.
		Model(model.Heya{}).
		Pluck("id", &heyaIDs).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get heyaIDs :%w", err)
	}

	return heyaIDs, nil
}

func (repo *GormRepository) GetHeyaByID(ctx context.Context, id uuid.UUID) (*model.Heya, error) {
	db, err := repo.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	var heya *model.Heya

	err = db.
		Where("db = ?", id).
		First(&heya).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get heya :%w", err)
	}

	return heya, nil
}

func (repo *GormRepository) CreateHeya(ctx context.Context, heya *model.Heya) error {
	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	err = db.Create(&heya).Error
	if err != nil {
		return fmt.Errorf("failed to create : %w", err)
	}

	return nil
}

func (repo *GormRepository) UpdateHeya(ctx context.Context, heya *model.Heya) error {

}

func (repo *GormRepository) DeleteHeya(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}
