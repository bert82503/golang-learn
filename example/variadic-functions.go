package main

import "fmt"

// a function that will take an arbitrary number of ints as arguments.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// Variadic functions can be called with any number of trailing arguments.
func main() {
	sum(1, 2)
	// [1 2] 3
	sum(1, 2, 3)
	// [1 2 3] 6

	// If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	// [1 2 3 4] 10
}
