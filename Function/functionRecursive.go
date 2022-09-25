// Function Recursive
// - Recursive function adalah function yang memanggil function dirinya sendiri
// - Kadang dalam pekerjaan, kita sering menemui kasus dimana menggunakan recursive function lebih
//   mudah dibandingkan tidak menggunakan recursive function
// - Contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah Faktorial

package main

import "fmt"

func main() {
	result := recursiveLoop(5)
	fmt.Println(5 * 4 * 3 * 2 * 1) // ini hanya untuk mengecek saja apakah faktorial sudah sesuai
	fmt.Println(result)
}

// func faktorialLoop(value int) int {
// 	total := 1

// 	for i := value; i > 0; i-- {  -> looping for ini adalah contoh jika ingin mencari faktorial
// 		total *= i
// 	}
// 	return total
// }

func recursiveLoop(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursiveLoop(value-1)
	}
}
