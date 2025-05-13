package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = 30
	var d = a + b * c
	fmt.Println(d)

	var i = 10
	i += 5
	fmt.Println(i)

	i += 5
	fmt.Println(i)

	var j = 10
	j++
	fmt.Println(j)
	j--
	fmt.Println(j)
}