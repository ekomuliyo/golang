package main

import "fmt"

func random() interface{} {
	return "Eko"
}

func main() {

	result := random()

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	case bool:
		fmt.Println("Bool", value)
	default:
		fmt.Println("Unknown", value)
	}
}