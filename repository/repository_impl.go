package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

// GetDB DBをコンテキストから取得
func (gr *GormRepository) getDB(ctx context.Context) (db *gorm.DB,err error) {
	iDB := ctx.Value(txKey)
	if iDB == nil {
		return gr.db.WithContext(ctx), nil
	}

	gormDB, ok := iDB.(*gorm.DB)
	if !ok {
		return nil, errors.New("failed to get gorm.DB")
	}

	return gormDB.WithContext(ctx), nil
}
