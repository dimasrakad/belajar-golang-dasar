package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func PointerDiMethod() {
	man := Man{"Dims"}
	man.Married()
	fmt.Println(man)
}
