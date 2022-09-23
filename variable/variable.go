// Variable

// - Variable adalah tempat untuk menyimpan data
// - Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau
// - di GO-lang variable hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan
//   data yang berbeda-beda jenis, kita harus membuat beberapa variable
// - Untuk membuat variable, kita bisa menggunakan kata kunci var, lalu diikuti dengan nama variable
//   dan tipe datanya.


package main

import "fmt"

func main() {
	var name string

	name = "Riwanto Sitinjak"
	fmt.Println(name)

	name = "Iwan Sitinjak"
	fmt.Println(name)
}
