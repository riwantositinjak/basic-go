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
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = month[11:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Riwanto")
	fmt.Println(slice3)
	fmt.Println(slice2)
}
