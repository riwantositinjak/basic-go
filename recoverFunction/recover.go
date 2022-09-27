// Recover

// - Recover adalah function yang bisa kita gunakan untuk menangkap data panic
// - Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan

package main

import "fmt"

func main() {
	runApp(true)
	fmt.Println("Program masih akan berjalan setelah di recover")
}

func endApp() {
	fmt.Println("Aplikasi selesai")
	message := recover()
	if message != nil {
		fmt.Println("Terjadi error dengan pesan: ", message)
	}

}

func runApp(value bool) {
	defer endApp()
	if value {
		panic("Terjadi error panic")
	}
}
