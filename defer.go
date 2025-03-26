package main

import "fmt"

func Defer(value int) {
	defer logging()
	fmt.Println("Run application")
	fmt.Println("Result:", 10/value)
}

func logging() {
	fmt.Println("Selesai memanggil function")
}
