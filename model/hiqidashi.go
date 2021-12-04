package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type Hiqidashis struct {
	ID           uuid.UUID      `json:"id"             gorm:"type:char(36);not null;primaryKey"`
	SheetID      uuid.UUID      `json:"sheet_id"       gorm:"type:char(36);not null;primaryKey"`
	CreatorID    uuid.UUID      `json:"creator_id"     gorm:"type:char(36);not null"`
	LastEditorID uuid.NullUUID  `json:"last_editor_id" gorm:"type:char(36)"`
	ParentID     uuid.NullUUID  `json:"parent_id"      gorm:"type:char(36)"`
	Title        string         `json:"title"          gorm:"type:char(50);not null"`
	Description  string         `json:"description"    gorm:"type:text;not null"`
	ImageID      uuid.UUID      `json:"image_id"       gorm:"type:char(36);not null"`
	CreatedAt    time.Time      `json:"created_at"     gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `json:"updated_at"     gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"     gorm:"type:DATETIME;default:NULL"`
}
