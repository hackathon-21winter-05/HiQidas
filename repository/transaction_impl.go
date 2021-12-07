package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gorm.io/gorm"
)


const (
	txKey string = "transaction"
)

// Do Transaction用のメソッド
func (tx *GormRepository) Do(ctx context.Context, options *sql.TxOptions, f func(context.Context) error) error {
	fc := func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, txKey, tx)

		err := f(ctx)
		if err != nil {
			return err
		}

		return nil
	}

	if options == nil {
		err := tx.db.Transaction(fc)
		if err != nil {
			return fmt.Errorf("failed to get transaciton:%w", err)
		}
	} else {
		err := tx.db.Transaction(fc, options)
		if err != nil {
			return fmt.Errorf("failed to get transaciton:%w", err)
		}
	}

	return nil
}
