package main

import "fmt"

// Variables are explicitly declared and used by the compiler to check type-correctness of function calls.
func main() {
	var a = "initial"
	fmt.Println(a)
	// initial

	var m, n int = 1, 2
	fmt.Println(m, n)
	// 1 2

	var b = true
	fmt.Println(b)
	// true

	var e int
	fmt.Println(e)
	// 0

	f := "apple"
	fmt.Println(f)
	// apple
}
