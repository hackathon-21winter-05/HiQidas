package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type Heyas struct {
	ID           uuid.UUID      `json:"id" gorm:"type:char(36);not null;primaryKey"`
	Title        string         `json:"title" gorm:"type:char(50);not null"`
	Description  string         `json:"description" gorm:"type:text;not null"`
	CreatorID    uuid.UUID      `json:"creator_id" gorm:"type:char(36);not null"`
	LastEditorID uuid.UUID      `json:"last_editor_id" gorm:"type:char(36);not null"`
	CreatedAt    time.Time      `json:"created_at"     gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `json:"updated_at"     gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"     gorm:"type:DATETIME;default:NULL"`
	CreatorUser        User          `json:"-" gorm:"foreignKey:ID"`
}
