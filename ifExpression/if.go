// If Expression

// - If adalah salah satu kata kunci yang digunakan untuk percabangan
// - Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi
// - Hampir di semua bahasa pemrograman mendukung if expression

// mari kita coba langsung

package main

import "fmt"

func main() {
	name := "Eiwanto"

	if name == "Riwanto" {
		fmt.Println("Hello " + name)
	}
}
