// Else Expression

// - Blok if akan dieksekusi ketika kondisi if bernilai true
// - Kadang kita ingin melakukan eksekusi program tertentu jika kondisi if bernilai false
// - Hal ini bisa dilakukan menggunakan else expression

package main

import "fmt"

func main() {
	name := "Eiwanto"

	if name == "Riwanto" {
		fmt.Println("Hello " + name)
	} else {
		fmt.Println("Tidak sesuai ")
	}
}
