package main

import "fmt"

func BreakAndContinue() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Print(i, " ")
	}

	fmt.Println()

	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Print(i, " ")
	}
}
