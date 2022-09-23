// For
// - Dalam bahasa pemrogramana, biasanya ada fitur yang bernama perulangan
// - For adalah satu satunya fitur perulangan untuk bahasa pemrograman go sampai pada saat ini

// mari kita coba
package main

import "fmt"

func main() {
	number := 1

	for number <= 10 {
		fmt.Println("Pengulangan ke - ", number)
		number++
	}
}
