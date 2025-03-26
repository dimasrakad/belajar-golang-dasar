package main

import "fmt"

func IfExpression() {
	name := "Raka"

	if name == "Dims" {
		fmt.Println("Name Dims")
	} else if name == "Raka" {
		fmt.Println("Name Raka")
	} else {
		fmt.Println("Name not Dims or Raka")
	}

	if length := len(name); length > 5 {
		fmt.Println("Name length is greater than 5")
	} else {
		fmt.Println("Name length is lesser than 5")
	}
}
