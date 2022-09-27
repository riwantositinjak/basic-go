// Struct

// - Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe
//   data lainnya dalam satu kesatuan
// - Struct biasanya representasi data dalam program aplikasi yang kita buat
// - Data di struct disimpan dalam field
// - Sederhananya struct adalah kumpulan dari field

// Membuat Struct
// - Struct adalah template data atau prototype data
// - Struct tidak bisa langsung digunakan
// - Namun kita bisa membuat data/object dari struct yang telah kita buat

// Let's try
package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var riwanto Customer
	riwanto.Name = "Riwanto"
	riwanto.Address = "Jakarta Timur"
	riwanto.Age = 26

	var lestaria Customer
	lestaria.Name = "Lestaria"
	lestaria.Address = "Bengkulu"
	lestaria.Age = 15

	fmt.Println(riwanto)
	fmt.Println(riwanto.Name)
	fmt.Println(riwanto.Address)
	fmt.Println(riwanto.Age)

	fmt.Println(lestaria)
}
