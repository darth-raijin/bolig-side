package registerUserDto

type CreateEventRequest struct {
	Email    string `json:"email" validate:"required, email"`
	Name     string `json:"name" validate:"required" example:"Name"`
	Password string `json:"password" validate:"required"`
}
