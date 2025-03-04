package main

import "fmt"

// Інтерфейс Посередника
type AirTrafficControl interface {
	RegisterAircraft(aircraft Aircraft)
	SendMessage(message string, sender Aircraft)
}

// Конкретний Посередник
type Tower struct {
	aircrafts map[string]Aircraft
}

func NewTower() *Tower {
	return &Tower{aircrafts: make(map[string]Aircraft)}
}

func (t *Tower) RegisterAircraft(aircraft Aircraft) {
	t.aircrafts[aircraft.GetID()] = aircraft
}

func (t *Tower) SendMessage(message string, sender Aircraft) {
	for id, ac := range t.aircrafts {
		if id != sender.GetID() {
			ac.ReceiveMessage(message)
		}
	}
}

// Інтерфейс Літака
type Aircraft interface {
	GetID() string
	SendMessage(message string)
	ReceiveMessage(message string)
}

// Конкретна Реалізація Літака
type ConcreteAircraft struct {
	id  string
	atc AirTrafficControl
}

func NewConcreteAircraft(id string, atc AirTrafficControl) *ConcreteAircraft {
	return &ConcreteAircraft{id: id, atc: atc}
}

func (a *ConcreteAircraft) GetID() string {
	return a.id
}

func (a *ConcreteAircraft) SendMessage(message string) {
	fmt.Printf("Літак %s надсилає повідомлення: %s\n", a.id, message)
	a.atc.SendMessage(message, a)
}

func (a *ConcreteAircraft) ReceiveMessage(message string) {
	fmt.Printf("Літак %s отримав повідомлення: %s\n", a.id, message)
}

// Головна функція
func main() {
	tower := NewTower()

	aircraft1 := NewConcreteAircraft("AC001", tower)
	aircraft2 := NewConcreteAircraft("AC002", tower)
	aircraft3 := NewConcreteAircraft("AC003", tower)

	tower.RegisterAircraft(aircraft1)
	tower.RegisterAircraft(aircraft2)
	tower.RegisterAircraft(aircraft3)

	aircraft1.SendMessage("Запит на посадку.")
}
