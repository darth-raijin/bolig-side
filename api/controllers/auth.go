package controllers

import (
	"net/http"
	"time"

	errorDto "github.com/darth-raijin/bolig-side/api/models/dtos/error"
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
// @Failure 422 {object} errorDto.DomainErrorWrapper{}
// @Router /api/auth/register [POST]
func RegisterUser(c *fiber.Ctx) error {
	payload := new(registerUserDto.CreateEventRequest)

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

	_, err := service.AuthService.CreateUser(*payload)

	if len(err.Errors) > 0 {
		return c.Status(err.Statuscode).JSON(err)
	}
	return c.Status(http.StatusServiceUnavailable).JSON(errorDto.DomainErrorWrapper{})
}
