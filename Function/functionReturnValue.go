// Function Return Value

// - Function bisa mengembalikan data
// - Untuk memberitahu bahwa function mengembalikan data, kita harus menuliskan tipe data kembalian
//   dari function tersebut
// - Jika function tersebut kita deklarasikan dengan tipe data pengembalian, maka wajib di dalam
//   function nya kita harus mengembalikan data
// - Untuk mengembalikan data dari function, kita bisa menggunakan kata kunci return, diikuti dengan
//   datanya

// Let's try
package main

import "fmt"

func main() {
	fmt.Println(getHello("Riwanto"))
	fmt.Println(getHello(""))
}

func getHello(name string) string {
	if name == "" {
		return "Hello, namamu siapa ?"
	} else {
		return "Hello " + name
	}
}
