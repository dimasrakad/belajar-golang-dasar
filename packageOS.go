package main

import (
	"fmt"
	"os"
)

func PackageOS() {
	args := os.Args
	fmt.Println(args)

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(hostname)
	}

	username := os.Getenv("USERNAME")
	fmt.Println(username)
}
