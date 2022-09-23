// Else if Expression
// - Kadang dalam if, kita butuh membuat beberapa kondisi
// - Kasus seperti ini, kita bisa menggunakan Else if expression

package main

import "fmt"

func main() {
	name := "Lestaria"

	if name == "Riwanto" {
		fmt.Println("Hello Riwanto")
	} else if name == "Lestaria" {
		fmt.Println("Hello  Lestaria")
	} else {
		fmt.Println("Tidak sesuai ")
	}
}
