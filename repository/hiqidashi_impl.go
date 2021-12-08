package repository

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

// GetHiqidashisByHeyaID ヘヤのすべてのヒキダシを取得
func (repo *GormRepository) GetHiqidashisByHeyaID(ctx context.Context, heyaID uuid.UUID) ([]*model.Hiqidashi, error) {
	db, err := repo.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	hiqidashis := make([]*model.Hiqidashi, 0)

	err = db.
		Where("heya_id = ?", heyaID).
		Find(&hiqidashis).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get hiqidashi by heyaID :%w", err)
	}

	return hiqidashis, nil
}

func (repo *GormRepository) GetHiqidashisByParentID(ctx context.Context, parentID uuid.UUID) ([]*model.Hiqidashi, error) {
	db, err := repo.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	hiqidashis := make([]*model.Hiqidashi, 0)

	err = db.
		Where("parent_id = ?", parentID).
		Find(&hiqidashis).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get hiqidashi by parentID :%w", err)
	}

	return hiqidashis, nil
}

func (repo *GormRepository) CreateHiqidashi(ctx context.Context, hiqidashi *model.Hiqidashi) error {
	panic("implement me")
}

func (repo *GormRepository) DeleteHiqidashi(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}

func (repo *GormRepository) UpdateHiqidashi(ctx context.Context, hiqidashi *model.Hiqidashi) error {
	panic("implement me")
}
