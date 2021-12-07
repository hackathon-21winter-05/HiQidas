package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var allTables = []interface{}{
	model.User{},
	model.Heya{},
	model.History{},
	model.Hiqidashi{},
	model.Drawing{},
	model.Tsuna{},
}

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(c *config.Config) (Repository, error) {
	db, err := newDBConnection(c)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db :%w", err)
	}

	return &GormRepository{db: db}, nil
}

func newDBConnection(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", c.MariaDBUsername, c.MariaDBPassword, c.MariaDBHostname, c.MariaDBDatabase) + "?parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4"
	logLevel := logger.Info

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logLevel)})
	if err != nil {
		return nil, fmt.Errorf("failed to connect DB : %w", err)
	}

	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci")

	err = db.AutoMigrate(allTables...)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate: %w", err)
	}

	return db, nil
}

// GetDB DBをコンテキストから取得
func (gr *GormRepository) getDB(ctx context.Context) (db *gorm.DB, err error) {
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
