package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var allTables = []interface{}{
	model.User{},
	model.Heya{},
	model.History{},
	model.Hiqidashi{},
	model.Drawing{},
	model.Tsuna{},
}

type ctxKey string

const (
	txKey ctxKey = "transaction"
)

type DB struct {
	db *gorm.DB
}

func NewDBConnect() (*DB, error) {
	user, ok := os.LookupEnv("DB_USERNAME")
	if !ok {
		return nil, errors.New("DB_USERNAME is not set")
	}

	pass, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return nil, errors.New("DB_PASSWORD is not set")
	}

	host, ok := os.LookupEnv("DB_HOSTNAME")
	if !ok {
		return nil, errors.New("DB_HOSTNAME is not set")
	}

	dbname, ok := os.LookupEnv("DB_DATABASE")
	if !ok {
		return nil, errors.New("DB_DATABASE is not set")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, pass, host, dbname) + "?parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4"

	var logLevel logger.LogLevel
	logLevel = logger.Info

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logLevel)})
	if err != nil {
		return nil, fmt.Errorf("failed to connect DB : %w", err)
	}

	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci")

	err = db.AutoMigrate(allTables...)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate: %w", err)
	}

	return &DB{
		db: db,
	}, nil
}

// GetDB DBをコンテキストから取得
func (d *DB) GetDB(ctx context.Context) (db *gorm.DB, err error) {
	iDB := ctx.Value(txKey)
	if iDB == nil {
		return d.db.WithContext(ctx), nil
	}

	gormDB, ok := iDB.(*gorm.DB)
	if !ok {
		return nil, errors.New("failed to get gorm.DB")
	}

	return gormDB.WithContext(ctx), nil
}

// Do Transaction用のメソッド
func (d *DB) Do(ctx context.Context, options *sql.TxOptions, f func(context.Context) error) error {
	fc := func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, txKey, tx)

		err := f(ctx)
		if err != nil {
			return err
		}

		return nil
	}

	if options == nil {
		err := d.db.Transaction(fc)
		if err != nil {
			return fmt.Errorf("failed to get transaciton:%w", err)
		}
	} else {
		err := d.db.Transaction(fc, options)
		if err != nil {
			return fmt.Errorf("failed to get transaciton:%w", err)
		}
	}

	return nil
}
