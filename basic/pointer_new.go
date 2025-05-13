package main

import "fmt"

type Address struct {
	City string
	Province string
	Country string
}

func main() {

	address1 := new(Address)
	address2 := address1

	address2.City = "Jakarta"

	fmt.Println(address1)
	fmt.Println(address2)

}