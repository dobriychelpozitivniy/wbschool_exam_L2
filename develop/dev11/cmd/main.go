package main

import (
	"dev11/pkg/handler"
	"dev11/pkg/model"
	"dev11/pkg/repository"
	"dev11/pkg/service"
)

func main() {
	r := repository.NewRepository(map[string]model.Event{})
	s := service.NewService(r)
	h := handler.NewHandler(s)
	server := h.InitHandlers()

	server.RunServer(":8080")

}
