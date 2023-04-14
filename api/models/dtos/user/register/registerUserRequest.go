package registerUserDto

type RegisterUserRequest struct {
	FirstName        string `json:"first_name" validate:"required"`
	LastName         string `json:"last_name" validate:"required"`
	Email            string `json:"email" validate:"required,email"`
	Country          string `json:"country" validate:"required,len=2"`
	Password         string `json:"password" validate:"required,min=8"`
	RepeatedPassword string `json:"repeatedPassword" validate:"required,min=8"`
	Realtor          bool   `json:"realtor"`
	SubscriptionID   string `json:"subscription_id,omitempty"`
}
