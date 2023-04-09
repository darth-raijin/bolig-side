package service

import (
	errorDto "github.com/darth-raijin/bolig-side/api/models/dtos/error"
	createEventDto "github.com/darth-raijin/bolig-side/api/models/dtos/event/create"
	"github.com/darth-raijin/bolig-side/api/models/entities"
	"github.com/darth-raijin/bolig-side/pkg/repository"
)

var EventService eventService

type eventService struct{}

func (eventService) CreateEvent(*createEventDto.CreateEventRequest) (createEventDto.CreateEventResponse, errorDto.DomainErrorWrapper) {
	entity := entities.Event{
		Name:        "Sample Event",
		Description: "Sample Description",
		City:        "Copenhagen",
		Country:     "Denmark",
		Questions:   []entities.Question{},
	}
	// DB transaction here -> pass err from here to return err
	repository.GormDB.Create(&entity)

	return createEventDto.CreateEventResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		City:        entity.City,
		Country:     entity.Country,
	}, errorDto.DomainErrorWrapper{}
}

func (eventService) FetchEvent(*createEventDto.CreateEventRequest) (createEventDto.CreateEventResponse, errorDto.DomainErrorWrapper) {

	return createEventDto.CreateEventResponse{}, errorDto.DomainErrorWrapper{}
}

func (eventService) UpdateEvent(*createEventDto.CreateEventRequest) (createEventDto.CreateEventResponse, errorDto.DomainErrorWrapper) {

	return createEventDto.CreateEventResponse{}, errorDto.DomainErrorWrapper{}
}

func (eventService) FetchQuestions(*createEventDto.CreateEventRequest) (createEventDto.CreateEventResponse, errorDto.DomainErrorWrapper) {
	return createEventDto.CreateEventResponse{}, errorDto.DomainErrorWrapper{}
}
