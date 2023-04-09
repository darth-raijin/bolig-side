package dtos

import (
	"github.com/google/uuid"
)

type Event struct {
	ID        uuid.UUID  `json:"id" validate:"required" format:"uuid"`
	Name      string     `json:"name" validate:"required" example:"Go Crash Course"`
	Questions []Question `json:"questions,omitempty" validate:"required"`
}
