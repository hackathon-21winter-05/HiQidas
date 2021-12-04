package model

import "github.com/gofrs/uuid"

type Histories struct {
	UserID  uuid.UUID `json:"user_id"  gorm:"type:char(36);not null;primaryKey"`
	SheetID uuid.UUID `json:"sheet_id" gorm:"type:char(36);not null;primaryKey"`
}
