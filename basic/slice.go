package main

import "fmt"

func main() {
	months := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	slice1 := months[2:6]
	fmt.Println(slice1) // [March April May June]

	slice2 := months[4:]
	fmt.Println(slice2) // [May June July August September October November December]

	slice3 := months[:5]
	fmt.Println(slice3) // [January February March April May]

	days := [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}

	daysSlice1 := days[5:]
	fmt.Println(daysSlice1) // [Friday Saturday]

	daysSlice1[0] = "Sunday Changed"
	daysSlice1[1] = "Monday Changed"
	fmt.Println(daysSlice1) // [Sunday Changed Monday Changed]
	fmt.Println(days) // [Sunday Monday Tuesday Wednesday Thursday Sunday Changed Monday Changed]

	daysSlice2 := append(daysSlice1, "Holiday")
	daysSlice2[0] = "Sunday"
	fmt.Println(daysSlice1) // [Sunday Changed Monday Changed]
	fmt.Println(daysSlice2) // [Sunday Monday Changed Holiday]
	fmt.Println(days) // [Sunday Monday Tuesday Wednesday Thursday Sunday Changed Monday Changed]


	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Eko"
	newSlice[1] = "Budi"
	fmt.Println(newSlice) // [Eko Budi]
	fmt.Println(len(newSlice)) // 2
	fmt.Println(cap(newSlice)) // 5

	newSlice2 := append(newSlice, "Andi")
	fmt.Println(newSlice2) // [Eko Budi Andi]
	fmt.Println(len(newSlice2)) // 3
	fmt.Println(cap(newSlice2)) // 5

	newSlice3 := append(newSlice, "Santoso")
	newSlice3[0] = "Eko Ubah"
	fmt.Println(newSlice3) // [Eko Ubah Budi Andi]
	fmt.Println(newSlice) // [Eko Budi]

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(fromSlice) // [Sunday Monday Tuesday Wednesday Thursday Sunday Changed Monday Changed]
	fmt.Println(toSlice) // [Sunday Monday Tuesday Wednesday Thursday Sunday Changed Monday Changed]


	thisIsArray := [...]int{
		1,
		2,
		3,
		4,
	}

	thisIsSlice := []int{
		1,
		2,
		3,
		4,
	}

	fmt.Println(thisIsArray) // [1 2 3 4]
	fmt.Println(thisIsSlice) // [1 2 3 4]

}