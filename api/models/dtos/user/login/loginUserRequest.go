package loginUserDto

type LoginUserRequest struct {
	Username string `json:"username" validate:"min=3, max=12" example:"DarthMaul1337"`
	Email    string `json:"email" validate:" email"`
	Password string `json:"password" validate:"required"`
}
