package main

import "fmt"

func main() {
	const firstName string = "Eko"
	const lastName = "Muliyo"

	const (
		firstNameBudi = "Budi"
		lastNameBudi  = "Budiman"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(firstNameBudi)
	fmt.Println(lastNameBudi)
}