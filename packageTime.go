package main

import (
	"fmt"
	"time"
)

func PackageTime() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())
	fmt.Println(now.Unix())
	fmt.Println(now.UTC())

	utc := time.Date(2023, 12, 25, 7, 7, 7, 7, time.UTC)
	fmt.Println(utc)

	layout := time.DateOnly
	parse, _ := time.Parse(layout, "2023-12-25")
	fmt.Println(parse)

}
