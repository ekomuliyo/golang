package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{
		"name": name,
	}
}

func main() {

	data := NewMap("")
	if data == nil {
		fmt.Println("Data is nil")
	} else {
		fmt.Println("Data is not nil")
	}

}