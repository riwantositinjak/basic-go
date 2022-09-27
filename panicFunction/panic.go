// Panic

// - Panic function adalah function yang bisa kita gunakan untuk menghentikan program
// - Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
// - Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan
//   dieksekusi

package main

import "fmt"

func main() {
	runningApp(false) // jika argumen diganti true akan mempengaruhi output dari func runningAppnya
}

func endApp() {
	fmt.Println("Program telah selesai dijalankan")
}

func runningApp(value bool) {
	defer endApp()
	if value {
		panic("Program ini tidak berhasil berjalan")
	}
	fmt.Println("Program sedang berjalan")
}
