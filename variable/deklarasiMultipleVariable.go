// Deklarasi Multiple Variable

// -Di Go-Lang kita bisa membuat variable secara sekaligus banyak
// -Code yang dibuat akan lebih bagus dan mudah di baca

// Mari kita lakukan
package main

import "fmt"

func main() {
	var (
		firstName = "Riwanto"
		lastName  = "Sitinjak"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
