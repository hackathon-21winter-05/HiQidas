package model

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
)

type Heya struct {
	ID           uuid.UUID      `json:"id"             gorm:"type:char(36);not null;primaryKey"`
	Title        string         `json:"title"          gorm:"type:char(50);not null"`
	Description  sql.NullString `json:"description"    gorm:"type:text"`
	CreatorID    uuid.UUID      `json:"creator_id"     gorm:"type:char(36);not null"`
	LastEditorID uuid.UUID      `json:"last_editor_id" gorm:"type:char(36);not null"`
	CreatedAt    time.Time      `json:"created_at"     gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `json:"updated_at"     gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP"`
	Deleted      bool           `json:"deleted"        gorm:"type:boolean;default:false;not null"`
}
