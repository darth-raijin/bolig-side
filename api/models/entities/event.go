package entities

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	CreatedBy   string
	Name        string
	Description string
	City        string
	Country     string
	Questions   []Question `gorm:"foreignKey:ID"`
	Created     time.Time  `gorm:"autoCreateTime"`
	Updated     time.Time
}
