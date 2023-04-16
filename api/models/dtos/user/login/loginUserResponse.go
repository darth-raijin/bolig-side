package loginUserDto

import (
	"time"
)

type LoginUserResponse struct {
	FirstName    string    `json:"firstname,omitempty"`
	LastName     string    `json:"lastname,omitempty"`
	Email        string    `json:"email,omitempty"`
	Country      string    `json:"country,omitempty"`
	Password     string    `json:"password,omitempty"`
	Realtor      bool      `json:"realtor,omitempty"`
	LastLoggedIn time.Time `json:"lastLoggedIn,omitempty"`
	AccessToken  string    `json:"access_token,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
}
