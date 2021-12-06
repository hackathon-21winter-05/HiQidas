package model

import "github.com/gofrs/uuid"

type User struct {
	ID         uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	Name       string    `gorm:"type:varchar(32);not null;unique"`
	IconFileID uuid.UUID `gorm:"type:char(36);not null"`
}
