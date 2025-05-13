package main

import "fmt"

func Ups() any {
	return "Test"
}

func main() {
	var a any = Ups()
	fmt.Println(a)
}