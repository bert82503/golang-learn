package main

import "fmt"

// The (int, int) in this function signature shows that the function returns 2 ints.
func vals() (int, int) {
	return 3, 7
}

// Go has built-in support for multiple return values.
// This feature is used often in idiomatic Go, to return both result and error values from a function.
func main() {
	a, b := vals()
	fmt.Println(a)
	// 3
	fmt.Println(b)
	// 7

	_, c := vals()
	fmt.Println(c)
	// 7
}
