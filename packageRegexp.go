package main

import (
	"fmt"
	"regexp"
)

func PackageRegexp() {
	regex := regexp.MustCompile("d([a-z])s")

	fmt.Println(regex.MatchString("dims"))
	fmt.Println(regex.MatchString("dim"))
	fmt.Println(regex.MatchString("dis"))
	fmt.Println(regex.MatchString("d4s"))
	fmt.Println(regex.MatchString("Dis"))

	fmt.Println(regex.FindAllString("dims dimas dima das dos dias d1m4s", 10))
}
