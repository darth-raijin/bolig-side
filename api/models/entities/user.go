package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key"`
	Username      string    `gorm:"unique_index;not null"`
	Name          string
	Email         string    `gorm:"unique_index;not null"`
	Password      string    // Do some crazy HASH + SALT
	Hosted        []Event   `gorm:"foreignKey:ID"`
	Participating []Event   `gorm:"foreignKey:ID"`
	Created       time.Time `gorm:"autoCreateTime"`
}
