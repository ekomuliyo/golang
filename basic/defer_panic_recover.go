package main

import "fmt"

func logging() {
	fmt.Println("Logging")
}

func runApp() {
	defer logging() // This will be executed after runApp() is finished
	fmt.Println("Run App")
}

func endApp() {
	fmt.Println("End App")
	message := recover() // This will recover from panic
	if message != nil {
		fmt.Println("Error:", message)
	}
}

func runAppWithPanic(error bool) {
	defer endApp() // This will be executed after runAppWithPanic() is finished
	if error {
		panic("Application Error")
	}
	fmt.Println("Run App With Panic")
}

func main() {
	// runApp()
	runAppWithPanic(true)
}