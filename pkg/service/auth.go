package service

import (
	"time"

	errorDto "github.com/darth-raijin/bolig-side/api/models/dtos/error"
	loginUserDto "github.com/darth-raijin/bolig-side/api/models/dtos/user/login"
	registerUserDto "github.com/darth-raijin/bolig-side/api/models/dtos/user/register"
	"github.com/darth-raijin/bolig-side/api/models/entities"
	entityrepository "github.com/darth-raijin/bolig-side/pkg/repository/entityRepository"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var AuthService authService

type authService struct{}

// Used for registering user
func (authService) CreateUser(user registerUserDto.RegisterUserRequest) (registerUserDto.RegisterUserResponse, errorDto.DomainErrorWrapper) {
	userEntity := entities.User{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Country:   user.Country,
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	// Error occured trying to Hash the password
	if err != nil {
		return registerUserDto.RegisterUserResponse{}, errorDto.DomainErrorWrapper{
			Statuscode: fiber.StatusInternalServerError,
			Timestamp:  time.Now(),
			Errors: []errorDto.DomainError{
				{
					Message: "Failed handling password",
				},
			},
		}
	}

	userEntity.Password = string(hashed)
	registerResult, err := entityrepository.RegisterUser(userEntity)

	if err.Error() == errorDto.EmailNotUnique.Message {
		wrapper := errorDto.DomainErrorWrapper{
			Statuscode: 409,
			Timestamp:  time.Now().UTC(),
			Errors: []errorDto.DomainError{
				{
					DomainErrorCode: errorDto.EmailNotUnique.DomainErrorCode,
					Message:         err.Error(),
				},
			},
		}
		return registerUserDto.RegisterUserResponse{}, wrapper
	}

	response := registerUserDto.RegisterUserResponse{
		ID:        registerResult.ID.String(),
		FirstName: registerResult.FirstName,
		LastName:  registerResult.LastName,
		Email:     registerResult.Email,
		Country:   registerResult.Country,
		Realtor:   registerResult.Realtor,
	}
	return response, errorDto.DomainErrorWrapper{}
}

func (authService) LoginUser(user loginUserDto.LoginUserRequest) loginUserDto.LoginUserResponse {

	return loginUserDto.LoginUserResponse{}
}

func (authService) ResetPassword() {

}

func (authService) FetchUserDetails() {

}
