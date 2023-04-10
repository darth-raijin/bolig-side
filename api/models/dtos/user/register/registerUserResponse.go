package registerUserDto

type RegisterUserResponse struct {
	ID             string `json:"id,omitempty"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Country        string `json:"country"`
	Password       string `json:"password"`
	Realtor        bool   `json:"realtor"`
	SubscriptionID string `json:"subscription_id,omitempty"`
}
