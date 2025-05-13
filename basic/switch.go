package main

import "fmt"

func main() {
	name := "Eko"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Budi":
		fmt.Println("Hello Budi")
	case "Andi":
		fmt.Println("Hello Andi")
	default:
		fmt.Println("Default")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name is too long")
	case false:
		fmt.Println("Name is short")
	}

	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Name is too long")
	case length < 5:
		fmt.Println("Name is short")
	default:
		fmt.Println("Name is exactly 5 characters")
	}
	
}