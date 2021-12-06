package model

import "github.com/gofrs/uuid"

type Tsuna struct {
	ID           uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	HiqidashiOne uuid.UUID `gorm:"type:char(36);not null"`
	HiqidashiTwo uuid.UUID `gorm:"type:char(36);not null"`
}
