package main

import "fmt"

func random() interface{} {
	return "OK"
}

func TypeAssertions() {
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("String:", value)
	case int:
		fmt.Println("Int:", value)
	default:
		fmt.Println("Unknown")
	}
}
