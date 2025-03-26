package main

func RecursiveFunction(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * RecursiveFunction(value-1)
	}
}
