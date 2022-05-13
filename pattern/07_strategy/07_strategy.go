package main

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

func main() {
	h := Hero{}

	h.Attack()

	s := NewSword()

	h.setWeapon(s)
	h.Attack()

	c := NewCar()
	h.setWeapon(c)
	h.Attack()
}

type Weapon interface {
	Attack()
}

type Hero struct {
	weapon Weapon
}

func (h *Hero) setWeapon(w Weapon) {
	h.weapon = w
}

func (h *Hero) Attack() {
	if h.weapon == nil {
		fmt.Println("Hero not have weapon")
		return
	}

	h.weapon.Attack()
}

type Sword struct {
}

func NewSword() *Sword {
	return &Sword{}
}

func (s *Sword) Attack() {
	fmt.Println("Attack with sword")
}

type Car struct {
}

func NewCar() *Car {
	return &Car{}
}

func (c *Car) Attack() {
	fmt.Println("Attack with car")
}
