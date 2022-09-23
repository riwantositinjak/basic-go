// Switch dengan short statement
// - Sama dengan if, switch juga mendukung short statement sebelum variable yang akan di cek
//   kondisinya

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

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama Sudah benar")
	}
}
