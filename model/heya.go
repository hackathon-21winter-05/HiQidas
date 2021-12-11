package model

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
)

type Heya struct {
	ID           uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	Title        string    `gorm:"type:char(50);not null"`
	Description  string    `gorm:"type:text;not null"`
	CreatorID    uuid.UUID `gorm:"type:char(36);not null"`
	LastEditorID uuid.UUID `gorm:"type:char(36);not null"`
	CreatedAt    time.Time `gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP"`

	Creator User           `gorm:"foreignKey:ID;references:CreatorID"`
}

type NullHeya struct {
	ID           uuid.UUID
	Title        sql.NullString
	Description  sql.NullString
	LastEditorID uuid.UUID
	UpdatedAt    sql.NullTime
}
