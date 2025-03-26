package main

import (
	"errors"
	"fmt"
)

func ErrorInterface() {
	hasil, err := Pembagian(5, 0)

	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Hasil:", hasil)
	}
}

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan nol")
	} else {
		return nilai / pembagi, nil
	}
}
