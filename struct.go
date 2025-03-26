package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello", customer.Name)
}

func Struct() {
	var customer Customer
	customer.Name = "Dims"
	customer.Address = "Sidoarjo"
	customer.Age = 8

	customer2 := Customer{
		Name:    "Raka",
		Address: "Surabaya",
		Age:     9,
	}

	customer3 := Customer{"Dewa", "Jakarta", 10}

	fmt.Println(customer)
	customer.sayHello()
	fmt.Println(customer2)
	customer2.sayHello()
	fmt.Println(customer3)
	customer3.sayHello()
}
