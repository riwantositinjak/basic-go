// Deklarasi Multiple Constant
// - Sama seperti variable, di Go-lang juga kita bisa membuat constant secara sekaligus banyak

// Mari kita coba

package main

import "fmt"

func main() {
	const (
		firstName = "Riwanto"
		lastName  = "Sitinjak"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
