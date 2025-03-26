package main

import "fmt"

func Array() {
	var names [3]string
	names[0] = "Dimas"
	names[1] = "Raka"
	names[2] = "Dewanggana"

	fmt.Println(names)

	var values = [3]int{
		1, 2, 3,
	}

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
}
