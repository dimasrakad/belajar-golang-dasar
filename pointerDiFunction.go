package main

import "fmt"

func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func PointerDiFunction() {
	address := Address{"Sidoarjo", "Jawa Timur", ""}
	ChangeAddressToIndonesia(&address)

	fmt.Println(address)
}
