package model

import "github.com/gofrs/uuid"

type Drawing struct {
	ID      uuid.UUID `json:"id"    gorm:"type:char(36);not null;primaryKey"`
	Drawing string    `json:"drawing" gorm:"type:text;not null"`
}
