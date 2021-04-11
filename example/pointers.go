package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

// Go supports pointers, allowing you to pass references to values and records.
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
