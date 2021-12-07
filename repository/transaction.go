package repository

import (
	"context"
	"database/sql"
)

type Transaction interface {
	Do(context.Context, *sql.TxOptions, func(context.Context) error) error
}
