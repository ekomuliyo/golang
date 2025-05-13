package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	name = filter(name)
	fmt.Println("Hello", name)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "*****"
	}
	return name
}

func main() {
	sayHelloWithFilter("Eko", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}