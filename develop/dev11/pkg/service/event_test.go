package service

import (
	"dev11/pkg/model"
	"dev11/pkg/repository"
	"reflect"
	"testing"
)

func initService() *Service {
	r := repository.NewRepository(map[string]model.Event{"12we12": model.Event{
		ID:          "12we12",
		Name:        "test",
		Description: "test",
		DateAdded:   "21-04-2022",
		DateTodo:    "22-04-2022",
	}})

	s := NewService(r)

	return s
}

func initEmptySerivce() *Service {
	r := repository.NewRepository(map[string]model.Event{})
	s := NewService(r)

	return s
}

func TestEventService_Create(t *testing.T) {
	tests := []struct {
		name    string
		s       *EventService
		event   model.CreateEvent
		wantErr bool
	}{
		{
			name: "OK",
			s:    initEmptySerivce().Event.(*EventService),
			event: model.CreateEvent{
				Name:        "test",
				Description: "test",
				DateAdded:   "21-04-2022",
				DateTodo:    "22-04-2022",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Create(tt.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("EventService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" {
				t.Errorf("EventService.Create() = %v, want %v", got, "not empty id")
			}
		})
	}
}

func TestEventService_Update(t *testing.T) {
	tests := []struct {
		name    string
		s       *EventService
		event   model.Event
		want    *model.Event
		wantErr bool
	}{
		{
			name: "OK",
			s:    initService().Event.(*EventService),
			event: model.Event{
				ID:          "12we12",
				Name:        "newname",
				Description: "test",
				DateAdded:   "21-04-2022",
				DateTodo:    "22-04-2022",
			},
			want: &model.Event{
				ID:          "12we12",
				Name:        "newname",
				Description: "test",
				DateAdded:   "21-04-2022",
				DateTodo:    "22-04-2022",
			},
			wantErr: false,
		},
		{
			name: "not exist id",
			s:    initService().Event.(*EventService),
			event: model.Event{
				ID:          "not exist id",
				Name:        "newname",
				Description: "test",
				DateAdded:   "21-04-2022",
				DateTodo:    "22-04-2022",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Update(tt.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("EventService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EventService.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventService_Delete(t *testing.T) {
	tests := []struct {
		name    string
		s       *EventService
		id      string
		want    string
		wantErr bool
	}{
		{
			name:    "OK",
			s:       initService().Event.(*EventService),
			id:      "12we12",
			want:    "12we12",
			wantErr: false,
		},
		{
			name:    "not exist id",
			s:       initService().Event.(*EventService),
			id:      "not exist id",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Delete(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("EventService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EventService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventService_GetForDay(t *testing.T) {
	tests := []struct {
		name    string
		s       *EventService
		date    string
		want    *[]model.Event
		wantErr bool
	}{
		{
			name: "OK",
			s:    initService().Event.(*EventService),
			date: "22-04-2022",
			want: &[]model.Event{{
				ID:          "12we12",
				Name:        "test",
				Description: "test",
				DateAdded:   "21-04-2022",
				DateTodo:    "22-04-2022",
			}},
			wantErr: false,
		},
		{
			name:    "empty result",
			s:       initService().Event.(*EventService),
			date:    "22-04-2050",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetForDay(tt.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("EventService.GetForDay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EventService.GetForDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventService_GetForWeek(t *testing.T) {
	type args struct {
		numberWeek int
		year       int
	}
	tests := []struct {
		name    string
		s       *EventService
		args    args
		want    *[]model.Event
		wantErr bool
	}{
		{
			name: "ok",
			s:    initService().Event.(*EventService),
			args: args{
				year:       2022,
				numberWeek: 16,
			},
			want: &[]model.Event{{
				ID:          "12we12",
				Name:        "test",
				Description: "test",
				DateAdded:   "21-04-2022",
				DateTodo:    "22-04-2022",
			}},
			wantErr: false,
		},
		{
			name: "empty result",
			s:    initService().Event.(*EventService),
			args: args{
				year:       2022,
				numberWeek: 17,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetForWeek(tt.args.numberWeek, tt.args.year)
			if (err != nil) != tt.wantErr {
				t.Errorf("EventService.GetForWeek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EventService.GetForWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventService_GetForMonth(t *testing.T) {
	type args struct {
		numberMonth int
		year        int
	}
	tests := []struct {
		name    string
		s       *EventService
		args    args
		want    *[]model.Event
		wantErr bool
	}{
		{
			name: "ok",
			s:    initService().Event.(*EventService),
			args: args{
				year:        2022,
				numberMonth: 4,
			},
			want: &[]model.Event{{
				ID:          "12we12",
				Name:        "test",
				Description: "test",
				DateAdded:   "21-04-2022",
				DateTodo:    "22-04-2022",
			}},
			wantErr: false,
		},
		{
			name: "empty result",
			s:    initService().Event.(*EventService),
			args: args{
				year:        2022,
				numberMonth: 12,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetForMonth(tt.args.numberMonth, tt.args.year)
			if (err != nil) != tt.wantErr {
				t.Errorf("EventService.GetForMonth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EventService.GetForMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}
