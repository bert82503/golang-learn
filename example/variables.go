package main

import "fmt"

/*
In Go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.

var declares 1 or more variables.

You can declare multiple variables at once.

Go will infer the type of initialized variables.

Variables declared without a corresponding initialization are zero-valued.
For example, the zero value for an int is 0.

The := syntax is shorthand for declaring and initializing a variable,
e.g. for var f string = "apple" in this case.
*/
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
