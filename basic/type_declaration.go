package main

import "fmt"

func main() {

	type PhoneNumber string

	var phoneNumber PhoneNumber = "08123456789"

	var example string = "081234567891"
	var examplePhoneNumber = PhoneNumber(example)

	fmt.Println(phoneNumber)
	fmt.Println(example)
	fmt.Println(examplePhoneNumber)
}