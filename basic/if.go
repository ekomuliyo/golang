package main

import "fmt"

func main() {
	name := "Eko"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hello Unknown")
	}

	if length := len(name); length > 5 {
		fmt.Println("Name is too long")
	} else {
		fmt.Println("Name is short")
	}
}