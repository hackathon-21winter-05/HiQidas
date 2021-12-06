package model

import "github.com/gofrs/uuid"

type Drawing struct {
	ID      uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	Drawing string    `gorm:"type:text;not null"`
}
