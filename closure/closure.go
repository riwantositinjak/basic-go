// Closure
// - Closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam
//   scope yang sama
// - Harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi

// Let's try

package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		counter := 2
		counter++
	}
	increment()
	increment()
	fmt.Println(counter)
}
