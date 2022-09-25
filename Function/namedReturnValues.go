// Named Return Values
// - Biasanya saat kita memberitahu bahwa sebuah function mengembalikan value, maka kita hanya
//   mendeklarasikan tipe data return value di function
// - Namun kita juga bisa membuat variable secara langsung di tipe data return functionnya

// let's try
package main

import "fmt"

func main() {
	firstName, middleName, lastName := getCompletedName()
	fmt.Println(firstName, middleName, lastName)
}

// variable return value bisa kita tulis penamaannya langsung
func getCompletedName() (firstName, middleName, lastName string) {
	firstName = "Rilista"
	middleName = "Olivia"
	lastName = "Sitinjak"
	// return firstName, middleName, lastName -> begini juga bisa tapi return juga bisa karena
	return // kita sudah mendeklarasikan variablenya di return value sebelumnya
}
