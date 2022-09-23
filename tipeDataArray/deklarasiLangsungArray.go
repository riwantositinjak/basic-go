// Dalam mendeklarasi array juga kita bisa mendeklarasikannya secara langsung, tanpa memetametakan
// indeksnya

package main

import "fmt"

func main() {
	var number = [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(number)
}
