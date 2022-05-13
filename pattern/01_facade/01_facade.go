package main

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// Минимизировать зависимость подсистем некоторой сложной системы и обмен информацией между ними.
//Паттерн «Фасад» предоставляет унифицированный интерфейс вместо набора интерфейсов некоторой подсистемы.
// Фасад определяет интерфейс более высокого уровня, кото-
// рый упрощает использование подсистемы.

func main() {
	f := Facade{
		db:      &DB{},
		http:    &HTTP{},
		service: &Service{},
	}

	f.SomeWork()
}

type Facade struct {
	db      *DB
	http    *HTTP
	service *Service
}

func (f *Facade) SomeWork() {
	f.db.SomeDBWork()
	f.http.SomeHTTPWork()
	f.service.SomeServiceWork()
}

type DB struct{}

func (db *DB) SomeDBWork() {
	// some code
}

type HTTP struct{}

func (http *HTTP) SomeHTTPWork() {
	// some code
}

type Service struct{}

func (s *Service) SomeServiceWork() {
	// some code
}
