package main

import "fmt"

type Filter func(string) string

func FunctionAsParameter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func toxicFilter(name string) string {
	if name == "Anjing" {
		return "*****"
	}
	return name
}
