package dtos

import (
	"github.com/google/uuid"
)

type UserDto struct {
	ID    uuid.UUID `json:"id" validate:"required" format:"uuid"`
	Name  string    `json:"name" validate:"required" example:"Go Crash Course"`
	Email string    `json:"email" validate:"required"`
}
