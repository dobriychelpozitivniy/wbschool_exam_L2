package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

func main() {
	d := Device{Name: "MyDevice"}
	uds := UpdateDataService{Name: "UpdateService-1"}
	sds := SaveDataService{Name: "SaveDataService-1"}

	d.SetNext(&uds)
	uds.SetNext(&sds)

	data := Data{}
	d.Execute(&data)
}

type Service interface {
	Execute(*Data)
	SetNext(Service)
}

type Data struct {
	GetSource    bool
	UpdateSource bool
}

type Device struct {
	Name string
	Next Service
}

func (d *Device) Execute(data *Data) {
	if data.GetSource {
		fmt.Printf("Данные из устройства %s уже были получены \n", d.Name)
		d.Next.Execute(data)
		return
	}

	fmt.Printf("Получение данных из устройства %s \n", d.Name)
	data.GetSource = true
	d.Next.Execute(data)
}

func (d *Device) SetNext(s Service) {
	d.Next = s
}

type UpdateDataService struct {
	Name string
	Next Service
}

func (uds *UpdateDataService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Printf("Данные в сервисе %s уже были обновлены \n", uds.Name)
		uds.Next.Execute(data)
		return
	}

	fmt.Printf("Обновление данных из сервиса %s \n", uds.Name)
	data.UpdateSource = true
	uds.Next.Execute(data)
}

func (uds *UpdateDataService) SetNext(s Service) {
	uds.Next = s
}

type SaveDataService struct {
	Name string
	Next Service
}

func (sds *SaveDataService) Execute(data *Data) {
	if !data.UpdateSource {
		fmt.Println("Данные не были обновлены")
		return
	}

	fmt.Println("Данные сохранены")
}

func (sds *SaveDataService) SetNext(s Service) {
	sds.Next = s
}
