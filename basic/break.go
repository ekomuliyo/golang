package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Loop index -",i)
	}
	fmt.Println("Loop terminated")

	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println("Loop index -",i)
	}
	fmt.Println("Loop terminated")

}