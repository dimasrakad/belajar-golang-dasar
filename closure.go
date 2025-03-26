package main

import "fmt"

func Closure() {
	counter := 0
	fmt.Println("Counter: ", counter)
	increment := func() {
		fmt.Println("Increment")
		counter++
	}
	increment()
	fmt.Println("Counter: ", counter)
	increment()
	fmt.Println("Counter: ", counter)
}
