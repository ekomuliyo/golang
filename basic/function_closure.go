package main

import "fmt"

func main() {
	counter := 0
	increment := func() {
		fmt.Println("Increment")
		counter++
	}
	
	increment()
	fmt.Println(counter)
	increment()
	fmt.Println(counter)
	increment()
	fmt.Println(counter)
}