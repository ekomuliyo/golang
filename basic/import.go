package main

import (
	"fmt"
	"example.com/golang/basic/helper"
)

func main() {

	result := helper.SayHello("Eko")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version)
	// fmt.Println(helper.sayGoodBy("Eko"))
}