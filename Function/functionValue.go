// Function Value
// - Function adalah first class citizen
// - Function juga merupakan tipe data, dan bisa disimpan di dalam variable

// lets try brother and sister
package main

import "fmt"

func main() {
	gutBye := goodBye
	fmt.Println(gutBye("Riwanto"))
}

func goodBye(name string) string {
	return "Selamat tinggal " + name
}
