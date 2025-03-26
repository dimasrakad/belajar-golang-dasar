package main

import "fmt"

func Variable() {
	var name string

	name = "Dims"
	fmt.Println(name)

	name = "Dimas"
	fmt.Println(name)

	var age = 20
	fmt.Println(age)

	birthYear := 2000
	fmt.Println(birthYear)

	var (
		firstName = "Dimas"
		lastName  = "Raka"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
