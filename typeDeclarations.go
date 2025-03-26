package main

import "fmt"

func TypeDeclarations() {
	type NoKTP string
	type Married bool

	var noKtpDims NoKTP = "1234"
	var marriedStatus Married = false
	fmt.Println(noKtpDims)
	fmt.Println(marriedStatus)
}
