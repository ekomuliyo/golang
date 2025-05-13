package main

import "fmt"

func main() {

	counter := 1

	for counter <= 10 {
		fmt.Println("Counter:", counter)
		counter++
	}
	fmt.Println("Done")

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Counter:", counter)
	}
	fmt.Println("Done")


	names := [5]string{"Eko", "Muliyo", "Budi", "Budiman", "Andi"}
	for i := 0; i < len(names); i++ {
		fmt.Println("Name:", names[i])
	}

	for i, name := range names {
		fmt.Println("Index:", i, "Name:", name)
	}

	for _, name := range names {
		fmt.Println("Name:", name)
	}
}