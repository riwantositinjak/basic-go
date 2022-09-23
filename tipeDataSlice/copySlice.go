package main

import "fmt"

func main() {
	// ini adalah cara paling aman untuk membuat slice dari awal, karena arraynya didalam slicenya
	// karena tidak disimpan di dalam sebuah variabel
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Riwanto"
	newSlice[1] = "Sitinjak"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println((copySlice))
}

// Catatan untuk membuat array
// - Saat membuat array kita harus berhati-hati, jika salah, maka yang kita buat bukanlah Array,
//   melainkan Slice

// begini contohnya :

// package main

// import "fmt"

// func main(){
// 	iniArray := [...]int{1,2,3,4,5,6} -> jangan sampai lupa menuliskan[nilai indexnya]
// 	iniSlice:= []int{1,2,3,4,5,6} -> jika lupa maka golang akan menganggap kalian menuliskan slice bukan array

// }
