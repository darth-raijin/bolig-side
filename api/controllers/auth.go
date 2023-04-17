package controllers

import (
	"time"

	errorDto "github.com/darth-raijin/bolig-side/api/models/dtos/error"
	loginUserDto "github.com/darth-raijin/bolig-side/api/models/dtos/user/login"
	registerUserDto "github.com/darth-raijin/bolig-side/api/models/dtos/user/register"
	"github.com/darth-raijin/bolig-side/pkg/service"
	"github.com/darth-raijin/bolig-side/pkg/utility"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Register a User
// @Summary Registers a user
// @Description DomainErrorCodes:
// @Description 2: Email is already in use
// @Description 3: Password not secure
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body registerUserDto.RegisterUserRequest true "User to register"
// @Sucess 201 {object} registerUserDto.RegisterUserResponse{}
// @Failure 422 {object} errorDto.DomainErrorWrapper{}
// @Router /api/auth/register [POST]
func RegisterUser(c *fiber.Ctx) error {
	payload := new(registerUserDto.RegisterUserRequest)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	validationError := validator.New().Struct(payload)
	if validationError != nil {
		validationErrors := validationError.(validator.ValidationErrors)

		// Check list and create a DomainErrorWrapper and return
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errorDto.DomainErrorWrapper{
			Timestamp: time.Now().UTC(),
			Errors:    utility.CreateValidationSlice(validationErrors),
		})
	}

	created, err := service.AuthService.CreateUser(*payload)

	if len(err.Errors) > 0 {
		return c.Status(err.Statuscode).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(created)
}

func GetRegisterView(c *fiber.Ctx) error {
	payload := new(loginUserDto.LoginUserRequest)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	validationError := validator.New().Struct(payload)
	if validationError != nil {
		validationErrors := validationError.(validator.ValidationErrors)

		// Check list and create a DomainErrorWrapper and return
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errorDto.DomainErrorWrapper{
			Timestamp: time.Now().UTC(),
			Errors:    utility.CreateValidationSlice(validationErrors),
		})
	}

	return nil
}
