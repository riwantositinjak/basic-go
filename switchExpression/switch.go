// Switch Expression
// - Selain if expression untuk melakukan percabangan, kita juga bisa menggunakan Switch Expression
// - Switch expression sangat sederhana dibandingkan if
// - Biasanya switch expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variable

// lets try
package main

import "fmt"

func main() {
	name := "Riwanto"

	switch name {
	case "Riwanto":
		fmt.Println("halo Riwanto")
	case "Lestaria":
		fmt.Println("Halo Lestaria")
	case "Rilista Olivia":
		fmt.Println("Halo Koling")
	default:
		fmt.Println("halo sebutkan namamu?")
	}
}
