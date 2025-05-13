package main

import "fmt"

func factorialRecursive(value int) int {
	if value == 0 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}


func main() {
	fmt.Println(factorialRecursive(10))
}