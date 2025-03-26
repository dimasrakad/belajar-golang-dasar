package helper

import "fmt"

var version = "1.0.0" // tidak bisa diakses dariluar package
var Application = "golang"

// tidak bisa diakses dari luar package
func sayGoodBye(name string) string {
	return "Good bye" + name
}

func SayHello(name string) {
	fmt.Println("Hello", name)
}
