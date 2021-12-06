package database

import (
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

type DB struct {
	db *gorm.DB
}

func NewDBConnect(isProduction bool) (*DB, error) {
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
	if isProduction {
		logLevel = logger.Silent
	} else {
		logLevel = logger.Info
	}

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
