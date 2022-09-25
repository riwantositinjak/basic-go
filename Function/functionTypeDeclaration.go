// Function Type Declarations

// - Kadang jika function terlalu panjang, agak ribet untuk menuliskannya di dalam parameter
// - Type Declarations juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah
//   kita menggunakan function sebagai parameter

package main

import "fmt"

type Filter func(string) string // ini adalah penggunaakan tipe declaration untuk aliasing function sebagai parameter

func main() {
	sayHelloWithFilter("Riwanto", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}

func sayHelloWithFilter(name string, filter Filter) { 
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
