package model

import "github.com/gofrs/uuid"

type Favorite struct {
	UserID uuid.UUID `gorm:"type:char(36);not null;primaryKey;index:idx_favorite_user_id,priority:1"`
	HeyaID uuid.UUID `gorm:"type:char(36);not null;primaryKey;index:idx_favorite_user_id,priority:2"`
}
