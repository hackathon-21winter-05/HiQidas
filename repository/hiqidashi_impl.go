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
	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	err = db.Create(&hiqidashi).Error
	if err != nil {
		return fmt.Errorf("failed to create hiqidashi :%w", err)
	}

	return nil
}

func (repo *GormRepository) DeleteHiqidashi(ctx context.Context, id uuid.UUID) error {
	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	result := db.
		Where("id = ?", id).
		Delete(&model.Hiqidashi{})
	err = result.Error
	if err != nil {
		return fmt.Errorf("failed to deleted hiqidashi :%w", err)
	}
	if result.RowsAffected == 0 {
		return ErrNoRecordDeleted
	}

	return nil
}

func (repo *GormRepository) UpdateHiqidashi(ctx context.Context, hiqidashi *model.Hiqidashi) error {
	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}
	hiqidashiMap := map[string]interface{}{
		"id":             hiqidashi.ID,
		"heya_id":        hiqidashi.HeyaID,
		"creator_id":     hiqidashi.CreatorID,
		"last_editor_id": hiqidashi.LastEditorID,
		"parent_id":      hiqidashi.ParentID,
		"title":          hiqidashi.Title,
		"description":    hiqidashi.Description,
		"drawing":        hiqidashi.Drawing,
		"colorID":        hiqidashi.ColorID,
		"created_at":     hiqidashi.CreatedAt,
		"updated_at":     hiqidashi.UpdatedAt,
	}
	result := db.Updates(&hiqidashiMap)
	err = result.Error
	if err != nil {
		return fmt.Errorf("failed to update hiqidashi :%w", err)
	}
	if result.RowsAffected == 0 {
		return ErrNoRecordUpdated
	}

	return nil
}
