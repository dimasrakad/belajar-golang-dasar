package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func Pointer() {
	PassByValue()
	PassByReference()
}

func PassByValue() {
	/* PASS BY VALUE */
	address1 := Address{"Sidoarjo", "Jawa Timur", "Indonesia"}
	address2 := address1

	address2.City = "Surabaya"

	fmt.Println("PASS BY VALUE")
	fmt.Println(address1) // address1 tidak berubah
	fmt.Println(address2)
	fmt.Println()
}

func PassByReference() {
	/* PASS BY REFERENCE */
	address1 := Address{"Sidoarjo", "Jawa Timur", "Indonesia"}
	address2 := &address1
	address3 := &address1
	address4 := new(Address)

	address2.City = "Surabaya"
	fmt.Println("PASS BY REFERENCE")
	fmt.Println(address1) // address1 berubah
	fmt.Println(address2)
	fmt.Println()

	address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1) // address1 tidak berubah
	fmt.Println(address2)
	fmt.Println()

	*address3 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1) // address1 berubah
	fmt.Println(address3)
	fmt.Println()

	address4.City = "Bandung"
	fmt.Println(address4)
}
