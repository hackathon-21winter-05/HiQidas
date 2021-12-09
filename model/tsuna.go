package model

import "github.com/gofrs/uuid"

type Tsuna struct {
	ID           uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	HeyaID       uuid.UUID `gorm:"type:char(36);not null"`
	HiqidashiOne uuid.UUID `gorm:"type:char(36);not null"`
	HiqidashiTwo uuid.UUID `gorm:"type:char(36);not null"`
}

// NullTsuna Oneを変えずにtwoをUpdateする
type NullTsuna struct {
	ID           uuid.UUID
	HiqidashiTwo uuid.UUID
}