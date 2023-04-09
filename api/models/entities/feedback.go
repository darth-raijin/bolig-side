package entities

import (
	"time"

	"github.com/google/uuid"
)

type Feedback struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Event    Event     `gorm:"foreignKey:ID"`
	Question Question  `gorm:"foreignKey:ID"`
	User     User      `gorm:"foreignKey:ID"`
	Created  time.Time `gorm:"autoCreateTime"`
	Updated  time.Time
}
