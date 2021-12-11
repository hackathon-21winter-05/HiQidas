package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"gorm.io/gorm"
)

func (repo *GormRepository) GetHiqidashiByID(ctx context.Context, id uuid.UUID) (*model.Hiqidashi, error) {
	if id == uuid.Nil {
		return nil, ErrNillUUID
	}

	db, err := repo.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	hiqidashi := &model.Hiqidashi{}
	err = db.
		Where("id = ?", id).
		First(&hiqidashi).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get	hiqidashi :%w", err)
	}

	return hiqidashi, nil
}

// GetHiqidashisByHeyaID ヘヤのすべてのヒキダシを取得
func (repo *GormRepository) GetHiqidashisByHeyaID(ctx context.Context, heyaID uuid.UUID) ([]*model.Hiqidashi, error) {
	if heyaID == uuid.Nil {
		return nil, ErrNillUUID
	}
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
	if hiqidashi.ID == uuid.Nil {
		return ErrNillUUID
	}

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

// DeleteHiqidashiByID  ヒキダシを削除
func (repo *GormRepository) DeleteHiqidashiByID(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return ErrNillUUID
	}

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

// UpdateHiqidashiByID UpdateHiqidashi ヒキダシを更新
func (repo *GormRepository) UpdateHiqidashiByID(ctx context.Context, hiqidashi *model.NullHiqidashi) error {
	if hiqidashi.ID == uuid.Nil || hiqidashi.LastEditorID == uuid.Nil {
		return ErrNillUUID
	}

	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	hiqidashiMap := map[string]interface{}{}
	hiqidashiMap["id"] = hiqidashi.ID
	hiqidashiMap["last_editor_id"] = hiqidashi.LastEditorID
	if hiqidashi.ParentID.Valid {
		hiqidashiMap["parent_id"] = hiqidashi.ParentID
	}
	if hiqidashi.Drawing.Valid {
		hiqidashiMap["drawing"] = hiqidashi.Drawing
	}
	if hiqidashi.Title.Valid {
		hiqidashiMap["title"] = hiqidashi.Title
	}
	if hiqidashi.Description.Valid {
		hiqidashiMap["description"] = hiqidashi.Description
	}
	hiqidashiMap["updated_at"] = hiqidashi.UpdatedAt

	result := db.
		Model(&model.Hiqidashi{}).Where("id = ?", hiqidashi.ID).
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

func (repo *GormRepository) DeleteHiqidashiDrawing(ctx context.Context, hiqidashi *model.Hiqidashi) error {
	if hiqidashi.ID == uuid.Nil {
		return ErrNillUUID
	}
	if hiqidashi.Drawing.Valid {
		return nil
	}

	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	err = db.
		Where("id = ?", hiqidashi.ID).
		Update("drawing", gorm.Expr("NULL")).Error

	if err != nil {
		return fmt.Errorf("failed to drawing nil :%w", err)
	}

	return nil
}

func (repo *GormRepository) DeleteHiqidashisByHeyaID(ctx context.Context, heyaID uuid.UUID) error {
	if heyaID == uuid.Nil {
		return ErrNillUUID
	}

	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	result := db.
		Where("heya_id = ?", heyaID).
		Delete(&model.Hiqidashi{})
	err = result.Error
	if err != nil {
		return fmt.Errorf("failed to delete hiqidashi by heyaID :%w", err)
	}
	if result.RowsAffected == 0 {
		return ErrNoRecordDeleted
	}

	return nil
}

func (repo *GormRepository) UpdateRootHiqidashiByHeyaID(ctx context.Context, heyaID uuid.UUID, title, description sql.NullString) error {
	if heyaID == uuid.Nil {
		return ErrNillUUID
	}
	hiqidashiMap := map[string]interface{}{}
	hiqidashiMap["heya_id"] = heyaID
	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}
	if title.Valid {
		hiqidashiMap["title"] = title.String
	}
	if description.Valid {
		hiqidashiMap["description"] = description.String
	}

	result := db.
		Model(&model.Hiqidashi{}).
		Where("heya_id = ? AND parent_id is NULL", heyaID).
		Updates(hiqidashiMap)
	err = result.Error
	if err != nil {
		return fmt.Errorf("failed to updates root hiqidashi :%w", err)
	}
	if result.RowsAffected == 0 {
		return ErrNoRecordUpdated
	}

	return nil
}
