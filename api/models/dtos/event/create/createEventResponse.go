package createEventDto

import "github.com/google/uuid"

type CreateEventResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	City        string    `json:"city,omitempty"`
	Country     string    `json:"country,omitempty"`
}
