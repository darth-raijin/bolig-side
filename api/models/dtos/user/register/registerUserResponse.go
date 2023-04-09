package registerUserDto

type CreateUserResponse struct {
	Username string `json:"username" validate:"required, min=3, max=12" example:"DarthMaul1337"`
	Name     string `json:"name" validate:"required" example:"Name"`
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
}
