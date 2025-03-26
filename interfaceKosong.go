package main

import "fmt"

func Ups(value int) interface{} {
	if value == 1 {
		return 1
	} else if value == 2 {
		return true
	}
	return "Ups"
}

func InterfaceKosong() {
	fmt.Println(Ups(1))
	fmt.Println(Ups(2))
	fmt.Println(Ups(3))
}
