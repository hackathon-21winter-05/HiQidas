package repository

import (
	"context"
	"database/sql"
)

type DB interface {
	GetDB() (db *sql.DB, err error)
	Do(context.Context, *sql.TxOptions, func(context.Context) error) error
}
