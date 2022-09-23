// For dengan Statement

// - Dalam for kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa ditambahkan
//   di for
// - Init statement, yaitu statement sebelum for dieksekusi
// - Post statement, yaitu statement yang akan selalu dieksekusi di akhir tiap perulangan

package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke - ", counter)
	}
}
