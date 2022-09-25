// Anonymous Function
// - Sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tersebut
// - Namun kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter
//   tanpa harus membuat function terlebih dahulu
// - Hal tersebut dinamakan anonymous function atau function tanpa nama

// lets try
package main

import "fmt"

type Blacklist func(string) bool

func main() {
	Blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", Blacklist)
	registerUser("Riwanto", Blacklist)
}

func registerUser(name string, blacklist Blacklist) {
	if name == "Admin" {
		fmt.Println("You are blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}
