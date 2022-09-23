// Function Parameter

// - Saat membuat function kadang-kadang kita membutuhkan data dari luar, atau kita sebut Parameter
// - Kita bisa menambahkan parameter di function, bisa lebih dari selanjutnya
// - Parameter tidaklah wajib, jadi kita bisa membuat function tanpa parameter seperti sebelumnya yang
//   sudah kita buat
// - Namun jika kita menambahkan parameter di function, maka ketika memanggil function tersebut, kita
//  Wajib memasukkan data ke parameternya.

// Lets try
package main

import "fmt"

func sayHello(firstName string, lastName string) { // masukkan paramater disertai tipe datanya
	fmt.Println("Hallo", firstName, lastName)
}

func main() {
	sayHello("Riwanto", "Sitinjak") // memasukkan argumennya untuk dipassing ke parameter func sayHello
}
