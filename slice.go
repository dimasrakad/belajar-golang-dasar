package main

import "fmt"

func Slice() {
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

	slice1 := months[3:9]
	fmt.Println(slice1)
	fmt.Println(slice1[0])
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice2 := months[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "New Month")
	fmt.Println(slice3)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Dimas"
	newSlice[1] = "Raka"
	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
