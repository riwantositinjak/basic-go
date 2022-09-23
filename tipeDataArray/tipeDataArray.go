// Tipe Data Array
// - Array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama
// - Saat membuat array, kita perlu menentukan jumlah data yang bisa ditampung oleh Array tersebut
// - Daya tampung array tidak bisa bertambah setelah Array dibuat
// - Index di golang dimulai dari index 0

package main

import "fmt"

func main() {
	// tanda [2] ini adalah jumlah maksimum data array dari namaSaya variable
	// jika lebih dari 2 maka akan terjadi error pada program
	var namaSaya [2]string
	namaSaya[0] = "Riwanto"
	namaSaya[1] = "Sitinjak"

	fmt.Println(namaSaya[0])
	fmt.Println(namaSaya[1])
	fmt.Print(namaSaya)
}
