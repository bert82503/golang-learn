package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

/*
Go supports pointers, allowing you to pass references to values and records within your program.
(指针，值引用)

We’ll show how pointers work in contrast to values with 2 functions: zeroval and zeroptr.
zeroval has an int parameter, so arguments will be passed to it by value.
zeroval will get a copy of ival distinct from the one in the calling function.

zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
The *iptr code in the function body then dereferences the pointer from
its memory address to the current value at that address.
Assigning a value to a dereferenced pointer changes the value at the referenced address.

The &i syntax gives the memory address of i, i.e. a pointer to i.

Pointers can be printed too.

zeroval doesn’t change the i in main, but zeroptr does because it has a reference to the memory address for that variable.
*/
func main() {
	i := 1
	fmt.Println("initial:", i)
	// initial: 1

	zeroval(i)
	fmt.Println("zeroval:", i)
	// zeroval: 1

	// The &i syntax gives the memory address of i, a pointer to i.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	// zeroptr: 0

	fmt.Println("pointer:", &i)
	// pointer: 0xc000198000
}
