package main

import "fmt"

/*
In Go, an array is a numbered sequence of elements of a specific length.

Here we create an array a that will hold exactly 5 ints.
The type of elements and length are both part of the array’s type.
By default an array is zero-valued, which for ints means 0s.

We can set a value at an index using the array[index] = value syntax,
and get a value with array[index].

The builtin len returns the length of an array.

Use this syntax to declare and initialize an array in one line.

Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.

Note that arrays appear in the form [v1 v2 v3 ...] when printed with fmt.Println.

You’ll see slices much more often than arrays in typical Go.
We’ll look at slices next.
*/
func main() {
	var a [5]int
	fmt.Println("emp:", a)
	// emp: [0 0 0 0 0]

	a[4] = 100
	fmt.Println("set:", a)
	// set: [0 0 0 0 100]
	fmt.Println("get:", a[4])
	// get: 100

	fmt.Println("len:", len(a))
	// len: 5

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	// dcl: [1 2 3 4 5]

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	// 2d:  [[0 1 2] [1 2 3]]
}
