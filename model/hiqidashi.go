package model

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
)

type Hiqidashi struct {
	ID           uuid.UUID      `gorm:"type:char(36);not null;primaryKey"`
	HeyaID       uuid.UUID      `gorm:"type:char(36);not null;index:idx_hiqidashi_heya_id,priority:1"`
	CreatorID    uuid.UUID      `gorm:"type:char(36);not null;index:idx_hiqidashi_creator_id,priority:1"`
	LastEditorID uuid.UUID      `gorm:"type:char(36);not null"`
	ParentID     uuid.NullUUID  `gorm:"type:char(36)"`
	Title        string         `gorm:"type:char(50);not null"`
	Description  string         `gorm:"type:text;not null"`
	Drawing      sql.NullString `gorm:"type:text"`
	ColorCode    string         `gorm:"type:char(7);default:#9E7A7A;not null"`
	CreatedAt    time.Time      `gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP;index:idx_hiqidashi_heya_id,priority:2,index:idx_hiqidashi_creator_id,priority:2"`
	UpdatedAt    time.Time      `gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP"`

	ChildrenID []uuid.UUID `gorm:"-"`
}

type NullHiqidashi struct {
	ID           uuid.UUID
	LastEditorID uuid.UUID
	ParentID     uuid.NullUUID
	Title        sql.NullString
	Description  sql.NullString
	Drawing      sql.NullString
	ColorID      sql.NullString
	UpdatedAt    time.Time
}
