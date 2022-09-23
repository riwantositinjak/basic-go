package main

import "fmt"

func main() {
	var a = 10
	a += 10 // a = a + 10

	var b = 20
	b += a; // b = b + a;

	fmt.Println(a)
	fmt.Println(b)
}
