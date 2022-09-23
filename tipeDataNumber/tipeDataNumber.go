/* Tipe Data Number

Ada 2 jenis tipe data Number
- Integer
- Floating Point

Golang tipe data bersifat case sensitive
tipe data int memiliki nilai minimum negatif sekian sedangkan tipe data uint
nilai minimumnya dimulai dari 0
tipe data float memiliki 2 tipe float32 dan float 64

Alias
alias adalah nama lain
byte = uint9
rune = int32
int = minimal int32 tergantung berapa bit sistem operasi kamu
uin = minimal uint32 tergantung berapa bit sistem operasi kamu */

package main

import "fmt"

func main(){
	fmt.Println("Satu = ", 1);
	fmt.Println("Dua = ", 2);
	fmt.Println("Tiga Koma Lima = ", 3.5);
}