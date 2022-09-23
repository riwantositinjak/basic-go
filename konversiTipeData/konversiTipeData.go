// Konversi Tipe Data

// - Di Go-Lang kadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe lain
// - Misal kita ingin mengkonversi tipe data int32 ke int 63, dan lain-lain

package main

import "fmt"

func main() {
	var nilai32 int32 = 299
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)
}
