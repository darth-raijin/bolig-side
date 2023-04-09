package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Feedback struct {
	gorm  gorm.Model
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name  string
	Email string
}

type Question struct {
	gorm    gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Author  User      `gorm:"embedded"`
	Upvotes int32
}
