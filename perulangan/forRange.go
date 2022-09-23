// For Range
// - For bisa digunakan untuk melakukan iterasi terhadap semua data collection
// - Data collection contohnya Array, Slice, Map

package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke - ", counter)
	}
	slice := []string{"Riwanto", "Sitinjak", "Lestaria", "Sitinjak"}

	// for i:= 0 ; i < len(slice); i++ {
	// 	fmt.Println(i)
	// 	fmt.Println(slice[i]);
	// }

	for index, value := range slice {
		fmt.Println("index", index, "=", value)
	}
}
