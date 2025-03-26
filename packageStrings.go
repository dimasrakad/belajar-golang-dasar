package main

import (
	"fmt"
	"strings"
)

func PackageStrings() {
	name := "Dimas Raka D"
	fmt.Println(strings.Contains(name, "Dimas"))
	fmt.Println(strings.Split(name, " "))
	fmt.Println(strings.ToLower(name))
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.Trim("     "+name+"     ", " "))
	fmt.Println(strings.ReplaceAll(name, "Dimas", "Dims"))
}
