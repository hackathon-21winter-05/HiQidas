package repository

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

// GetTsunasByHeyaID ヘヤにあるすべてのツナを取得する
func (repo *GormRepository) GetTsunasByHeyaID(ctx context.Context, heyaID uuid.UUID) ([]*model.Tsuna, error) {
	db, err := repo.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	tsunas := make([]*model.Tsuna, 0)

	err = db.
		Where("heya_id = ?", heyaID).
		Find(&tsunas).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get tsunas :%w ", err)
	}

	return tsunas, nil
}

func (repo *GormRepository) CreateTsuna(ctx context.Context, tsuna *model.Tsuna) error {
	panic("implement me")
}

func (repo *GormRepository) DeleteTsuna(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}

func (repo *GormRepository) UpdateTsuna(Ctx context.Context, tsuna *model.Tsuna) error {
	panic("implement me")
}
