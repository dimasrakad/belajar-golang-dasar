package main

import (
	"fmt"
)

func ForLoops() {
	counter := 1

	for counter <= 10 {
		fmt.Print(counter, " ")
		counter++
	}

	fmt.Println()

	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()

	slice := []string{"Dims", "Raka", "Dewa"}

	for k, v := range slice {
		fmt.Println(k, v)
	}

	person := make(map[string]string)
	person["name"] = "Dims"
	person["age"] = "8"

	for k, v := range person {
		fmt.Println(k, ":", v)
	}
}
