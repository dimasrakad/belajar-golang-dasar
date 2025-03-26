package main

import (
	"flag"
	"fmt"
)

func PackageFlag() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")

	flag.Parse()

	fmt.Println("host\t\t:", *host)
	fmt.Println("username\t:", *username)
	fmt.Println("password\t:", *password)
}
