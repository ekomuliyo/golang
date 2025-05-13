package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Eko"
	names[1] = "Budi"
	names[2] = "Andi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		1,
		2,
		3,
	}
	fmt.Println(values)

	var values2 = [...]int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(values2)
}