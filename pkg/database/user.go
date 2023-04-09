package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm     gorm.Model
	ID       string `gorm:"primaryKey"`
	Name     string
	Username string
	Password string
}
