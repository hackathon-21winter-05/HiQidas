package model

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
)

type Hiqidashi struct {
	ID           uuid.UUID      `json:"id"             gorm:"type:char(36);not null;primaryKey"`
	HeyaID       uuid.UUID      `json:"heya_id"        gorm:"type:char(36);not null;index:idx_hiqidashi_heya_id,priority:1"`
	CreatorID    uuid.UUID      `json:"creator_id"     gorm:"type:char(36);not null;index:idx_hiqidashi_creator_id,priority:1"`
	LastEditorID uuid.UUID      `json:"last_editor_id" gorm:"type:char(36);not null"`
	ParentID     uuid.NullUUID  `json:"parent_id"      gorm:"type:char(36)"`
	Title        string         `json:"title"          gorm:"type:char(50);not null"`
	Description  sql.NullString `json:"description"    gorm:"type:text"`
	ImageID      uuid.UUID      `json:"image_id"       gorm:"type:char(36)"`
	ColorID      int            `json:"color_id"       gorm:"type:TINYINT UNSIGNED;not null"`
	CreatedAt    time.Time      `json:"created_at"     gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP;index:idx_hiqidashi_heya_id,priority:2,index:idx_hiqidashi_creator_id,priority:2"`
	UpdatedAt    time.Time      `json:"updated_at"     gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP"`

	ChildrenID []uuid.UUID `gorm:"-"`
}
