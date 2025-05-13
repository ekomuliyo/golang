package main

import "fmt"

type Address struct {
	City string
	Province string
	Country string
}

func main() {

	address1 := Address{"Palembang", "Sumatera Selatan", "Indonesia"}
	address2 := &address1

	address2.City = "Jakarta"
	address2.Province = "DKI Jakarta"
	address2.Country = "Indonesia"
	fmt.Println(address1) 
	fmt.Println(address2) 

	//address2 = &Address{"Bandung", "Jawa Barat", "Indonesia"}
	*address2 = Address{"Bandung", "Jawa Barat", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)

}