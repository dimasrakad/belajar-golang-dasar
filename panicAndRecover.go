package main

import (
	"fmt"
)

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Terjadi error:", message)
	}
	fmt.Println("End app")
}

func PanicAndRecover(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
}
