package service

import (
	"dev11/pkg/model"
	"dev11/pkg/repository"
)

type Event interface {
	Create(event model.CreateEvent) (string, error)
	Update(event model.Event) (*model.Event, error)
	GetForDay(date string) (*[]model.Event, error)
	GetForWeek(numberWeek int, year int) (*[]model.Event, error)
	GetForMonth(numberMonth int, year int) (*[]model.Event, error)
	Delete(id string) (string, error)
}

type Service struct {
	Event
}

func NewService(r *repository.Repository) *Service {
	return &Service{Event: NewEventService(r.Local)}
}
