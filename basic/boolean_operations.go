package main

import "fmt"

func main() {
	var valueTotal = 90
	var absense = 81

	var passTotal bool = valueTotal > 80
	var passAbsence bool = absense > 80

	var pass = passTotal && passAbsence
	fmt.Println(pass)
}