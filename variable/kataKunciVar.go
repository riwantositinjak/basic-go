// Kata Kunci Var

// - Di Golang, kata kunci var saat membuat variable tidaklah wajib
// - Asalkan saat membuat variable kita langsung menginisialisasi datanya
// - Agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan kata kunci := saat
//   menginisialisasikan data pada variable tersebut

// MARI KITA LAKUKAN

package main

import "fmt"

func main() {
	nama := "Riwanto"
	fmt.Println(nama)

	nama = "Sitinjak"
	fmt.Println(nama)
}
