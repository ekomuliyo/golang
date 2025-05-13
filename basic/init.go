package main

import (
	"fmt"
	"example.com/golang/basic/database"
	_ "example.com/golang/basic/internal"

)

func main() {
	fmt.Println(database.GetDatabase())
}