// Break and Continue
// - Break & Continue adalah kata kunci yang bisa digunakan dalam perulangan
// - Break digunakan untuk menghentikan seluruh perulangan
// - Continue adalah digunakan untuk menghentikan perulangan yang berjalan, dan langsung
//   melanjutkan ke perulangan selanjutnya

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke- ", i)
	}
}
