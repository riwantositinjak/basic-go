// Tipe Data Slice
// - Tipe data Slice adalah potongan dari data array
// - Slice mirip dengan array, yang membedakan adalah ukuran Slice bisa berubah
// - Slice dan Array selalu terkoneksi, dimana slice adalah data yang mengakses
//   sebagian atau seluruh data di Array

// Detail Tipe Data Slice
// - Tipe data slice memiliki 3 data, yaitu pointer, length, dan capacity
// - Pointer adalah penunjuk data pertama di array pada Slice
// - Length adalah panjang dari slice, dan
// - Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

// let's Try

package main

import "fmt"

func main() {
	var month = [12]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Agustus",
		"November",
		"Desember",
	}
	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // untuk mencetak panjang slicenya
	fmt.Println(cap(slice1)) // untuk mencetak kapasitas slice di array month

	// month[5] = "Changed"  -> hati hati ketika mengganti value dari slice nya maka arraynya juga akan ikut terganti
	// fmt.Println(slice1)
 
}
