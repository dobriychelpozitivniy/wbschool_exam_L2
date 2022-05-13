package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

func main() {
	bb := NewBurgerBuilder()

	p := NewPovar(&bb)

	v := p.Vopper()
	c := p.Cheeseburger()

	fmt.Println(v, c)
}

type Povar struct {
	bb BurgerBuilding
}

func (p *Povar) Vopper() Burger {
	p.bb.AddCheese()
	p.bb.AddMeat()
	p.bb.AddMeat()
	p.bb.AddMayonese()
	p.bb.AddSalat()
	p.bb.AddOnion()
	return p.bb.GetBurger()
}

func (p *Povar) Cheeseburger() Burger {
	p.bb.AddCheese()
	p.bb.AddMeat()
	p.bb.AddMayonese()
	return p.bb.GetBurger()
}

func NewPovar(bb BurgerBuilding) Povar {
	return Povar{
		bb: bb,
	}
}

type BurgerBuilding interface {
	AddCheese()
	AddMeat()
	AddSalat()
	AddMayonese()
	AddOnion()
	GetBurger() Burger
}

type BurgerBuilder struct {
	burger Burger
}

func NewBurgerBuilder() BurgerBuilder {
	return BurgerBuilder{
		burger: Burger{},
	}
}

func (bb *BurgerBuilder) AddCheese() {
	bb.burger.Cheese = +1
}

func (bb *BurgerBuilder) AddMeat() {
	bb.burger.Meat = +1
}
func (bb *BurgerBuilder) AddSalat() {
	bb.burger.Salat = +1
}
func (bb *BurgerBuilder) AddMayonese() {
	bb.burger.Mayonese = true
}
func (bb *BurgerBuilder) AddOnion() {
	bb.burger.Onion = +1
}
func (bb *BurgerBuilder) GetBurger() Burger {
	b := bb.burger
	bb.burger = Burger{}
	return b
}

type Burger struct {
	Cheese   int
	Meat     int
	Salat    int
	Mayonese bool
	Onion    int
}
