package controllers

import (
	"net/http"
	"time"

	errorDto "github.com/darth-raijin/bolig-side/api/models/dtos/error"

	createEventDto "github.com/darth-raijin/bolig-side/api/models/dtos/event/create"
	"github.com/darth-raijin/bolig-side/pkg/service"
	"github.com/darth-raijin/bolig-side/pkg/utility"
)

// Get a specific event by ID
// @Summary Get a specific event by ID
// @Description Get a specific event by ID
// @Tags Event
// @Accept json
// @Produce json
// @Param id query string true "id for fetching given event"
// @Failure 404 {object} errorDto.DomainErrorWrapper{}
// @Router /api/v1/events/{id} [get]
func GetEventById(c *fiber.Ctx) error {
	return c.Status(http.StatusServiceUnavailable).JSON(errorDto.DomainErrorWrapper{})
}

// Creates an event
// @Summary Creates an event
// @Description Creates an event
// @Tags Event
// @Accept json
// @Produce json
// @Param Event body dtos.Event true "Required parameters to create an event"
// @Success 201 {object} createEventDto.CreateEventRequest{}
// @Failure 419 {object} errorDto.DomainErrorWrapper{}
// @Router /api/v1/events [post]
func CreateEvent(c *fiber.Ctx) error {
	payload := new(createEventDto.CreateEventRequest)

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

	// Passing DTO to service class for handling
	event, err := service.EventService.CreateEvent(payload)

	if len(err.Errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	return c.Status(http.StatusCreated).JSON(event)
}
