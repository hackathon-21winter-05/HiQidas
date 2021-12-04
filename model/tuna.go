package model

import "github.com/gofrs/uuid"

type Tuna struct {
	ID           uuid.UUID    `json:"id"            gorm:"type:char(36);not null;primaryKey"`
	HiqidashiOne uuid.UUID    `json:"hiqidashi_one" gorm:"type:char(36);not null"`
	HiqidashiTwo uuid.UUID    `json:"hiqidashi_two" gorm:"type:char(36);not null"`
	Hiqidashis   []Hiqidashi `json:"-"              gorm:"foreignKey:ID"`
}
