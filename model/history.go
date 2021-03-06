package model

import "github.com/gofrs/uuid"

type History struct {
	UserID uuid.UUID `gorm:"type:char(36);not null;primaryKey;index:idx_history_user_id,priority:1"`
	HeyaID uuid.UUID `gorm:"type:char(36);not null;primaryKey;index:idx_history_user_id,priority:2"`

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Heya Heya `gorm:"foreignKey:HeyaID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
