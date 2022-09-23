package main

import "fmt"

func main() {
	var nilai = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilai > absensi
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus) // output false kareena salah satu dari boolean operator and bernilai false
}
