package model

import "github.com/gofrs/uuid"

type Tsuna struct {
	ID             uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	HeyaID         uuid.UUID `gorm:"type:char(36);not null"`
	HiqidashiOneID uuid.UUID `gorm:"type:char(36);not null"`
	HiqidashiTwoID uuid.UUID `gorm:"type:char(36);not null"`

	Heya         Heya      `gorm:"foreignKey:HeyaID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	HiqidashiOne Hiqidashi `gorm:"foreignKey:HiqidashiOneID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	HiqidashiTwo Hiqidashi `gorm:"foreignKey:HiqidashiTwoID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// NullTsuna Oneを変えずにtwoをUpdateする
type NullTsuna struct {
	ID           uuid.UUID
	HiqidashiTwo uuid.UUID
}
