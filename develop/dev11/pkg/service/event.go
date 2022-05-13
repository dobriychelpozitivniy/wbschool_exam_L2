package service

import (
	"dev11/pkg/model"
	"dev11/pkg/repository"
	"fmt"
	"time"
)

type EventService struct {
	r repository.Local
}

func NewEventService(r repository.Local) *EventService {
	return &EventService{r: r}
}

func (s *EventService) Create(event model.CreateEvent) (string, error) {
	return s.r.Add(event)
}

func (s *EventService) Update(event model.Event) (*model.Event, error) {
	return s.r.Update(event.ID, event)
}

func (s *EventService) Delete(id string) (string, error) {
	return s.r.Delete(id)
}

func (s *EventService) GetForDay(date string) (*[]model.Event, error) {
	var events *[]model.Event = &[]model.Event{}

	evs, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}

	for _, e := range *evs {
		if date == e.DateTodo {
			*events = append(*events, e)
		}
	}

	if len(*events) == 0 {
		return nil, fmt.Errorf("Empty result")
	}

	return events, nil
}

func (s *EventService) GetForWeek(numberWeek int, year int) (*[]model.Event, error) {
	var events *[]model.Event = &[]model.Event{}

	wStart := weekStart(year, numberWeek)

	wEnd := wStart.AddDate(0, 0, 7)

	evs, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}

	for _, e := range *evs {
		t, err := time.Parse("02-01-2006", e.DateTodo)
		if err != nil {
			return nil, fmt.Errorf("Error parse time: %s", err)
		}

		if t.After(wStart) && t.Before(wEnd) {
			*events = append(*events, e)
		}
	}

	if len(*events) == 0 {
		return nil, fmt.Errorf("Empty result")
	}

	return events, nil
}

func (s *EventService) GetForMonth(numberMonth int, year int) (*[]model.Event, error) {
	var events *[]model.Event = &[]model.Event{}

	mStart := monthStart(year, numberMonth)

	mEnd := mStart.AddDate(0, 1, 0)

	evs, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}

	fmt.Println(mStart)
	fmt.Println(mEnd)

	for _, e := range *evs {
		t, err := time.Parse("02-01-2006", e.DateTodo)
		if err != nil {
			return nil, fmt.Errorf("Error parse time: %s", err)
		}

		if t.After(mStart) && t.Before(mEnd) {
			*events = append(*events, e)
		}
	}

	if len(*events) == 0 {
		return nil, fmt.Errorf("Empty result")
	}

	return events, nil
}

func monthStart(year, month int) time.Time {
	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
}

func weekStart(year, week int) time.Time {
	// Start from the middle of the year:
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)

	// Roll back to Monday:
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	// Difference in weeks:
	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)

	return t
}
