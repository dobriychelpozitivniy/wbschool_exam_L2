package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Вынести логику выбора создания новых объектов в один метод, который вызывает конструкторы с логикой создания
необходимых объектов в рантайме
*/

func main() {
	m, err := NewMemebership("g")
	if err != nil {
		panic(err)
	}

	fmt.Println(m.Name())
}

func NewMemebership(m string) (Membership, error) {
	switch m {
	case "g":
		return NewGymMembership(40), nil
	case "t":
		return NewPersonalTrainingMembership(100), nil
	case "p":
		return NewGymPlusPoolMembership(70), nil
	default:
		return nil, fmt.Errorf("Такого абонемента не существует")
	}
}

type Membership interface {
	Name() string
	Price() float64
}

type GymMembership struct {
	name  string
	price float64
}

func NewGymMembership(price float64) *GymMembership {
	return &GymMembership{
		name:  "Gym membership",
		price: price,
	}
}

func (gm *GymMembership) Name() string {
	return gm.name
}

func (gm *GymMembership) Price() float64 {
	return gm.price
}

type GymPlusPoolMembership struct {
	name  string
	price float64
}

func NewGymPlusPoolMembership(price float64) *GymPlusPoolMembership {
	return &GymPlusPoolMembership{
		name:  "Gym + pool membership",
		price: price,
	}
}

func (gm *GymPlusPoolMembership) Name() string {
	return gm.name
}

func (gm *GymPlusPoolMembership) Price() float64 {
	return gm.price
}

type PersonalTrainingMembership struct {
	name  string
	price float64
}

func NewPersonalTrainingMembership(price float64) *PersonalTrainingMembership {
	return &PersonalTrainingMembership{
		name:  "Personal Training membership",
		price: price,
	}
}

func (gm *PersonalTrainingMembership) Name() string {
	return gm.name
}

func (gm *PersonalTrainingMembership) Price() float64 {
	return gm.price
}
