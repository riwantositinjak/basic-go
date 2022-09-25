// Mengabaikan Return Value
// - Multiple return value wajib ditangkap semua valuenya
// - Jika kita ingin mengabaikan return value tersebut, kita bisa menggunakan tanda _ (garis bawah)

// Let's try
package main

import "fmt"

func main() {
	firstName, _ := getFullName() // return value yang kedua bisa ditulis begini untuk menandakan
	fmt.Println(firstName)        // mengabaikan return second valuenya
}

func getFullName() (string, string) {
	return "Riwanto", "Sitinjak"
}
