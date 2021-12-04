package model

import "github.com/gofrs/uuid"

type Image struct {
	ID    uuid.UUID `json:"id"    gorm:"type:char(36);not null;primaryKey"`
	Image []byte    `json:"image" gorm:"type:MEDIUMBLOB;not null"`
}
