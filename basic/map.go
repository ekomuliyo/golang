package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "John Doe",
		"address": "Jakarta",
	}

	fmt.Println("Name:", person["name"])
	fmt.Println("Address:", person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Book Golang"
	book["author"] = "Eko"
	book["ups"] = "wrong"

	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)
}