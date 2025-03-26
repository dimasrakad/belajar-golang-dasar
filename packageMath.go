package main

import (
	"fmt"
	"math"
)

func PackageMath() {
	fmt.Println(math.Round(1.23))
	fmt.Println(math.Round(1.98))
	fmt.Println(math.Floor(1.98))
	fmt.Println(math.Ceil(1.23))
	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1, 2))
}
