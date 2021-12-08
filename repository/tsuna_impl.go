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

// CreateTsuna ツナを作成
func (repo *GormRepository) CreateTsuna(ctx context.Context, tsuna *model.Tsuna) error {
	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	err = db.Create(&tsuna).Error
	if err != nil {
		return fmt.Errorf("failed to create tsuna :%w", err)
	}

	return nil
}

// DeleteTsuna ツナを削除する
func (repo *GormRepository) DeleteTsuna(ctx context.Context, id uuid.UUID) error {
	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	result := db.
		Where("id = ?", id).
		Delete(&model.Tsuna{})
	err = result.Error
	if err != nil {
		return fmt.Errorf("failed to delete tsuna :%w", err)
	}
	if result.RowsAffected == 0 {
		return ErrNoRecordDeleted
	}

	return nil
}

func (repo *GormRepository) UpdateTsuna(Ctx context.Context, tsuna *model.Tsuna) error {
	panic("implement me")
}
