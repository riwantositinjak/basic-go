// Type Declaration
// - Type Declaration adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
// - Type Declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada,
//   dengan tujuan agar lebih mudah dimengerti

package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPiwan NoKTP = "12312313123123123123"
	var marriedStatus Married = false

	fmt.Println(noKTPiwan)
	fmt.Println(marriedStatus)
}
