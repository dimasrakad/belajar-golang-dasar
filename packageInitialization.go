package main

import (
	"belajar-golang-dasar/database"
	"fmt"
)

// Blank Identifier
// import _ "belajar-golang-dasar/internal"

func PackageInitialization() {
	result := database.GetDatabase()
	fmt.Println(result)
}
