package handler

import (
	"dev11/pkg/router"
	"dev11/pkg/service"
)

type Handler struct {
	s *service.Service
	router.Router
}

func (h *Handler) InitHandlers() *router.Router {
	rr := router.NewRouter()

	rr.RegisterMethod("POST", "/create_event", h.CreateEvent)
	rr.RegisterMethod("POST", "/update_event", h.UpdateEvent)
	rr.RegisterMethod("POST", "/delete_event", h.DeleteEvent)
	rr.RegisterMethod("GET", "/events_for_day", h.GetEventsForDay)
	rr.RegisterMethod("GET", "/events_for_week", h.GetEventsForWeek)
	rr.RegisterMethod("GET", "/events_for_month", h.GetEventsForMonth)
	rr.RegisterMiddleware(h.Logger)

	return rr
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{s: s}
}
