package model

import "github.com/gofrs/uuid"

type History struct {
	UserID uuid.UUID `json:"user_id"  gorm:"type:char(36);not null;primaryKey;index:idx_history_user_id,priority:1"`
	HeyaID uuid.UUID `json:"heya_id" gorm:"type:char(36);not null;primaryKey;index:idx_history_user_id,priority:2"`
}
