// Constant

// - Constant adalah variable yang nilainya tidak bis adiubah lagi setelah pertama kali diberi nilai
// - Cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan
//   adalah const, bukan var
// - Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya

// mari kita coba
package main

import "fmt"

func main() {
	const firstName string = "Riwanto" // penggunaan variable constan harus langsung di inisialisasikan nilainya
	const lastName = "Sitinjak"        // penulisan constant juga bisa tidak perlu menuliskan tipe datanya karena golang akan mengetahuinya

	// firstName = "Lestaria"; -> jika ini kita jalankan maka golang akan error karena const tidak bisa diubah value dari variabel nya
	// dibawah ini untuk mencetak variablenya
	fmt.Println(firstName)
	fmt.Println(lastName)
}
