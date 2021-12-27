package model

import "github.com/gofrs/uuid"

type Favorite struct {
	UserID uuid.UUID `gorm:"type:char(36);not null;primaryKey;index:idx_favorite_user_id,priority:1"`
	User   User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	HeyaID uuid.UUID `gorm:"type:char(36);not null;primaryKey;index:idx_favorite_user_id,priority:2"`
	Heya   Heya      `gorm:"foreignKey:HeyaID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
