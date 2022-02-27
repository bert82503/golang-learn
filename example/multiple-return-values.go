package main

import "fmt"

// The (int, int) in this function signature shows that the function returns 2 ints.
func vals() (int, int) {
	return 3, 7
}

/*
Go has built-in support for multiple return values.
This feature is used often in idiomatic Go, for example to return both result and error values from a function.

The (int, int) in this function signature shows that the function returns 2 ints.

Here we use the 2 different return values from the call with multiple assignment.

If you only want a subset of the returned values, use the blank identifier _.

Accepting a variable number of arguments is another nice feature of Go functions;
we’ll look at this next.
*/
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
