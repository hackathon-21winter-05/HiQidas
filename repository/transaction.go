package repository

import (
	"context"
	"database/sql"
)

type TransactionRepository interface {
	Do(context.Context, *sql.TxOptions, func(context.Context) error) error
}
