// Struct Literals

// - Sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa
//   kita gunakan untuk membuat data dari struct

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

	lestaria := Customer{"lestaria", "bengkulu", 15} // cara ini memang lebih singkat tapi cara ini tergantung dari urutan field di type struct nya
	// apabila struct nya kita ganti urutan fieldnya maka akan terjadi error
	fmt.Println(lestaria)
}
