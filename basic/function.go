package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) {
	return "Eko", "Muliyo"
}

func getCompleteName() (firstName string, lastName string) {
	firstName = "Eko"
	lastName = "Muliyo"
	return firstName, lastName
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func getGoodbye(name string) string {
	return "Goodbye " + name
}

func main() {
	sayHello("Eko")
	sayHello("Budi")
	sayHello("Andi")
	sayHello("Muliyo")

	result := getHello("Eko")
	fmt.Println(result)

	firstName, _ := getFullName()
	fmt.Println(firstName)

	a, b := getFullName()
	fmt.Println(a, b)

	sum := sumAll(1, 2, 3, 4, 5)
    fmt.Println(sum)

    numbers := []int{1, 2, 3, 4, 5}
    sum = sumAll(numbers...)
    fmt.Println(sum)

	goodbye := getGoodbye("Eko")
	fmt.Println(goodbye)
}