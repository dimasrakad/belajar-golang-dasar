package main

import "fmt"

func main() {
	str := "  abcd  efg h   "
	fmt.Printf("before:\t\"%s\"\n", str)

	for (string(str[0]) == " ") || (string(str[len(str)-1]) == " ") {
		if string(str[0]) == " " {
			str = string(str[1:])
		}
		if string(str[len(str)-1]) == " " {
			str = string(str[:len(str)-1])
		}
	}
	fmt.Printf("after:\t\"%s\"\n", str)

	// helloworld
	// Helloworld()

	// number
	// Number()

	// boolean
	// Boolean()

	// string
	// String()

	// variable
	// Variable()

	// konversi tipe data
	// KonversiTipeData()

	// type declarations
	// TypeDeclarations()

	// operasi matematika
	// OperasiMatematika()

	// operasi perbandingan
	// OperasiPerbandingan()

	// operasi boolean
	// OperasiBoolean()

	// array
	// Array()

	// slice
	// Slice()

	// map
	// Map()

	// if expression
	// IfExpression()

	// switch expression
	// SwitchExpression()

	// for loops
	// ForLoops()

	// break and continue
	// BreakAndContinue()

	// function parameter
	// FunctionParameter("Dims", "Raka")

	// function return value
	// fmt.Println(FunctionReturnValue("Dims"))

	// returning multiple values
	// firstName, lastName := ReturningMultipleValues()
	// fmt.Println(firstName, lastName)
	// firstName2, _ := ReturningMultipleValues()
	// fmt.Println(firstName2)

	// named return value
	// firstName, middleName, lastName := NamedReturnValue()
	// fmt.Println(firstName, middleName, lastName)
	// firstName2, middleName2, _ := NamedReturnValue()
	// fmt.Println(firstName2, middleName2)

	// variadic function
	// fmt.Println(VariadicFunction(1, 2, 3, 4))
	// slice := []int{5, 10, 15, 20}
	// fmt.Println(VariadicFunction(slice...))

	// function value
	// hello := FunctionValue
	// fmt.Println(hello("Dims"))

	// function as parameter
	// FunctionAsParameter("Dims", toxicFilter)
	// FunctionAsParameter("Anjing", toxicFilter)

	// anonymous function
	// blacklist := func(name string) bool {
	// 	return name == "Anjing"
	// }
	// AnonymousFunction("Dims", blacklist)
	// AnonymousFunction("Anjing", blacklist)

	// recursive function
	// fmt.Println(RecursiveFunction(5))

	// closure
	// Closure()

	// defer
	// Defer(0)

	// panic and recover
	// PanicAndRecover(true)

	// comment
	// single line
	/*
		multi
		line
	*/

	// struct
	// Struct()

	// interface
	// Interface()

	// interface kosong
	// InterfaceKosong()

	// nil
	// Nil()

	// error interface
	// ErrorInterface()

	// type assertions
	// TypeAssertions()

	// pointer
	// Pointer()

	// pointer di function
	// PointerDiFunction()

	// pointer di method
	// PointerDiMethod()

	// package dan import
	// PackageDanImport()

	// package initialization
	// PackageInitialization()

	// package OS
	// PackageOS()

	// package flag
	// PackageFlag()

	// package strings
	// PackageStrings()

	// package strconv
	// PackageStrconv()

	// package math
	// PackageMath()

	// package container/list
	// PackageContainerList()

	// package container/ring
	// PackageContainerRing()

	// package sort
	// PackageSort()

	// package time
	// PackageTime()

	// package reflect
	// PackageReflect()

	// package regexp
	// PackageRegexp()

	// Custom Error
	CustomError()
}
