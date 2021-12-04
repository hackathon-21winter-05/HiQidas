package model

import "github.com/gofrs/uuid"

type Users struct {
	ID         uuid.UUID `json:"id"           gorm:"type:char(36);not null;primaryKey"`
	Name       string    `json:"name"         gorm:"type:varchar(32);not null;unique"`
	IconFileID uuid.UUID `json:"icon_file_id" gorm:"type:char(36);not null"`
}
