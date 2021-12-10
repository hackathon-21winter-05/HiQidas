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
	if id == uuid.Nil {
		return nil, model.ErrNillUUID
	}

	db, err := repo.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	var heya *model.Heya

	err = db.
		Where("id = ?", id).
		First(&heya).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get heya :%w", err)
	}

	return heya, nil
}

func (repo *GormRepository) CreateHeya(ctx context.Context, heya *model.Heya) error {
	if heya.ID == uuid.Nil {
		return model.ErrNillUUID
	}

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

func (repo *GormRepository) UpdateHeyaByID(ctx context.Context, heya *model.NullHeya) error {
	if heya.ID == uuid.Nil || heya.LastEditorID == uuid.Nil {
		return model.ErrNillUUID
	}

	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	heyaMap := map[string]interface{}{}

	heyaMap["id"] = heya.ID
	heyaMap["last_editor_id"] = heya.LastEditorID
	if heya.Title.Valid {
		heyaMap["title"] = heya.Title.String
	}
	if heya.Description.Valid {
		heyaMap["description"] = heya.Description.String
	}
	if heya.UpdatedAt.Valid {
		heyaMap["updated_at"] = heya.UpdatedAt
	}

	result := db.
		Model(&model.Heya{}).
		Updates(&heyaMap)
	err = result.Error
	if err != nil {
		return fmt.Errorf("failed to update heya:%w", err)
	}
	if result.RowsAffected == 0 {
		return model.ErrNoRecordUpdated
	}

	return nil
}

func (repo *GormRepository) DeleteHeyaByID(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return model.ErrNillUUID
	}

	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	result := db.
		Where("id = ?", id).
		Delete(&model.Heya{})
	err = result.Error
	if err != nil {
		return fmt.Errorf("failed to deleted heya :%w", err)
	}
	if result.RowsAffected == 0 {
		return model.ErrNoRecordDeleted
	}

	return nil
}
