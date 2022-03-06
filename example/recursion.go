package main

import "fmt"

// This fact function calls itself until it reaches the base case of fact(0).
// 递归函数
func fact(n int) int {
	// 退出条件
	if n == 0 {
		return 1
	}
	// 递归调用，分而治之的思想
	return n * fact(n-1)
}

/*
Go supports recursive functions.
Here’s a classic example.
(递归函数)

This fact function calls itself until it reaches the base case of fact(0).
(终止条件)

Closures can also be recursive, but this requires the closure to be declared with a typed var explicitly before it’s defined.

Since fib was previously declared in main, Go knows which function to call with fib here.
*/
func main() {
	fmt.Println(fact(7))
	// 5040

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
