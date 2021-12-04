package model

import (
	"github.com/gofrs/uuid"
	"time"
)

type Hiqidashi struct {
	ID             uuid.UUID      `json:"id"             gorm:"type:char(36);not null;primaryKey"`
	HeyaID         uuid.UUID      `json:"heya_id"        gorm:"type:char(36);not null;primaryKey"`
	CreatorID      uuid.UUID      `json:"creator_id"     gorm:"type:char(36);not null"`
	LastEditorID   uuid.NullUUID  `json:"last_editor_id" gorm:"type:char(36);not null"`
	ParentID       uuid.NullUUID  `json:"parent_id"      gorm:"type:char(36)"`
	Title          string         `json:"title"          gorm:"type:char(50);not null"`
	Description    string         `json:"description"    gorm:"type:text"`
	ImageID        uuid.UUID      `json:"image_id"       gorm:"type:char(36)"`
	CreatedAt      time.Time      `json:"created_at"     gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time      `json:"updated_at"     gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP"`
	CreatorUser    *User          `json:"-"              gorm:"foreignKey:CreatorID"`
	LastEditorUser *User          `json:"-"              gorm:"foreignKey:LastEditorID"`
	Image          *Image         `json:"-"              gorm:"foreignKey:ImageID"`
	Heya           *Heya          `json:"-"              gorm:"foreignKey:HeyaID"`

	ChildrenID []uuid.UUID `gorm:"-"`
}
