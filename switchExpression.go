package main

import "fmt"

func SwitchExpression() {
	name := "Dims"

	switch name {
	case "Dims":
		fmt.Println("Hello Dims")
	case "Raka":
		fmt.Println("Hello Raka")
	default:
		fmt.Println("Hello unknown")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Name is greater than 5")
	// case false:
	// 	fmt.Println("Name is lesser than 5")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Name is greater than 10")
	case length > 5:
		fmt.Println("Name is greater than 5")
	default:
		fmt.Println("Name is lesser than 5")
	}
}
