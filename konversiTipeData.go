package main

import "fmt"

func KonversiTipeData() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	name := "Dims"
	d := name[0]
	dString := string(d)

	fmt.Println(name)
	fmt.Println(dString)
}
