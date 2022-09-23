// if dengan Short Statement

// - If mendukung short statement sebelum kondisi
// - Hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan
//   terhadap kondisi

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

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
