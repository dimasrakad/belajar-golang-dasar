package main

import "fmt"

func Nil() {
	person := NewMap("")

	if person == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(person)
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{
		"name": name,
	}
}
