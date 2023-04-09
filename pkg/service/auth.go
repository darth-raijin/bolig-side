package service

import (
	"errors"
	"time"

	errorDto "github.com/darth-raijin/bolig-side/api/models/dtos/error"
	loginUserDto "github.com/darth-raijin/bolig-side/api/models/dtos/user/login"
	registerUserDto "github.com/darth-raijin/bolig-side/api/models/dtos/user/register"
	"github.com/darth-raijin/bolig-side/api/models/entities"
	"github.com/darth-raijin/bolig-side/pkg/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var AuthService authService

type authService struct{}

// Used for registering user
func (authService) CreateUser(user registerUserDto.CreateEventRequest) (registerUserDto.CreateUserResponse, errorDto.DomainErrorWrapper) {
	userEntity := entities.User{
		Email: user.Email,
		Name:  user.Name,
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {

		return registerUserDto.CreateUserResponse{}, errorDto.DomainErrorWrapper{
			Timestamp: time.Now(),
			Errors: []errorDto.DomainError{
				{
					Message: err.Error(),
				},
			},
		}
	}

	userEntity.Password = string(hashed)

	// Check if EMAIL and USERNAME is distinct
	result := repository.GormDB.Create(&userEntity)
	if result.Error != nil {
		// Handle database error
		if errors.Is(err, gorm.ErrDuplicatedKey) {
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

			return registerUserDto.CreateUserResponse{}, wrapper
		}

		return registerUserDto.CreateUserResponse{}, errorDto.DomainErrorWrapper{}
	}
	return registerUserDto.CreateUserResponse{}, errorDto.DomainErrorWrapper{}
}

func (authService) LoginUser(user loginUserDto.LoginUserRequest) loginUserDto.LoginUserResponse {

	return loginUserDto.LoginUserResponse{}
}

func (authService) ResetPassword() {

}

func (authService) FetchUserDetails() {

}
