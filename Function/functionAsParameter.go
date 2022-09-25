// Function as Parameter
// - Function tidak hanya bisa kita simpan di dalam variable sebagai value
// - Namun juga bisa kita gunakan sebagai parameter untuk function lain

// lets try

package main

import "fmt"

func main() {
	sayHelloWithFilter("Riwanto", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}

func sayHelloWithFilter(name string, filter func(string) string) {
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
