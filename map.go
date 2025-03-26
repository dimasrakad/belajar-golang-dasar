package main

import "fmt"

func Map() {
	person := map[string]string{
		"name":    "Dims",
		"address": "Sidoarjo",
	}

	person["email"] = "dimas@email.com"

	fmt.Println(person)

	person["name"] = "Dimas"

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(len(person))

	var book map[string]string = make(map[string]string)
	book["title"] = "Book1"
	book["author"] = "Author1"
	book["salah"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "salah")

	fmt.Println(book)
	fmt.Println(len(book))
}
