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

/*
Go supports anonymous functions, which can form closures.
Anonymous functions are useful when you want to define a function inline without having to name it.

This function intSeq returns another function, which we define anonymously in the body of intSeq.
The returned function closes over the variable i to form a closure.

We call intSeq, assigning the result (a function) to nextInt.
This function value captures its own i value, which will be updated each time we call nextInt.

See the effect of the closure by calling nextInt a few times.

To confirm that the state is unique to that particular function, create and test a new one.

The last feature of functions we’ll look at for now is recursion.
*/
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
