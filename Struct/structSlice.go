package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var slc []Person = []Person{
		Person{
			name: "Riwanto",
			age:  26,
		},
		Person{
			name: "Lisnawati Sitinjak",
			age:  20,
		},
		Person{
			name: "Lestaria Sitinjak",
			age:  14,
		},
		Person{
			name: "Rilista Olivia Sitinjak",
			age:  10,
		},
	}
	for i, v := range slc {
		fmt.Println(i)
		fmt.Println(v)
	}

	fmt.Println()
}
