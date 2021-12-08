package repository

import (
	"context"
	"database/sql"
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
	panic("implement me")
}

func (repo *GormRepository) CreateHeya(ctx context.Context, title string, description sql.NullString) (*model.Heya, error) {
	panic("implement me")
}

func (repo *GormRepository) UpdateHeya(ctx context.Context, heya *model.Heya) error {
	panic("implement me")
}

func (repo *GormRepository) DeleteHeya(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}
