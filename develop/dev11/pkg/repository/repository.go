package repository

import "dev11/pkg/model"

type Local interface {
	Add(event model.CreateEvent) (string, error)
	Update(id string, event model.Event) (*model.Event, error)
	Delete(id string) (string, error)
	Get(id string) (*model.Event, error)
	GetAll() (*map[string]model.Event, error)
}

type Repository struct {
	Local
}

func NewRepository(es map[string]model.Event) *Repository {
	return &Repository{Local: NewLocalRepository(es)}
}
