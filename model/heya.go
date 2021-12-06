package model

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
)

type Heya struct {
	ID           uuid.UUID      `gorm:"type:char(36);not null;primaryKey"`
	Title        string         `gorm:"type:char(50);not null"`
	Description  sql.NullString `gorm:"type:text"`
	CreatorID    uuid.UUID      `gorm:"type:char(36);not null"`
	LastEditorID uuid.UUID      `gorm:"type:char(36);not null"`
	CreatedAt    time.Time      `gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP"`
	Deleted      bool           `gorm:"type:boolean;default:false;not null"`
}
