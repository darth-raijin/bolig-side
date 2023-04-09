package createEventDto

type CreateEventRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=32"`
	Description string `json:"description" validate:"required,max=255" example:"Sample description"`
	City        string `json:"city,omitempty" validate:"max=255"`
	Country     string `json:"country,omitempty" validate:"iso3166_1_alpha2"`
}
