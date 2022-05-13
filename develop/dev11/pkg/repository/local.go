package repository

import (
	"crypto/rand"
	"crypto/sha1"
	"dev11/pkg/model"
	"fmt"
	"sync"
)

type LocalRepository struct {
	mu     sync.Mutex
	events map[string]model.Event
}

func NewLocalRepository(es map[string]model.Event) *LocalRepository {
	return &LocalRepository{events: es}
}

func (r *LocalRepository) Add(event model.CreateEvent) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := createID()

	_, ok := r.events[id]
	if ok {
		return "", fmt.Errorf("Error create new ID")
	}

	e := model.Event{
		ID:          id,
		Name:        event.Name,
		Description: event.Description,
		DateAdded:   event.DateAdded,
		DateTodo:    event.DateTodo,
	}

	r.events[id] = e

	return id, nil
}

func (r *LocalRepository) Update(id string, event model.Event) (*model.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.events[id]
	if !ok {
		return nil, fmt.Errorf("Event with id: %s not exist", id)
	}

	r.events[id] = event

	return &event, nil
}

func (r *LocalRepository) Delete(id string) (string, error) {
	_, ok := r.events[id]
	if !ok {
		return "", fmt.Errorf("Event with id: %s not exist", id)
	}

	delete(r.events, id)

	return id, nil
}

func (r *LocalRepository) GetAll() (*map[string]model.Event, error) {
	r.mu.Lock()

	defer r.mu.Unlock()

	if len(r.events) == 0 {
		return nil, fmt.Errorf("Empty events")
	}

	return &r.events, nil
}

func (r *LocalRepository) Get(id string) (*model.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	event, ok := r.events[id]
	if !ok {
		return nil, fmt.Errorf("Event with id: %s not exist", id)
	}

	return &event, nil
}

func createID() string {
	data := make([]byte, 10)

	_, _ = rand.Read(data)

	id := fmt.Sprintf("%x", sha1.Sum(data))

	return id
}
