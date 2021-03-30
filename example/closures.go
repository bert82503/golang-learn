package main

import "fmt"

// The returned function closes over the variable i to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Go supports anonymous functions, which can form closures.
// Anonymous functions are useful when you want to define a function inline without having to name it.
func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	// 1
	fmt.Println(nextInt())
	// 2
	fmt.Println(nextInt())
	// 3

	// To confirm that the state is unique to that particular function, create and test a new one.
	newInts := intSeq()
	fmt.Println(newInts())
	// 1
}
