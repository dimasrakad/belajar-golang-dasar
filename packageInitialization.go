package main

import (
	"belajar-golang-dasar/database"
	"fmt"
)

// import _ "belajar-golang-dasar/database"

func PackageInitialization() {
	result := database.GetDatabase()
	fmt.Println(result)
}
