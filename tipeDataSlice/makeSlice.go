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
}
