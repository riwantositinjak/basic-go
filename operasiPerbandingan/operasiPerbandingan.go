// Operasi Perbandingan

// - Operasi perbandingan adalah operasi untuk membandingkan dua buah datanya
// - Operasi perbandingan adalah operasi yang menghasilkan nilai boolean(benar atau salah);
// - Jika hasil operasinya adalah benar, maka nilainya adalah true
// - Jika hasil operasinya adalah salah, maka nilai adalah false

// Mari kita coba

package main

import "fmt"

func main() {
	var namaPertama string = "Riwanto"
	var namaKedua string = "Riwanto"
	var namaKetiga string = "Lestaria"
	var perbandinganTrue bool = namaPertama == namaKedua
	var perbandinganFalse bool = namaPertama == namaKetiga
	var perbandinganLebihDari bool = namaPertama > namaKetiga

	fmt.Println(perbandinganTrue)
	fmt.Println(perbandinganFalse)
	fmt.Println(perbandinganLebihDari)
}
