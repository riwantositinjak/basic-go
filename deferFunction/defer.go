// Defer

// - Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah
//   function selesai di eksekusi
// - Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi

package main

import "fmt"

func main() {
	runApplication(0)
}

func logging() {
	fmt.Println("Program selesai dijalankan")
}

func runApplication(value int) {
	defer logging() // penulisan defer harus yang paling atas, agar ketika program berhasil atau tidak akan tetap dijalankan oleh golang
	result := 10 / value
	fmt.Println("Hasilnya akan", result)
}
