package main

import "fmt"

type Person struct {
	Name string
}


func (person Person) GetName() string {
	return person.Name
}

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{"Eko"}
	SayHello(person)

	animal := Animal{"Dog"}
	SayHello(animal)
}