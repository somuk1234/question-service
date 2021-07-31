package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Question struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;unique"`
	CreatedAt time.Time `gorm:"autoCreateTime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;" json:"updated_at"`
	Status    bool      `gorm:"type:bool;not null;" json:"status"`
	Link      string    `gorm:"type:VARCHAR(220);unique;not null;" json:"link"`
	Rating    int       `gorm:"type:INT;" json:"rating"`
}
