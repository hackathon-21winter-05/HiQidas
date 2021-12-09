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

// GetHiqidashisByParentID　親ヒキダシに対する子供のヒキダシを取得
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

// CreateHiqidashi  ヒキダシを作成
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

// DeleteHiqidashi ヒキダシを削除
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

// UpdateHiqidashi ヒキダシを更新
func (repo *GormRepository) UpdateHiqidashi(ctx context.Context, hiqidashi *model.Hiqidashi) error {
	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	hiqidashiMap := map[string]interface{}{}

	if hiqidashi.ParentID.Valid {
		hiqidashiMap["parent_id"] = hiqidashi.ParentID
	}
	if hiqidashi.Drawing.Valid {
		hiqidashiMap["drawing"] = hiqidashi.Drawing
	}

	hiqidashiMap = map[string]interface{}{
		"id":             hiqidashi.ID,
		"heya_id":        hiqidashi.HeyaID,
		"creator_id":     hiqidashi.CreatorID,
		"last_editor_id": hiqidashi.LastEditorID,
		"title":          hiqidashi.Title,
		"description":    hiqidashi.Description,
		"colorID":        hiqidashi.ColorID,
		"created_at":     hiqidashi.CreatedAt,
		"updated_at":     hiqidashi.UpdatedAt,
	}

	result := db.
		Model(&hiqidashi).
		Updates(&hiqidashiMap)
	err = result.Error
	if err != nil {
		return fmt.Errorf("failed to update hiqidashi :%w", err)
	}
	if result.RowsAffected == 0 {
		return ErrNoRecordUpdated
	}

	return nil
}
