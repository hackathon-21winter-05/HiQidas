package model

import "github.com/gofrs/uuid"

type Credential struct {
	UserID      uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	User        User      `gorm:"foreignkey:UserID"`
	MailAddress string    `gorm:"type:varchar(254);not null;unique"`
	HashedPass  string    `gorm:"type:varchar(32);not null"`
}
