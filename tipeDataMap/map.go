// Tipe Data Map

// - Pada Array atau slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0
// - Map adalah tipe data lain yang berisikan kumpulan data yagn sama, namun kita bisa menentukan
//   jenis tipe data index yang akan kita gunakan
// - Sederhananya, Map adalah tipe data kumpulan key-value (kata kunci- nilai) dimana kata kuncinya
//   bersifat unik, tidak boleh sama
// - Berbeda dengan Array dan Slice, jumlah data yang kita masukkan ke dalam Map boleh sebanyak
//   banyaknya, asalkan kata kunci nya berbeda, jika kita gunakan kata kunci sama, maka secara
//   otomatis data sebelumnya akan diganti dengan data baru

// Ada banyak function yang bisa kita gunakan untuk memanipulasi mapnya diantaranya adalah
// - len(map) -> berguna untuk mendapatkan jumlah data di map
// - map[key] -> berguna untuk mengambil data di map dengan key
// - map[key] = value -> berguna untuk mengubah data di map dengan key
// - make(map[TypeKey]TypeValue) -> berguna untuk membuat map baru
// - delete(map, key) -> berguna untuk menghapus data di map dengan key

package main

import "fmt"

func main() {
	var person = map[string]string{
		"name":   "Riwanto",
		"alamat": "Jakarta",
	}

	person["hobby"] = "Ngoding" // menambah key dan value pada objek person

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["alamat"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["penulis"] = "Riwanto"
	book["ini salah"] = "salah"

	delete(book, "ini salah") // menghapus 

	fmt.Println(book)
	fmt.Println(len(book));
}
