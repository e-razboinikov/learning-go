package main

import "log"

// Объявление интрфейса
type Animal interface {
	Says() string
	NumberOfLegs() int
}

// Объявление полиаморфной функции
func PrintInfo(animal Animal) {
	log.Println("This animal says", animal.Says(), "and has", animal.NumberOfLegs(), "legs")
}

// Объявление типа данных
type Dog struct {
	Name  string
	Breed string
}

// Имплементация функции интерфейса
func (dog Dog) Says() string {
	return "woof"
}

// Имплементация функции интерфейса
func (dog Dog) NumberOfLegs() int {
	return 4
}

// Объявление типа данных
type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

// Имплементация функции интерфейса
func (gorilla Gorilla) Says() string {
	return "argh"
}

// Имплементация функции интерфейса
func (gorilla Gorilla) NumberOfLegs() int {
	return 2
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepherd",
	}

	gorilla := Gorilla{
		Name:          "Gorilla",
		Color:         "Black",
		NumberOfTeeth: 32,
	}

	// Вызов полиаморфной функции
	PrintInfo(dog)
	PrintInfo(gorilla)
}
