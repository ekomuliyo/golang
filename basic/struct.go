package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {

	var customer1 Customer
	customer1.Name = "Eko"
	customer1.Address = "Jakarta"
	customer1.Age = 30
	fmt.Println(customer1)

	customer1.Name = "Budi"
	customer1.Address = "Bandung"
	customer1.Age = 25
	fmt.Println(customer1)

	andi := Customer{
		Name:    "Andi",
		Address: "Semarang",
		Age:     28,
	}
	fmt.Println(andi)

	santoso := Customer{"Santoso", "Yogyakarta", 35}
	fmt.Println(santoso)

	santoso.sayHello("Eko")
}