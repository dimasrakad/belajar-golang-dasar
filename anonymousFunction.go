package main

import "fmt"

type Blacklist func(string) bool

func AnonymousFunction(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Hello", name)
	}
}
