package main

func VariadicFunction(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
