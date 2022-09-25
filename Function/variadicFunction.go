// Variadic Function
// - Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
// - Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam Array
// - Apa bedanya dengan parameter biasa dengan tipe data Array ?
// 	- Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
// 	- Jika parameter menggunakan varargs, kita bisa langsung mengirim datanya, jika lebih dari satu, cukup
// 	  gunakan tanda koma.

// let's try
package main

import "fmt"

func main() {
	result := sumAll(10, 10, 20, 22, 21)

	fmt.Println(result)
}

func sumAll(number ...int) int {
	total := 0
	for _, numbers := range number {
		total += numbers
	}
	return total
}
