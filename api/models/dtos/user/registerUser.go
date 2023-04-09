package dtos

type RegisterUserDto struct {
	Username string `json:"username" validate:"required, min=3, max=12" example:"Random Name"`
	Name     string `json:"name" validate:"required" example:"Username"`
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
}
