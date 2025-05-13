package main

import "fmt"

func main() {
	var value32 int32 = 1000
	var value64 int64 = int64(value32)
	var value16 int16 = int16(value32)

	fmt.Println("value32=", value32)
	fmt.Println("value64=", value64)
	fmt.Println("value16=", value16)

	var name = "Eko"
	var e uint8 = name[0]
	var eString string = string(e)
	fmt.Println("e=", e)
	fmt.Println("e=", eString)
}