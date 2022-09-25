// Slice Parameter
// - Kadang ada kasus dimana kita menggunakan Variadic Function, namun memiliki variable berupa Slice
// - Kita bisa menjadikan slice sebagai vararg parameter

package main

import "fmt"

func main() {
	total := sumAll(10, 20, 212, 122, 2, 2)
	fmt.Println(total)

	slice := []int{10, 12, 2, 2, 2, 2, 2}
	total = sumAll(slice...)
	fmt.Println(total)
}

func sumAll(number ...int) int {
	total := 0
	for _, numbers := range number {
		total += numbers
	}
	return total
}
