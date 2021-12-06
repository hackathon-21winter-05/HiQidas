package repository

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type DB interface {
	GetDB(ctx context.Context) (db *gorm.DB, err error)
	Do(context.Context, *sql.TxOptions, func(context.Context) error) error
}
