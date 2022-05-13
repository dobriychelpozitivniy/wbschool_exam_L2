package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// Необходимо иметь эффективное представление запросов к некоторой системе,
// не обладая при этом знаниями ни об их природе ни о способах их обработки.

func main() {

	light := NewLightOutside(1)
	heater := NewHeatingCooling(22.5)

	lightOnCommand := NewSwitchOnLightCommand(light)
	heatCommand := NewStartHeatingCommand(heater)

	eveningProgramm := NewProgramm(make([]Command, 0))
	eveningProgramm.AppendCmd(lightOnCommand)
	eveningProgramm.AppendCmd(heatCommand)
	eveningProgramm.start()
}

type LightOutside struct {
	intensity float64
}

func NewLightOutside(i float64) LightOutside {
	return LightOutside{
		intensity: i,
	}
}

func (lo *LightOutside) switchOn() {
	fmt.Println("Ligth's switched on")
}

func (lo *LightOutside) switchOff() {
	fmt.Println("Light's switched off")
}

type HeatingCooling struct {
	temperature float64
}

func NewHeatingCooling(t float64) HeatingCooling {
	return HeatingCooling{
		temperature: t,
	}
}

func (hc *HeatingCooling) mode() string {
	if hc.temperature >= 25 {
		return "heating"
	}
	return "cooling"
}

func (hc *HeatingCooling) start() {
	fmt.Printf("Start %s", hc.mode())
}

func (hc *HeatingCooling) stop() {
	fmt.Printf("Stop %s", hc.mode())
}

type Command interface {
	execute()
}

type SwitchOnLightCommand struct {
	light LightOutside
}

func (son SwitchOnLightCommand) execute() {
	son.light.switchOn()
}

func NewSwitchOnLightCommand(l LightOutside) SwitchOnLightCommand {
	return SwitchOnLightCommand{
		light: l,
	}
}

type StartHeatingCommand struct {
	heater HeatingCooling
}

func (shc StartHeatingCommand) execute() {
	if shc.heater.temperature < 25 {
		shc.heater.temperature = 25
	}
	shc.heater.start()
}

func NewStartHeatingCommand(h HeatingCooling) StartHeatingCommand {
	return StartHeatingCommand{
		heater: h,
	}
}

type Programm struct {
	commands []Command
}

func NewProgramm(cmds []Command) Programm {
	return Programm{
		commands: cmds,
	}
}

func (p *Programm) AppendCmd(c Command) {
	p.commands = append(p.commands, c)
}

func (p *Programm) start() {
	for _, v := range p.commands {
		v.execute()
	}
}
