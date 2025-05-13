package main

import (
	"fmt"
)

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "ID cannot be empty"}
	}
	if data == nil {
		return &notFoundError{Message: "Data not found"}
	}
	return nil
}

func main() {

	err := SaveData("", nil)
	if err != nil {
		if validationErr, ok := err.(*validationError); ok {
			// Handle validation error
			println("Validation Error:", validationErr.Message)
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			// Handle not found error
			println("Not Found Error:", notFoundErr.Message)
		} else {
			// Handle other errors
			println("Unknown Error:", err.Error())
		}
	} else {
		fmt.Println("Data saved successfully")
	}

}