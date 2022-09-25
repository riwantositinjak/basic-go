// Returning Multiple Values
// - Function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value
// - Untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data
//   return valuenya di function

// Lets Try

package main

import "fmt"

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	fmt.Println(firstName)
	fmt.Println(lastName)
}

func getFullName() (string, string) {
	return "Riwanto", "Sitinjak"
}
