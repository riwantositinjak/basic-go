// Struct Method

// - Struct adalah tipe data seperti tipe data seperti tipe data lainnya, dia bisa digunakan
//   sebagai parameter untuk function
// - Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct
//   memiliki function
// - Method adalah function

// Let's try
package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	riwanto := Customer{
		Name:    "Riwanto",
		Address: "Jakarta Timur",
		Age:     26,
	}
	fmt.Println(riwanto)

	riwanto.sayHi("Riwanto")
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "Nama saya adalah", customer.Name)
}
