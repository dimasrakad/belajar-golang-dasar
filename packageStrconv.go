package main

import (
	"fmt"
	"strconv"
)

func PackageStrconv() {
	boolean, err := strconv.ParseBool("true")

	if err != nil {
		fmt.Println("Error :", err.Error())
	} else {
		fmt.Println(boolean)
	}

	int, err := strconv.ParseInt("100", 10, 32)

	if err != nil {
		fmt.Println("Error :", err.Error())
	} else {
		fmt.Println(int)
	}

	int2 := strconv.FormatInt(100, 2) // binary
	fmt.Println(int2)
	int2 = strconv.FormatInt(100, 8) // octal
	fmt.Println(int2)
	int2 = strconv.FormatInt(100, 16) // hexadecimal
	fmt.Println(int2)

	int3, error := strconv.Atoi("100")

	if err != nil {
		fmt.Println("Error :", error.Error())
	} else {
		fmt.Println(int3)
	}
}
